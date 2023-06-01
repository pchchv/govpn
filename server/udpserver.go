package server

import (
	"net"

	"github.com/patrickmn/go-cache"
)

type Forwarder struct {
	localConn *net.UDPConn
	connCache *cache.Cache
}
