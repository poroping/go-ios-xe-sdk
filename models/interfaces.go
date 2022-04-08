package models

import (
	"encoding/json"

	"github.com/poroping/go-ios-xe-sdk/utils"
)

// "GigabitEthernet"
// TwentyFiveGigE https://{{host}}/restconf/data/Cisco-IOS-XE-native:native/interface/TwentyFiveGigE=1%2F0%2F1?with-defaults=report-all
// HundredGigE
// Loopback
// Port-channel https://{{host}}/restconf/data/Cisco-IOS-XE-native:native/interface/Port-channel=69
// VirtualPortGroup
// Port-channel-subinterface https://{{host}}/restconf/data/Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel=69.420

const (
	VlanPath                      = "/restconf/data/Cisco-IOS-XE-native:native/interface/Vlan"
	LoopbackPath                  = "/restconf/data/Cisco-IOS-XE-native:native/interface/Loopback"
	GigabitEthernetPath           = "/restconf/data/Cisco-IOS-XE-native:native/interface/GigabitEthernet"
	HundredGigabitEthernetPath    = "/restconf/data/Cisco-IOS-XE-native:native/interface/HundredGigE"
	TwentyFiveGigabitEthernetPath = "/restconf/data/Cisco-IOS-XE-native:native/interface/TwentyFiveGigE"
	PortChannelPath               = "/restconf/data/Cisco-IOS-XE-native:native/interface/Port-channel"
	PortChannelSubinterfacePath   = "/restconf/data/Cisco-IOS-XE-native:native/interface/Port-channel-subinterface/Port-channel"
)

const (
	VlanName = "Vlan"
)

type Interfaces struct {
	GigabitEthernet         []Interface `json:"GigabitEthernet,omitempty"`
	HundredGigE             []Interface `json:"HundredGigE,omitempty"`
	Loopback                []Interface `json:"Loopback,omitempty"`
	TwentyFiveGigE          []Interface `json:"TwentyFiveGigE,omitempty"`
	Vlan                    []Interface `json:"Vlan,omitempty"`
	PortChannel             []Interface `json:"Port-channel,omitempty"`
	PortChannelSubinterface []Interface `json:"Port-channel-subinterface,omitempty"`
	VirtualPortGroup        []Interface `json:"VirtualPortGroup,omitempty"`
}

type Vlan struct {
	Vlan Interface `json:"Cisco-IOS-XE-native:Vlan,omitempty"`
}

type Loopback struct {
	Loopback Interface `json:"Cisco-IOS-XE-native:Loopback,omitempty"`
}

type GigabitEthernet struct {
	GigabitEthernet Interface `json:"Cisco-IOS-XE-native:GigabitEthernet,omitempty"`
}

type TwentyFiveGigabitEthernet struct {
	TwentyFiveGigabitEthernet Interface `json:"Cisco-IOS-XE-native:TwentyFiveGigE,omitempty"`
}

type HundredGigabitEthernet struct {
	HundredGigabitEthernet Interface `json:"Cisco-IOS-XE-native:HundredGigE,omitempty"`
}

type PortChannel struct {
	PortChannel Interface `json:"Cisco-IOS-XE-native:Port-channel,omitempty"`
}

type PortChannelSubinterface struct {
	PortChannelSubinterface Interface `json:"Cisco-IOS-XE-native:Port-channel,omitempty"`
}

type Interface struct {
	Name          string                  `json:"name,omitempty"`
	ChannelGroup  *InterfaceChannelGroup  `json:"Cisco-IOS-XE-ethernet:channel-group,omitempty"`
	Description   *string                 `json:"description,omitempty"`
	Encapsulation *InterfaceEncapsulation `json:"encapsulation,omitempty"`
	IP            *InterfaceIP            `json:"ip,omitempty"`
	Shutdown      *json.RawMessage        `json:"shutdown,omitempty"`
	Vrf           *InterfaceVrf           `json:"vrf,omitempty"`
}

func (d *Interface) UnmarshalJSON(data []byte) error {
	type InterfaceAlias Interface
	aux := &struct {
		*InterfaceAlias
		Name interface{} `json:"name,omitempty"`
	}{
		InterfaceAlias: (*InterfaceAlias)(d),
	}

	if err := json.Unmarshal(data, aux); err != nil {
		return err
	}

	name := utils.ForceString(aux.Name)
	if name == nil {
		return nil // name should always be set
	}
	d.Name = *name

	return nil
}

type InterfaceIP struct {
	Address *Address `json:"address,omitempty"`
}

type InterfaceEncapsulation struct {
	Dot1Q *InterfaceEncapsulationDot1Q `json:"dot1Q,omitempty"`
}

type InterfaceEncapsulationDot1Q struct {
	VlanID *int64 `json:"vlan-id,omitempty"`
}

type InterfaceVrf struct {
	Forwarding *string `json:"forwarding,omitempty"`
}

type InterfaceChannelGroup struct {
	Number *int    `json:"number,omitempty"`
	Mode   *string `json:"mode,omitempty"`
}
