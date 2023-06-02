package register

import (
	"log"
	"net"
	"strings"

	"github.com/patrickmn/go-cache"
)

var _register *cache.Cache

func AddClientIP(ip string) {
	_register.Add(ip, 0, cache.DefaultExpiration)
}

func ExistClientIP(ip string) bool {
	_, ok := _register.Get(ip)
	return ok
}

func DeleteClientIP(ip string) {
	_register.Delete(ip)
}

func PickClientIP(cidr string) (clientIP string, prefixLength string) {
	ip, ipNet, err := net.ParseCIDR(cidr)
	if err != nil {
		log.Panicf("error cidr %v", cidr)
	}
	total := addressCount(ipNet) - 3
	index := uint64(0)

	ip = incr(ipNet.IP.To4())

	for {
		ip = incr(ip)
		index++

		if index >= total {
			break
		}

		if !ExistClientIP(ip.String()) {
			AddClientIP(ip.String())
			return ip.String(), strings.Split(cidr, "/")[1]
		}
	}

	return "", ""
}

func KeepAliveClientIP(ip string) {
	if ExistClientIP(ip) {
		_register.Increment(ip, 1)
	} else {
		AddClientIP(ip)
	}
}

func ListClientIP() (result []string) {
	for k := range _register.Items() {
		result = append(result, k)
	}

	return result
}

func addressCount(network *net.IPNet) uint64 {
	prefixLen, bits := network.Mask.Size()
	return 1 << (uint64(bits) - uint64(prefixLen))
}

func checkIPv4(ip net.IP) net.IP {
	if v4 := ip.To4(); v4 != nil {
		return v4
	}

	return ip
}

func incr(IP net.IP) net.IP {
	IP = checkIPv4(IP)
	incIP := make([]byte, len(IP))

	copy(incIP, IP)

	for j := len(incIP) - 1; j >= 0; j-- {
		incIP[j]++
		if incIP[j] > 0 {
			break
		}
	}

	return incIP
}
