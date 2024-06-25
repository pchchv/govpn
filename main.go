package main

import (
	"context"
	"flag"
	"os"
	"os/signal"

	"github.com/pchchv/govpn/client"
	"github.com/pchchv/govpn/common/config"
	"github.com/pchchv/govpn/server"
)

func main() {
	config := config.Config{}

	flag.StringVar(&config.CIDR, "c", "172.16.0.1/24", "vpn interface CIDR")
	flag.StringVar(&config.LocalAddr, "l", "0.0.0.0:3000", "local address")
	flag.StringVar(&config.ServerAddr, "s", "0.0.0.0:3001", "server address")
	flag.StringVar(&config.Key, "k", "6w9z$C&F)J@NcRfWjXn3r4u7x!A%D*G-", "encryption key")
	flag.StringVar(&config.Protocol, "p", "wss", "protocol ws/wss/udp/rtc")
	flag.BoolVar(&config.ServerMode, "S", false, "server mode")
	flag.Parse()

	config.Init()
	
	switch config.Protocol {
	case "udp":
		if config.ServerMode {
			server.StartUDPServer(config)
		} else {
			client.StartUDPClient(config)
		}
	case "ws":
		if config.ServerMode {
			server.StartWSServer(config)
		} else {
			client.StartWSClient(config)
		}
	case "rtc":
		if config.ServerMode {
			ctx, cancel := context.WithCancel(context.Background())
			server.StartWebRTCServer(ctx, config)

			c := make(chan os.Signal, 1)
			signal.Notify(c, os.Interrupt)
			<-c
			cancel()
			return
		} else {
			client.StartWebRTCClient(config)
		}
	default:
	}
}
