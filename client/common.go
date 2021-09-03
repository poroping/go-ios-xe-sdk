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