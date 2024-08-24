package conn

import (
	"context"
	"crypto/tls"
	"fmt"
	"net"
	"time"

	"github.com/quic-go/quic-go"
	"github.com/spf13/viper"
)

func NewConn() {
	ctx := context.Background()
	if viper.GetBool("IsHub") {
		udpConn, err := net.ListenUDP("udp", &net.UDPAddr{Port: 62442})
		if err != nil {
			panic(fmt.Errorf("Fatal error creating hub connect: %w", err))
		}
		cfg := &quic.Config{
			HandshakeIdleTimeout: 20 * time.Second,
			MaxIdleTimeout: 90 * time.Second,
			Allow0RTT: true,
			DisablePathMTUDiscovery: false,
		}
		tr := quic.Transport {
			Conn: udpConn,
		}
		ln, err := tr.Listen(&tls.Config{}, cfg)
		if err != nil {
			panic(fmt.Errorf("Fatal error creating hub listener"))
		}
		for {
			conn, err := ln.Accept(ctx)
			if err != nil {
				panic(fmt.Errorf("Fatal error accepting connection"))
			}
			fmt.Printf("Remote Address: %s", conn.RemoteAddr())
		}
	} else {

	}
}
