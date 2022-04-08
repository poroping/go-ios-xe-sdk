package models

const L2VlanPath = "/restconf/data/Cisco-IOS-XE-native:native/vlan/vlan-list"

type L2VlanList struct {
	VlanList L2Vlan `json:"Cisco-IOS-XE-vlan:vlan-list,omitempty"`
}

type L2Vlan struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
