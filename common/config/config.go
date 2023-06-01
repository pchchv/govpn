package config

type Config struct {
	Key        string
	CIDR       string
	Protocol   string
	LocalAddr  string
	ServerAddr string
	ServerMode bool
}
