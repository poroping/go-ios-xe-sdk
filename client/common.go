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

func GetBgpURI(asn int) string {
	return fmt.Sprintf("%s=%d", BgpBaseURI, asn)
}

func GetBgpNeighborURI(asn int, neighbor string) string {
	return fmt.Sprintf("%s/neighbor=%s", string(GetBgpURI(asn)), neighbor)
}

func GetBgpNeighborConfigURI(asn int, neighbor string, vrf *string) string {
	if vrf != nil {
		return fmt.Sprintf("%s/address-family/with-vrf/ipv4/unicast/vrf=%s/ipv4-unicast/neighbor=%s", string(GetBgpURI(asn)), *vrf, neighbor)
	}
	return fmt.Sprintf("%s/address-family/no-vrf/ipv4/unicast/ipv4-unicast/neighbor=%s", string(GetBgpURI(asn)), neighbor)
}
