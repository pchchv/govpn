package cipher

var _key []byte

func GenerateKey(key string) {
	_key = []byte(key)
}
