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
		go func() {
			for {
				fmt.Println("Here")
				conn, err := ln.Accept(ctx)
				if err != nil {
					panic(fmt.Errorf("Fatal error accepting connection"))
				}
				fmt.Println("hi")
				fmt.Printf("Remote Address: %s", conn.RemoteAddr())
			}
		}()
	}
	if true {
		addr, err := net.ResolveUDPAddr("udp", "0.0.0.0:6121")
		if err != nil {
			panic(fmt.Errorf("Fatal error resolving addr: %w", err))
		}
		c, err := net.ListenUDP("udp", addr)
		if err != nil {
			panic(fmt.Errorf("Fatal error creating hub lisetn connect: %w", err))
		}
		cfg := &quic.Config{
			HandshakeIdleTimeout: 20 * time.Second,
			MaxIdleTimeout: 90 * time.Second,
			Allow0RTT: true,
			DisablePathMTUDiscovery: false,
		}
		tr := quic.Transport {
			Conn: c,
		}
		ctx := context.Background()
		conn, err := tr.Dial(ctx, addr, &tls.Config{}, cfg)
		if err != nil {
			panic(fmt.Errorf("Fatal error quic dialing: %w", err))
		}
		fmt.Printf("Remote Address: %s", conn.RemoteAddr())
		stream, err := conn.OpenStreamSync(ctx)
		if err != nil {
			panic(fmt.Errorf("Fatal error quic stream: %w", err))
		}
		stream.Write([]byte("hello"))
	}
}
