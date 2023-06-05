# govpn [![Go Report Card](https://goreportcard.com/badge/github.com/pchchv/govpn)](https://goreportcard.com/report/github.com/pchchv/govpn)

A simple VPN client built in Go

# Installation
```console
$ git clone https://github.com/pchchv/govpn
```

# Build:
```console
$ bash scripts/build.sh
```

# Server:
```console
sudo ./main -S -l=:3001 -c=172.16.0.1/24 -k=123456
```

# Client:
```console
sudo ./main -l=:3000 -s=server-addr:3001 -c=172.16.0.10/24 -k=123456
```

# Server Setup:

* Add TLS for websocket, reverse proxy server (3001) via nginx/caddy(443)

* Enable IP forwarding on the server

```console
  sudo echo 1 > /proc/sys/net/ipv4/ip_forward
  sudo sysctl -p
  sudo iptables -t nat -A POSTROUTING -s 172.16.0.0/24 -o ens3 -j MASQUERADE
  sudo apt-get install iptables-persistent
  sudo iptables-save > /etc/iptables/rules.v4
```