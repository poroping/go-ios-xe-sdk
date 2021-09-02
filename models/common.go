package models

import (
	"net"
	"strconv"
)

type IPAddress struct {
	Address string `json:"address"`
	Mask    string `json:"mask"`
	CIDR    string
}

func (ip *IPAddress) SetCIDR() {
	prefix_length, _ := net.IPMask(net.ParseIP(ip.Mask).To4()).Size()
	cidr := ip.Address + "/" + strconv.Itoa(prefix_length)
	ip.CIDR = cidr
}

func (ip *IPAddress) SetNetmask() error {
	ipv4Address, ipv4Net, err := net.ParseCIDR(ip.CIDR)
	if err != nil {
		return err
	}
	ip.Address = ipv4Address.String()
	ip.Mask = ipv4Net.Mask.String()

	return nil
}
