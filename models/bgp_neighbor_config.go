package models

type BgpNeighborConfig struct {
	NeighborConfig NeighborConfig `json:"ned:neighbor"`
}

type BgpNeighborConfigList struct {
	Collection struct {
		NeighborConfig []NeighborConfig `json:"ned:neighbor"`
	} `json:"collection"`
}

type NeighborConfig struct {
	ID               string        `json:"id,omitempty"`
	Activate         []interface{} `json:"activate,omitempty"`
	DefaultOriginate struct {
	} `json:"default-originate,omitempty"`
	LocalAs struct {
		AsNo int `json:"as-no,omitempty"`
	} `json:"local-as,omitempty"`
	PrefixList          []PrefixList  `json:"prefix-list,omitempty"`
	RemovePrivateAs     []interface{} `json:"remove-private-as,omitempty"`
	SoftReconfiguration string        `json:"soft-reconfiguration,omitempty"`
}
