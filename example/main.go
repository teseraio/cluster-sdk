package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/hashicorp/raft"
	"github.com/teseraio/cluster-sdk/cluster"
)

func main() {
	logger := log.New(os.Stderr, "", log.LstdFlags)

	config := cluster.DefaultConfig()

	flag.StringVar(&config.NodeName, "node-name", "", "Node name")
	flag.IntVar(&config.GRPCAddr.Port, "grpc-port", 0, "GRPC bind port")
	flag.Int64Var(&config.BootstrapExpected, "bootstrap", 1, "Bootstrap expected")

	flag.Parse()

	server := cluster.NewServer(config, logger)

	if err := server.SetupGRPC(); err != nil {
		panic(err)
	}
	if err := server.SetupRaft(NewFSM()); err != nil {
		panic(err)
	}
	if err := server.SetupSerf(); err != nil {
		panic(err)
	}

	server.Serf.Join([]string{"127.0.0.1:5000"}, false)

	go func() {
		for {
			select {
			case m := <-server.EventCh:
				fmt.Printf("New member: %s\n", m.String())
			}
		}
	}()

	go func() {
		for {
			select {
			case isLeader := <-server.LeaderCh:
				if isLeader {
					fmt.Println("Is leader now")
				} else {
					fmt.Println("Not leader anymore")
				}
			}
		}
	}()

	done := make(chan bool)
	<-done
}

type SimpleFSM struct {
}

func NewFSM() raft.FSM {
	return &SimpleFSM{}
}

func (f *SimpleFSM) Apply(l *raft.Log) interface{} {
	return nil
}

func (f *SimpleFSM) Snapshot() (raft.FSMSnapshot, error) {
	return nil, nil
}

func (f *SimpleFSM) Restore(io.ReadCloser) error {
	return nil
}
