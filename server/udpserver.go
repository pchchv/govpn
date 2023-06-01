package server

import (
	"fmt"
	"net"

	"github.com/patrickmn/go-cache"
	"github.com/pchchv/govpn/common/cipher"
	"github.com/pchchv/govpn/common/netutil"
	"github.com/songgao/water"
	"github.com/songgao/water/waterutil"
)

type Forwarder struct {
	localConn *net.UDPConn
	connCache *cache.Cache
}

func (f *Forwarder) forward(iface *water.Interface, conn *net.UDPConn) {
	packet := make([]byte, 1500)

	for {
		n, err := iface.Read(packet)
		if err != nil || n == 0 {
			continue
		}
		b := packet[:n]
		if !waterutil.IsIPv4(b) {
			continue
		}
		srcAddr, dstAddr := netutil.GetAddr(b)
		if srcAddr == "" || dstAddr == "" {
			continue
		}
		key := fmt.Sprintf("%v->%v", dstAddr, srcAddr)
		v, ok := f.connCache.Get(key)
		if ok {
			b = cipher.XOR(b)
			f.localConn.WriteToUDP(b, v.(*net.UDPAddr))
		}
	}
}
