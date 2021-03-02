package cluster

import (
	"fmt"
	"net"
	"sync/atomic"
	"time"

	sockaddr "github.com/hashicorp/go-sockaddr"
	"github.com/hashicorp/memberlist"
	"github.com/teseraio/cluster-sdk/cluster/proto"
)

// Based on https://github.com/hashicorp/memberlist/blob/master/net_transport.go

const (
	udpPacketBufSize = 65536
	udpRecvBufSize   = 2 * 1024 * 1024
)

type streamTransportConfig struct {
	BindAddr string
	BindPort int
}

type streamTransport struct {
	server   *Server
	config   *streamTransportConfig
	packetCh chan *memberlist.Packet
	streamCh chan net.Conn

	tcpListener *net.TCPListener
	udpListener *net.UDPConn

	shutdown int32
}

func newStreamTransport(config *streamTransportConfig) (*streamTransport, error) {
	t := streamTransport{
		config:   config,
		packetCh: make(chan *memberlist.Packet),
		streamCh: make(chan net.Conn),
	}

	port := config.BindPort
	ip := net.ParseIP(config.BindAddr)

	udpAddr := &net.UDPAddr{IP: ip, Port: port}
	udpLn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		return nil, fmt.Errorf("Failed to start UDP listener on %q port %d: %v", udpAddr, port, err)
	}
	if err := setUDPRecvBuf(udpLn); err != nil {
		return nil, fmt.Errorf("Failed to resize UDP buffer: %v", err)
	}
	t.udpListener = udpLn

	go t.udpListen(t.udpListener)
	return &t, nil
}

// GetAutoBindPort implements the Transport interface
func (t *streamTransport) GetAutoBindPort() int {
	return t.tcpListener.Addr().(*net.TCPAddr).Port
}

// FinalAdvertiseAddr implements the Transport interface
func (t *streamTransport) FinalAdvertiseAddr(ip string, port int) (net.IP, int, error) {
	var advertiseAddr net.IP
	var advertisePort int
	if ip != "" {
		// If they've supplied an address, use that.
		advertiseAddr = net.ParseIP(ip)
		if advertiseAddr == nil {
			return nil, 0, fmt.Errorf("Failed to parse advertise address %q", ip)
		}

		// Ensure IPv4 conversion if necessary.
		if ip4 := advertiseAddr.To4(); ip4 != nil {
			advertiseAddr = ip4
		}
		advertisePort = port
	} else {
		if t.config.BindAddr == "0.0.0.0" {
			// Otherwise, if we're not bound to a specific IP, let's
			// use a suitable private IP address.
			var err error
			ip, err = sockaddr.GetPrivateIP()
			if err != nil {
				return nil, 0, fmt.Errorf("Failed to get interface addresses: %v", err)
			}
			if ip == "" {
				return nil, 0, fmt.Errorf("No private IP address found, and explicit IP not provided")
			}

			advertiseAddr = net.ParseIP(ip)
			if advertiseAddr == nil {
				return nil, 0, fmt.Errorf("Failed to parse advertise address: %q", ip)
			}
		} else {
			// Use the IP that we're bound to, based on the first
			// TCP listener, which we already ensure is there.
			advertiseAddr = t.tcpListener.Addr().(*net.TCPAddr).IP
		}

		// Use the port we are bound to.
		advertisePort = t.GetAutoBindPort()
	}

	return advertiseAddr, advertisePort, nil
}

// WriteTo implements the Transport interface
func (t *streamTransport) WriteTo(b []byte, addr string) (time.Time, error) {
	udpAddr, err := net.ResolveUDPAddr("udp", addr)
	if err != nil {
		return time.Time{}, err
	}
	_, err = t.udpListener.WriteTo(b, udpAddr)
	return time.Now(), err
}

// PacketCh implements the Transport interface
func (t *streamTransport) PacketCh() <-chan *memberlist.Packet {
	return t.packetCh
}

// DialTimeout implements the Transport interface
func (t *streamTransport) DialTimeout(addr string, timeout time.Duration) (net.Conn, error) {
	return t.server.createStream(addr, timeout, proto.StreamObj_Start_SERF)
}

// StreamCh implements the Transport interface
func (t *streamTransport) StreamCh() <-chan net.Conn {
	return t.server.streams[proto.StreamObj_Start_SERF]
}

// Shutdown implements the Transport interface
func (t *streamTransport) Shutdown() error {
	atomic.StoreInt32(&t.shutdown, 1)
	t.udpListener.Close()
	return nil
}

func (t *streamTransport) udpListen(udpLn *net.UDPConn) {
	for {
		buf := make([]byte, udpPacketBufSize)
		n, addr, err := udpLn.ReadFrom(buf)
		ts := time.Now()
		if err != nil {
			if s := atomic.LoadInt32(&t.shutdown); s == 1 {
				break
			}
			continue
		}

		if n < 1 {
			continue
		}
		t.packetCh <- &memberlist.Packet{
			Buf:       buf[:n],
			From:      addr,
			Timestamp: ts,
		}
	}
}

func setUDPRecvBuf(c *net.UDPConn) error {
	size := udpRecvBufSize
	var err error
	for size > 0 {
		if err = c.SetReadBuffer(size); err == nil {
			return nil
		}
		size = size / 2
	}
	return err
}
