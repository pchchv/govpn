package vpn

import (
	"log"

	"github.com/pchchv/govpn/common/osutil"
	"github.com/songgao/water"
)

func CreateServerVpn(cidr string) (iface *water.Interface) {
	c := water.Config{DeviceType: water.TAP}
	iface, err := water.New(c)
	if err != nil {
		log.Fatalln("failed to allocate vpn interface:", err)
	}

	log.Println("interface allocated:", iface.Name())

	osutil.ConfigVpnServer(cidr, iface)

	return iface
}

func CreateClientVpn(cidr string) (iface *water.Interface) {
	c := water.Config{DeviceType: water.TAP}
	iface, err := water.New(c)
	if err != nil {
		log.Fatalln("failed to allocate vpn interface:", err)
	}

	log.Println("interface allocated:", iface.Name())

	osutil.ConfigVpnClient(cidr, iface)

	return iface
}

