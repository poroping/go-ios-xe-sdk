package models

type Vlan struct {
	VlanList VlanList `json:"Cisco-IOS-XE-vlan:vlan-list"`
}

type VlanList struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
