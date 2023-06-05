package client

import (
	"io"
	"time"

	"github.com/gorilla/websocket"
	"github.com/patrickmn/go-cache"
	"github.com/pchchv/govpn/common/cipher"
	"github.com/pchchv/govpn/common/netutil"
	"github.com/songgao/water"
	"github.com/songgao/water/waterutil"
)

func wsToVpn(c *cache.Cache, key string, wsConn *websocket.Conn, iface *water.Interface) {
	defer netutil.CloseWS(wsConn)

	for {
		wsConn.SetReadDeadline(time.Now().Add(time.Duration(30) * time.Second))
		_, b, err := wsConn.ReadMessage()
		if err != nil || err == io.EOF {
			break
		}

		b = cipher.XOR(b)
		if !waterutil.IsIPv4(b) {
			continue
		}

		iface.Write(b[:])
	}

	c.Delete(key)
}
