package cluster

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"github.com/hashicorp/raft"
	"github.com/hashicorp/serf/serf"
	"github.com/teseraio/cluster-sdk/cluster/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type Server struct {
	logger *log.Logger
	Config *Config

	fsm raft.FSM

	// serf
	Serf         *serf.Serf
	localEventCh chan serf.Event
	EventCh      chan serf.Event

	// raft
	Raft        *raft.Raft
	LeaderCh    chan bool
	reconcileCh chan serf.Member

	// grpc
	grpcServer  *grpc.Server
	rpcListener net.Listener
	raftLayer   *RaftLayer

	// conn watchers
	streams map[proto.StreamObj_Start_Channel]chan net.Conn
}

func NewServer(config *Config, logger *log.Logger) *Server {
	s := &Server{
		logger:       logger,
		Config:       config,
		reconcileCh:  make(chan serf.Member, 10),
		localEventCh: make(chan serf.Event, 10),
		EventCh:      make(chan serf.Event, 10),
		LeaderCh:     make(chan bool, 10),
		streams: map[proto.StreamObj_Start_Channel]chan net.Conn{
			proto.StreamObj_Start_SERF: make(chan net.Conn),
			proto.StreamObj_Start_RAFT: make(chan net.Conn),
		},
		grpcServer: grpc.NewServer(),
	}
	return s
}

func (s *Server) SetupGRPC() error {
	proto.RegisterSystemServiceServer(s.grpcServer, &systemService{s})

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", s.Config.GRPCAddr.Port))
	if err != nil {
		return err
	}
	s.rpcListener = lis

	go func() {
		s.grpcServer.Serve(lis)
	}()
	return nil
}

func (s *Server) GRPCServer() *grpc.Server {
	return s.grpcServer
}

func (s *Server) handoff(channel proto.StreamObj_Start_Channel, conn net.Conn) {
	s.streams[channel] <- conn
}

func (s *Server) IsLeader() bool {
	return s.Raft.State() == raft.Leader
}

func (s *Server) Apply(cmd []byte, timeout time.Duration) raft.ApplyFuture {
	return s.Raft.Apply(cmd, timeout)
}

func (s *Server) createStream(addr string, timeout time.Duration, channel proto.StreamObj_Start_Channel) (net.Conn, error) {
	grpcConn, err := grpc.Dial(addr, grpc.WithInsecure(), grpc.WithTimeout(timeout))
	if err != nil {
		return nil, err
	}

	client := proto.NewSystemServiceClient(grpcConn)
	stream, err := client.Stream(context.Background())
	if err != nil {
		return nil, err
	}

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return nil, fmt.Errorf("peer not found in client stream context")
	}

	// send an start connection
	err = stream.Send(&proto.StreamObj{
		Event: &proto.StreamObj_Start_{
			Start: &proto.StreamObj_Start{
				Channel: channel,
			},
		},
	})
	if err != nil {
		return nil, err
	}

	// wait for an open connection
	resp, err := stream.Recv()
	if err != nil {
		return nil, err
	}
	if _, ok := resp.Event.(*proto.StreamObj_Open_); !ok {
		return nil, fmt.Errorf("expected open message")
	}

	// create the stream
	conn := streamWriter(stream)
	wConn := &wrapConn{
		Conn:       conn,
		remoteAddr: p.Addr,
	}
	return wConn, nil
}
