package models

import (
	"fmt"
	"net"
	"strconv"
)

type IPAddress struct {
	Address string `json:"address"`
	Mask    string `json:"mask"`
	CIDR    string `json:"-"`
}

func (ip *IPAddress) SetCIDR() error {
	prefix_length, _ := net.IPMask(net.ParseIP(ip.Mask).To4()).Size()
	if ip.Mask != "0.0.0.0" && prefix_length == 0 {
		return fmt.Errorf("expected %s to be a valid IPv4 netmask", ip.Mask)
	}
	cidr := ip.Address + "/" + strconv.Itoa(prefix_length)
	ip.CIDR = cidr

	return nil
}

func (ip *IPAddress) SetNetmask() error {
	ipv4Address, ipv4Net, err := net.ParseCIDR(ip.CIDR)
	if err != nil {
		return err
	}
	ip.Address = ipv4Address.String()
	ip.Mask = fmt.Sprintf("%d.%d.%d.%d", ipv4Net.Mask[0], ipv4Net.Mask[1], ipv4Net.Mask[2], ipv4Net.Mask[3])

	return nil
}


const BaseURI = "restconf/data/Cisco-IOS-XE-native:native"
var (
	BgpBaseURI = fmt.Sprintf("%s/%s", BaseURI, "router/Cisco-IOS-XE-bgp:bgp")
)

func GetBaseURI() string {
	return BaseURI
}

func GetBgpURI(asn string) string {
	return fmt.Sprintf("%s=%s", BgpBaseURI, asn)
}

func GetBgpNeighborURI(asn, neighbor string) string {
	return fmt.Sprintf("%s/neighbor=%s", string(GetBgpURI(asn)), neighbor)
}