package models

import (
	"encoding/json"
	"fmt"
)

func BgpNeighborPath(asn int, neighbor string) string {
	return fmt.Sprintf("%s/neighbor=%s", string(BgpPath(asn)), neighbor)
}

type BgpNeighbor struct {
	Neighbor Neighbor `json:"Cisco-IOS-XE-bgp:neighbor,omitempty"`
}

type Neighbor struct {
	ID                    string           `json:"id,omitempty"`
	RemoteAs              *int             `json:"remote-as,omitempty"`
	ClusterID             interface{}      `json:"cluster-id,omitempty"`
	Description           *string          `json:"description,omitempty"`
	DisableConnectedCheck *json.RawMessage `json:"disable-connected-check,omitempty"`
	EbgpMultihop          *struct {
		MaxHop *int `json:"max-hop,omitempty"`
	} `json:"ebgp-multihop,omitempty"`
	LocalAs *struct {
		AsNo *int `json:"as-no,omitempty"`
	} `json:"local-as,omitempty"`
	Shutdown     *json.RawMessage `json:"shutdown,omitempty"`
	Timers       *Timers          `json:"timers,omitempty"`
	UpdateSource *UpdateSource    `json:"update-source,omitempty"`
	ASN          int              `json:"-"`
}

type Timers struct {
	KeepaliveInterval   int `json:"keepalive-interval,omitempty"`
	Holdtime            int `json:"holdtime,omitempty"`
	MinimumNeighborHold int `json:"minimum-neighbor-hold,omitempty"`
}

type UpdateSource struct {
	Interface *Interface `json:"interface,omitempty"`
}
