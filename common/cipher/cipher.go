package cipher

import (
	"crypto/rc4"
	"log"
)

var _key []byte

func GenerateKey(key string) {
	_key = []byte(key)
}

func XOR(src []byte) []byte {
	c, err := rc4.NewCipher(_key)
	if err != nil {
		log.Fatalln(err)
	}

	dst := make([]byte, len(src))
	c.XORKeyStream(dst, src)

	return dst
}
