package models

type L2Vlan struct {
	VlanList VlanList `json:"ned:vlan-list"`
}

type L2VlanList struct {
	Collection struct {
		VlanList []VlanList `json:"ned:vlan-list"`
	} `json:"collection"`
}

type VlanList struct {
	ID   int    `json:"id,omitempty"`
	Name string `json:"name,omitempty"`
}
