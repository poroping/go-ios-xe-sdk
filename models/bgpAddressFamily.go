package models

import (
	"fmt"
	"log"

	"github.com/poroping/go-ios-xe-sdk/utils"
)

var (
	BgpAddressFamilyTypes = []string{"ipv4", "ipv6"}
)

func BgpAddressFamilyBasePath(m BgpAddressFamily) string {
	AFType := m.AddressFamilyType
	// restrict cause slapping into URL path
	// TODO: deal with this gracefully.
	if !utils.StringInSlice(AFType, BgpAddressFamilyTypes) {
		log.Fatalf("[FATAL] Address family type %q not valid", AFType)
	}
	if m.VRF != nil {
		return fmt.Sprintf("%s/address-family/with-vrf/%s=unicast", string(BgpPath(m.ASN)), AFType)
	}
	return fmt.Sprintf("%s/address-family/no-vrf/%s=unicast", string(BgpPath(m.ASN)), AFType)

	// return fmt.Sprintf("%s/address-family/no-vrf/ipv4/unicast", string(BgpPath(m.ASN)))
}

func BgpAddressFamilyVRFBasePath(m BgpAddressFamily) string {
	AFType := m.AddressFamilyType
	// restrict cause slapping into URL path
	// TODO: deal with this gracefully.
	if !utils.StringInSlice(AFType, BgpAddressFamilyTypes) {
		log.Fatalf("[FATAL] Address family type %q not valid", AFType)
	}
	if m.VRF == nil {
		log.Fatalf("[FATAL] VRF must be defined.")
	}
	return fmt.Sprintf("%s/address-family/with-vrf/%s/unicast/vrf", string(BgpPath(m.ASN)), AFType)

	// return fmt.Sprintf("%s/address-family/no-vrf/ipv4/unicast", string(BgpPath(m.ASN)))
}

// func BgpAddressFamilyPath(m BgpAddressFamily) string {
// 	AFType := m.AddressFamilyType
// 	// restrict cause slapping into URL path
// 	// TODO: deal with this gracefully.
// 	if AFType != "ipv4" || AFType != "ipv6" {
// 		log.Fatalf("[FATAL] Address family type %q not valid", AFType)
// 	}
// 	if m.VRF != nil {
// 		return fmt.Sprintf("%s/address-family/with-vrf/%s/unicast/vrf=%s", AFType, string(BgpPath(m.ASN)), *m.VRF)
// 	}
// 	return fmt.Sprintf("%s/address-family/no-vrf/%s/unicast", AFType, string(BgpPath(m.ASN)))
// }

type BgpAddressFamily struct {
	VRF               *string
	ASN               int
	AddressFamilyType string //ipv4 ipv6
}
