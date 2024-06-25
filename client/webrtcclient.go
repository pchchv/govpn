package client

import (
	"fmt"
	"time"

	"github.com/pchchv/govpn/common/cipher"
	"github.com/pchchv/govpn/common/config"
	"github.com/pchchv/govpn/common/sdputil"
	"github.com/pion/webrtc/v3"
)
func createConnection() (*webrtc.PeerConnection, error) {
	config := webrtc.Configuration{
		ICEServers: []webrtc.ICEServer{
			{
				URLs: []string{"stun:stun.l.google.com:19302"},
			},
		},
	}
	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		return nil, err
	}

	return peerConnection, nil
}

func StartWebRTCClient(config config.Config) {
	peerConnection, err := createConnection()
	if err != nil {
		panic(err)
	}

	peerConnection.OnICEConnectionStateChange(func(state webrtc.ICEConnectionState) {
		if state == webrtc.ICEConnectionStateFailed {
			fmt.Printf("\nICE Connection State has changed: %s\n\n", state.String())
			return
		}

		if state != webrtc.ICEConnectionStateClosed {
			fmt.Printf("\nICE Connection State has changed: %s\n\n", state.String())
		}
	})

	ordered := true
	mplt := uint16(5000)
	channel, err := peerConnection.CreateDataChannel("control", &webrtc.DataChannelInit{
		Ordered:           &ordered,
		MaxPacketLifeTime: &mplt,
	})
	if err != nil {
		panic(err)
	}

	answer, err := sdputil.SDPPrompt()
	if err != nil {
		panic(err)
	}

	err = PrintSDP(peerConnection, answer)
	if err != nil {
		panic(err)
	}
	
	// iface := vpn.CreateVpn(config.CIDR)

	channel.OnMessage(func(msg webrtc.DataChannelMessage) {
		// relay packets
		// b := cipher.XOR(msg.Data)
		// if !waterutil.IsIPv4(b) {
		// 	return
		// }

		// iface.Write(b)

		println(string(msg.Data))
	})

	//packet := make([]byte, 1500)
	for {
		// n, err := iface.Read(packet)
		// if err != nil || n == 0 {
		// 	continue
		// }
		// if !waterutil.IsIPv4(packet) {
		// 	continue
		// }
		time.Sleep(5 * time.Second)

		//b := cipher.XOR([]byte("HELLO"))
		err = channel.Send([]byte("Helllo"))
		if err != nil {
			println(err.Error())
			continue
		}
		println("sent hello packet")
	}
}

func PrintSDP(p *webrtc.PeerConnection, offer webrtc.SessionDescription) error {
	sdp, err := GenSDP(p, offer)
	if err != nil {
		return err
	}
	fmt.Println(sdp)
	return nil
}

func GenSDP(p *webrtc.PeerConnection, offer webrtc.SessionDescription) (string, error) {
	var sdp string
	err := p.SetRemoteDescription(offer)
	if err != nil {
		return sdp, err
	}

	answer, err := p.CreateAnswer(nil)
	if err != nil {
		return sdp, err
	}

	gatherDone := webrtc.GatheringCompletePromise(p)
	err = p.SetLocalDescription(answer)
	if err != nil {
		return sdp, err
	}
	<-gatherDone

	//Encode the SDP to base64
	sdp, err = cipher.Encode(p.LocalDescription())
	return sdp, err
}