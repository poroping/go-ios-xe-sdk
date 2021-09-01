package models

type L3VlanList struct {
	Collection struct {
		Vlan []Vlan `json:"Cisco-IOS-XE-native:Vlan"`
	} `json:"collection"`
}

type L3Vlan struct {
	Vlan Vlan `json:"Cisco-IOS-XE-native:Vlan"`
}

type Vlan struct {
	Name        int    `json:"name,omitempty"`
	Description string `json:"description,omitempty"`
	IP          struct {
		Vrf struct {
		} `json:"vrf,omitempty"`
		Address struct {
			Primary struct {
				Address string `json:"address"`
				Mask    string `json:"mask"`
			} `json:"primary,omitempty"`
		} `json:"address,omitempty"`
	} `json:"ip"`
}
