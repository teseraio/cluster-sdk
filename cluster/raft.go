package cluster

import (
	"fmt"
	"time"

	"github.com/hashicorp/raft"
)

const (
	raftTimeout = 10 * time.Second
)

func (s *Server) SetupRaft(fsm raft.FSM) error {
	s.fsm = fsm

	s.raftLayer = NewRaftLayer(s.Config.GRPCAddr)
	s.raftLayer.srv = s

	raftConfig := raft.DefaultConfig()

	trans := raft.NewNetworkTransport(s.raftLayer, 3, raftTimeout, raftConfig.LogOutput)

	raftConfig.LogOutput = s.Config.LogOutput
	raftConfig.LocalID = raft.ServerID(s.Config.NodeName)

	store := raft.NewInmemStore()
	stable := store
	log := store
	snap := raft.NewDiscardSnapshotStore()

	raftConfig.NotifyCh = s.LeaderCh

	if s.Config.BootstrapExpected == 1 {
		configuration := raft.Configuration{
			Servers: []raft.Server{
				{
					ID:      raftConfig.LocalID,
					Address: trans.LocalAddr(),
				},
			},
		}

		if err := raft.BootstrapCluster(raftConfig, log, stable, snap, trans, configuration); err != nil {
			return fmt.Errorf("Failed to bootstrap initial cluster: %v", err)
		}
	}

	client, err := raft.NewRaft(raftConfig, s.fsm, log, stable, snap, trans)
	if err != nil {
		return fmt.Errorf("Failed to start raft: %v", err)
	}

	s.Raft = client

	go s.reconcile()
	return nil
}

func (s *Server) reconcile() {
	for {
		select {
		case member := <-s.reconcileCh:
			fmt.Println("reconcile")

			id := member.Tags["id"]
			if id == "" {
				s.logger.Printf("Id not found for member addr: %s", member.Addr.String())
				continue
			}

			raftAddr := fmt.Sprintf("%s:%d", member.Addr.String(), member.Port)
			addFuture := s.Raft.AddVoter(raft.ServerID(id), raft.ServerAddress(raftAddr), 0, 0)

			if err := addFuture.Error(); err != nil {
				s.logger.Printf("Failed to add peer %s to raft", id)
			}
		}
	}
}
