package models

import (
	"encoding/json"
	"fmt"
	"net"
	"strconv"
)

const ExplicitNull string = (`[null]`)

var CiscoEnabled json.RawMessage = []byte(ExplicitNull)

const BasePath = "/restconf/data/Cisco-IOS-XE-native:native"

type IOSXERequest struct {
	HTTPMethod string
	Payload    []byte
	Params     []string
	Path       string
	Key        *string
}

type Address struct {
	Primary   *IPAddress            `json:"primary,omitempty"`
	Secondary *[]SecondaryIPAddress `json:"secondary,omitempty"`
}

type IPAddress struct {
	Address string `json:"address,omitempty"`
	Mask    string `json:"mask,omitempty"`
	CIDR    string `json:"-"`
}

type SecondaryIPAddress struct {
	Address   string           `json:"address,omitempty"`
	Mask      string           `json:"mask,omitempty"`
	Secondary *json.RawMessage `json:"secondary,omitempty"`
	CIDR      string           `json:"-"`
}

func (ip *IPAddress) SetCIDR() error {
	prefix_length, _ := net.IPMask(net.ParseIP(ip.Mask).To4()).Size()
	if ip.Mask != "0.0.0.0" && prefix_length == 0 {
		return fmt.Errorf("expected %s to be a valid IPv4 netmask", ip.Mask)
	}
	cidr := ip.Address + "/" + strconv.Itoa(prefix_length)
	if cidr == "/" {
		return nil
	}
	ip.CIDR = cidr

	return nil
}

func (ip *IPAddress) SetNetmask() error {
	ipv4Address, ipv4Net, err := net.ParseCIDR(ip.CIDR)
	if err != nil {
		return err
	}
	ip.Address = ipv4Address.String()
	ip.Mask = fmt.Sprintf("%d.%d.%d.%d", ipv4Net.Mask[0], ipv4Net.Mask[1], ipv4Net.Mask[2], ipv4Net.Mask[3])

	return nil
}

func (ip *SecondaryIPAddress) SetCIDR() error {
	prefix_length, _ := net.IPMask(net.ParseIP(ip.Mask).To4()).Size()
	if ip.Mask != "0.0.0.0" && prefix_length == 0 {
		return fmt.Errorf("expected %s to be a valid IPv4 netmask", ip.Mask)
	}
	cidr := ip.Address + "/" + strconv.Itoa(prefix_length)
	if cidr == "/" {
		return nil
	}
	ip.CIDR = cidr

	return nil
}

func (ip *SecondaryIPAddress) SetNetmask() error {
	ipv4Address, ipv4Net, err := net.ParseCIDR(ip.CIDR)
	if err != nil {
		return err
	}
	ip.Address = ipv4Address.String()
	ip.Mask = fmt.Sprintf("%d.%d.%d.%d", ipv4Net.Mask[0], ipv4Net.Mask[1], ipv4Net.Mask[2], ipv4Net.Mask[3])

	return nil
}

type PrefixList struct {
	Inout          string `json:"inout,omitempty"`
	PrefixListName string `json:"prefix-list-name,omitempty"`
}
