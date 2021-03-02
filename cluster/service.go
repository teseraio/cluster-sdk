package cluster

import (
	"context"
	"fmt"
	"net"

	gproto "github.com/golang/protobuf/proto"
	"github.com/golang/protobuf/ptypes/empty"

	grpc_net_conn "github.com/mitchellh/go-grpc-net-conn"
	"github.com/teseraio/cluster-sdk/cluster/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type systemService struct {
	s *Server
}

type wrapConn struct {
	net.Conn
	remoteAddr net.Addr
}

func (w *wrapConn) RemoteAddr() net.Addr {
	return w.remoteAddr
}

func (s *systemService) IsLeader(ctx context.Context, req *empty.Empty) (*proto.IsLeaderResp, error) {
	resp := &proto.IsLeaderResp{
		IsLeader: s.s.IsLeader(),
	}
	return resp, nil
}

func (s *systemService) Stream(stream proto.SystemService_StreamServer) error {
	resp, err := stream.Recv()
	if err != nil {
		return err
	}
	start, ok := resp.Event.(*proto.StreamObj_Start_)
	if !ok {
		return fmt.Errorf("expected start message")
	}

	// send an open message
	err = stream.Send(&proto.StreamObj{
		Event: &proto.StreamObj_Open_{
			Open: &proto.StreamObj_Open{},
		},
	})
	if err != nil {
		return err
	}

	p, ok := peer.FromContext(stream.Context())
	if !ok {
		return fmt.Errorf("peer not found in context")
	}

	// create net.Conn stream
	conn := streamWriter(stream)

	wConn := &wrapConn{
		Conn:       conn,
		remoteAddr: p.Addr,
	}

	s.s.handoff(start.Start.Channel, wConn)
	<-stream.Context().Done()

	return nil
}

func streamWriter(stream grpc.Stream) net.Conn {
	return &grpc_net_conn.Conn{
		Stream:   stream,
		Response: &proto.StreamObj{},
		Request:  &proto.StreamObj{},
		Encode: grpc_net_conn.SimpleEncoder(func(msg gproto.Message) *[]byte {
			req := msg.(*proto.StreamObj)
			if req.Event == nil {
				req.Event = &proto.StreamObj_Input_{
					Input: &proto.StreamObj_Input{},
				}
			}
			return &req.Event.(*proto.StreamObj_Input_).Input.Data
		}),
		Decode: grpc_net_conn.SimpleDecoder(func(msg gproto.Message) *[]byte {
			req := msg.(*proto.StreamObj)
			if req.Event == nil {
				req.Event = &proto.StreamObj_Input_{
					Input: &proto.StreamObj_Input{},
				}
			}
			return &req.Event.(*proto.StreamObj_Input_).Input.Data
		}),
	}
}
