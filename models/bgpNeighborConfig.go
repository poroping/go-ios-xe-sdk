package models

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/poroping/go-ios-xe-sdk/utils"
)

func BgpNeighborConfigPath(m BgpNeighborConfig) string {
	if m.NeighborConfig.ASN == 0 {
		log.Printf("[WARN] ASN not defined.")
	}
	// restrict cause slapping into URL path
	// TODO: deal with this gracefully.
	if !utils.StringInSlice(m.NeighborConfig.AddressFamilyType, BgpAddressFamilyTypes) {
		log.Fatalf("[FATAL] Address family type %q not valid", m.NeighborConfig.AddressFamilyType)
	}
	if m.NeighborConfig.Vrf != nil {
		return fmt.Sprintf("%s/address-family/with-vrf/%s/unicast/vrf=%s/ipv4-unicast/neighbor=%s", string(BgpPath(m.NeighborConfig.ASN)), m.NeighborConfig.AddressFamilyType, *m.NeighborConfig.Vrf, m.NeighborConfig.ID)
	}
	return fmt.Sprintf("%s/address-family/no-vrf/%s/unicast/ipv4-unicast/neighbor=%s", string(BgpPath(m.NeighborConfig.ASN)), m.NeighborConfig.AddressFamilyType, m.NeighborConfig.ID)
}

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
	Vrf                 *string      `json:"-"`
	ASN                 int          `json:"-"`
	AddressFamilyType   string       `json:"-"`
}
