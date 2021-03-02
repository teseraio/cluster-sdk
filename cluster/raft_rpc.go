package cluster

import (
	"fmt"
	"net"
	"time"

	"github.com/hashicorp/raft"
	"github.com/teseraio/cluster-sdk/cluster/proto"
)

type RaftLayer struct {
	srv     *Server
	addr    net.Addr
	connCh  chan net.Conn
	closeCh chan struct{}
}

func NewRaftLayer(addr net.Addr) *RaftLayer {
	layer := &RaftLayer{
		addr:    addr,
		connCh:  make(chan net.Conn),
		closeCh: make(chan struct{}),
	}
	return layer
}

func (l *RaftLayer) Accept() (net.Conn, error) {
	select {
	case conn := <-l.srv.streams[proto.StreamObj_Start_RAFT]:
		return conn, nil
	case <-l.closeCh:
		return nil, fmt.Errorf("Raft RPC layer closed")
	}
}

func (l *RaftLayer) Close() error {
	close(l.closeCh)
	return nil
}

func (l *RaftLayer) Addr() net.Addr {
	return l.addr
}

func (l *RaftLayer) Dial(address raft.ServerAddress, timeout time.Duration) (net.Conn, error) {
	return l.srv.createStream(string(address), timeout, proto.StreamObj_Start_RAFT)
}
