package models

import "encoding/json"

const VRFPath = "/restconf/data/Cisco-IOS-XE-native:native/vrf/definition"

type VRFDefinitionList struct {
	VRFDefinitionList []VRF `json:"Cisco-IOS-XE-native:definition,omitempty"`
}

type VRFDefinition struct {
	VRFDefinition VRF `json:"Cisco-IOS-XE-native:definition,omitempty"`
}

type VRF struct {
	Name          string              `json:"name,omitempty"`
	Description   *string             `json:"description,omitempty"`
	RD            *string             `json:"rd,omitempty"`
	AddressFamily *VRFAddressFamily   `json:"address-family,omitempty"`
	RouteTarget   *VRFRouteTargetList `json:"route-target,omitempty"`
}

type VRFAddressFamily struct {
	Ipv4 *VRFAddressFamilyIpv4 `json:"ipv4,omitempty"`
	Ipv6 *VRFAddressFamilyIpv6 `json:"ipv6,omitempty"`
}

type VRFAddressFamilyIpv4 struct {
	Maximum *VRFAddressFamilyIpv4Maximum `json:"maximum,omitempty"`
}

type VRFAddressFamilyIpv4Maximum struct {
	Routes      *int64           `json:"routes,omitempty"`
	WarningOnly *json.RawMessage `json:"warning-only,omitempty"`
}

type VRFAddressFamilyIpv6 struct {
}

type VRFRouteTargetList struct {
	Export []VRFRouteTarget `json:"export,omitempty"`
	Import []VRFRouteTarget `json:"import,omitempty"`
}

type VRFRouteTarget struct {
	AsnIP *string `json:"asn-ip,omitempty"`
}
