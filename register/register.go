package register

import "net"

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
