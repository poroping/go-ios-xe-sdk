package models

type SVIList struct {
	Vlan []L3Interface `json:"Cisco-IOS-XE-native:Vlan,omitempty"`
}

type SVI struct {
	Vlan L3Interface `json:"Cisco-IOS-XE-native:Vlan,omitempty"`
}

type L3Interface struct {
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
	Shutdown []string `json:"shutdown,omitempty"`
}
