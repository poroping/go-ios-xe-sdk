package client

import (
	"fmt"
)

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

func GetBgpNeighborConfigURI(asn, neighbor string, vrf *string) string {
	if vrf == nil {
		noVrf := "no-vrf"
		vrf = &noVrf
	}
	return fmt.Sprintf("%s/address-family/%s/ipv4/unicast/ipv4-unicast/neighbor=%s", string(GetBgpURI(asn)), &vrf, neighbor)
}
