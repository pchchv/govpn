package server

import (
	"fmt"
	"io"

	"github.com/gorilla/websocket"
	"github.com/patrickmn/go-cache"
	"github.com/pchchv/govpn/common/cipher"
	"github.com/pchchv/govpn/common/netutil"
	"github.com/songgao/water"
	"github.com/songgao/water/waterutil"
)

func vpnToWs(iface *water.Interface, c *cache.Cache) {
	buffer := make([]byte, 1500)

	for {
		n, err := iface.Read(buffer)
		if err != nil || err == io.EOF || n == 0 {
			continue
		}

		b := buffer[:n]
		if !waterutil.IsIPv4(b) {
			continue
		}

		srcAddr, dstAddr := netutil.GetAddr(b)
		if srcAddr == "" || dstAddr == "" {
			continue
		}

		key := fmt.Sprintf("%v->%v", dstAddr, srcAddr)
		v, ok := c.Get(key)
		if ok {
			b = cipher.XOR(b)
			v.(*websocket.Conn).WriteMessage(websocket.BinaryMessage, b)
		}
	}
}
