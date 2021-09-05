package models

import "encoding/json"

type BgpNeighborConfig struct {
	NeighborConfig NeighborConfig `json:"Cisco-IOS-XE-bgp:neighbor"`
}

type BgpNeighborConfigList struct {
	Collection struct {
		NeighborConfig []NeighborConfig `json:"Cisco-IOS-XE-bgp:neighbor"`
	} `json:"collection"`
}

type NeighborConfig struct {
	ID               string           `json:"id,omitempty"`
	Activate         *json.RawMessage `json:"activate,omitempty"`
	DefaultOriginate *struct{}        `json:"default-originate,omitempty"`
	LocalAs          *struct {
		AsNo int `json:"as-no,omitempty"`
	} `json:"local-as,omitempty"`
	PrefixList          []PrefixList `json:"prefix-list,omitempty"`
	RemoteAs            *int         `json:"remote-as,omitempty"`
	RemovePrivateAs     *struct{}    `json:"remove-private-as,omitempty"`
	SoftReconfiguration *string      `json:"soft-reconfiguration,omitempty"`
}
