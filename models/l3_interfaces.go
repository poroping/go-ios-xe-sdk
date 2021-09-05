package models

import "encoding/json"

const SVIName string = "Vlan"

type SVIList struct {
	Vlan []L3Interface `json:"Cisco-IOS-XE-native:Vlan,omitempty"`
}

type SVI struct {
	Vlan L3Interface `json:"Cisco-IOS-XE-native:Vlan,omitempty"`
}

type L3Interface struct {
	Name        interface{}      `json:"name,omitempty"`
	Description *string          `json:"description,omitempty"`
	IP          L3IP             `json:"ip"`
	Shutdown    *json.RawMessage `json:"shutdown,omitempty"`
}

type L3IP struct {
	Vrf struct {
	} `json:"vrf,omitempty"`
	Address struct {
		Primary   IPAddress             `json:"primary,omitempty"`
		Secondary *[]SecondaryIPAddress `json:"secondary,omitempty"`
	} `json:"address,omitempty"`
}
