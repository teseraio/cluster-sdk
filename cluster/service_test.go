package cluster

import (
	"bytes"
	"crypto/rand"
	"io"
	"net"
	"testing"
	"time"

	"github.com/hashicorp/raft"
	"github.com/teseraio/cluster-sdk/cluster/proto"
)

func TestStream(t *testing.T) {
	c0 := NewServer(&Config{GRPCAddr: &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 5000}}, nil)
	c1 := NewServer(&Config{GRPCAddr: &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 5001}}, nil)

	if err := c0.SetupGRPC(); err != nil {
		t.Fatal(err)
	}
	if err := c1.SetupGRPC(); err != nil {
		t.Fatal(err)
	}

	conn0, err := c0.createStream("127.0.0.1:5001", 5*time.Second, proto.StreamObj_Start_SERF)
	if err != nil {
		t.Fatal(err)
	}
	conn1 := <-c1.streams[proto.StreamObj_Start_SERF]

	wdata := make([]byte, 1024)
	rand.Read(wdata)

	closeCh := make(chan struct{})

	go func(t *testing.T) {
		if _, err := conn0.Write(wdata); err != nil {
			t.Fatal(err)
		}

		data2 := make([]byte, 1024)
		if _, err := conn0.Read(data2); err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(wdata, data2) {
			t.Fatal(err)
		}
		close(closeCh)
	}(t)

	data := make([]byte, 1024)
	if _, err := conn1.Read(data); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(wdata, data) {
		t.Fatal("bad")
	}

	if _, err := conn1.Write(data); err != nil {
		t.Fatal(err)
	}
	<-closeCh
}

type mockFSM struct {
}

func newMockFSM() raft.FSM {
	return &mockFSM{}
}

func (f *mockFSM) Apply(l *raft.Log) interface{} {
	return nil
}

func (f *mockFSM) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}

func (f *mockFSM) Restore(io.ReadCloser) error {
	return nil
}
