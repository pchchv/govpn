package sdputil

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	"github.com/pchchv/govpn/common/cipher"
	"github.com/pion/webrtc/v3"
)

func SDPPrompt() (webrtc.SessionDescription, error) {
	fmt.Println("Paste the remote SDP: ")

	//take remote SDP in answer
	answer := webrtc.SessionDescription{}
	for {
		r := bufio.NewReader(os.Stdin)
		var in string
		for {
			var err error
			in, err = r.ReadString('\n')
			if err != io.EOF {
				if err != nil {
					return webrtc.SessionDescription{}, err
				}
			}
			in = strings.TrimSpace(in)
			if len(in) > 0 {
				break
			}
		}
	
		fmt.Println("")

		if err := cipher.Decode(in, &answer); err == nil {
			break
		}
		fmt.Println("Invalid SDP. Enter again.")
	}
	return answer, nil
}
