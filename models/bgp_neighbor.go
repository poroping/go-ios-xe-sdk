package models

type BgpNeighbor struct {
	Neighbor Neighbor `json:"Cisco-IOS-XE-bgp:neighbor,omitempty"`
}

type Neighbor struct {
	ID       string `json:"id,omitempty"`
	RemoteAs int    `json:"remote-as,omitempty"`
	ClusterID               string `json:"cluster-id,omitempty"`
	Description             string `json:"description,omitempty"`
	DisableConnectedCheck   string `json:"disable-connected-check,omitempty"`
	Shutdown string `json:"shutdown,omitempty"`
}