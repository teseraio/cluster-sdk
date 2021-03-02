package cluster

import (
	"io"
	"net"
	"os"
)

type Config struct {
	NodeName          string
	ServiceName       string
	Tags              map[string]string
	BootstrapExpected int64
	LogOutput         io.Writer
	GRPCAddr          *net.TCPAddr
}

func DefaultConfig() *Config {
	c := &Config{
		NodeName:          "cluster",
		ServiceName:       "cluster",
		Tags:              map[string]string{},
		BootstrapExpected: 0,
		LogOutput:         os.Stderr,
		GRPCAddr:          &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 5000},
	}
	return c
}
