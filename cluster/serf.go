package cluster

import (
	"fmt"
	"net"
	"sync/atomic"

	"github.com/hashicorp/raft"
	"github.com/hashicorp/serf/serf"
)

func (s *Server) SetupSerf() error {
	conf := serf.DefaultConfig()
	conf.Init()

	conf.Tags["id"] = s.Config.NodeName
	for k, v := range s.Config.Tags {
		conf.Tags[k] = v
	}

	nc := &streamTransportConfig{
		BindAddr: s.Config.GRPCAddr.IP.String(),
		BindPort: s.Config.GRPCAddr.Port,
	}
	trans, err := newStreamTransport(nc)
	if err != nil {
		return err
	}
	trans.server = s
	trans.tcpListener = s.rpcListener.(*net.TCPListener)

	conf.NodeName = s.Config.NodeName
	conf.MemberlistConfig.Transport = trans
	conf.MemberlistConfig.LogOutput = s.Config.LogOutput
	conf.LogOutput = s.Config.LogOutput
	conf.EventCh = s.localEventCh

	client, err := serf.Create(conf)
	if err != nil {
		return err
	}

	s.Serf = client
	go s.eventHandler()

	return nil
}

func (s *Server) eventHandler() {
	for {
		select {
		case e := <-s.localEventCh:

			select {
			case s.EventCh <- e:
			default:
			}

			switch e.EventType() {
			case serf.EventMemberJoin:
				s.nodeJoin(e.(serf.MemberEvent))
				s.localMemberEvent(e.(serf.MemberEvent))
			case serf.EventMemberLeave, serf.EventMemberFailed:
				s.nodeLeave(e.(serf.MemberEvent))
				s.localMemberEvent(e.(serf.MemberEvent))
			}
		}
	}
}

func (s *Server) localMemberEvent(me serf.MemberEvent) {
	if s.Raft == nil {
		return
	}

	if !s.IsLeader() {
		return
	}

	for _, m := range me.Members {
		select {
		case s.reconcileCh <- m:
		default:
		}
	}
}

func (s *Server) nodeJoin(me serf.MemberEvent) {
	for _, m := range me.Members {
		s.logger.Printf("[INFO]: Member join: %s\n", m.Name)
	}

	if s.Raft != nil {
		if atomic.LoadInt64(&s.Config.BootstrapExpected) != 0 {
			s.tryBootstrap()
		}
	}
}

func (s *Server) nodeLeave(me serf.MemberEvent) {
	for _, m := range me.Members {
		s.logger.Printf("[INFO]: Member leave: %s\n", m.Name)
	}
}

func (s *Server) tryBootstrap() {
	servers := []raft.Server{}
	members := s.Serf.Members()
	for _, member := range members {
		id := member.Tags["id"]
		if id == "" {
			s.logger.Printf("Id not found for member addr: %s", member.Addr.String())
			continue
		}

		raftAddr := fmt.Sprintf("%s:%d", member.Addr.String(), member.Port)

		peer := raft.Server{
			ID:      raft.ServerID(id),
			Address: raft.ServerAddress(raftAddr),
		}
		servers = append(servers, peer)
	}

	if len(servers) < int(atomic.LoadInt64(&s.Config.BootstrapExpected)) {
		return
	}

	configuration := raft.Configuration{
		Servers: servers,
	}
	if err := s.Raft.BootstrapCluster(configuration).Error(); err != nil {
		s.logger.Printf("Failed to bootstrap cluster: %v\n", err)
	}

	atomic.StoreInt64(&s.Config.BootstrapExpected, 0)
}
