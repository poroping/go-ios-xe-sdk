package models

type BgpNeighborPrefixListList struct {
	Collection struct {
		PrefixList []PrefixList `json:"ned:prefix-list"`
	} `json:"collection"`
}

type BgpNeighborPrefixList struct {
	PrefixList PrefixList `json:"ned:prefix-list"`
}

type PrefixList struct {
	Inout          string `json:"inout"`
	PrefixListName string `json:"prefix-list-name"`
}
