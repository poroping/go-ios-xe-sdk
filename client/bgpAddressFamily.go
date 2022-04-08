package client

import (
	"fmt"
	"log"

	"github.com/poroping/go-ios-xe-sdk/models"
	"github.com/poroping/go-ios-xe-sdk/request"
)

func (c *CiscoIOSXEClient) CreateBgpAddressFamily(m models.BgpAddressFamily) error {
	// Crazy flow cause API sucks
	// PATCH new VRF to AF
	// https://{{host}}/restconf/data/Cisco-IOS-XE-native:native/router/bgp=65421/address-family/with-vrf/ipv4/unicast/vrf
	// {
	//     "Cisco-IOS-XE-bgp:vrf": [
	//         {
	//             "name": "BOOBS"
	//         }
	//     ]
	// }
	// PATCH error -> PUT to create
	// "patch to a nonexistent resource"
	if m.ASN == 0 || m.AddressFamilyType == "" {
		return fmt.Errorf("both ASN and AddressFamilyType must be set")
	}
	var rb []byte
	// if no vrf the address family is different cause cIScO
	if m.VRF == nil {
		rb = []byte(fmt.Sprintf(`{
			 	"Cisco-IOS-XE-bgp:%s": {
			 		"af-name": "unicast"
			 	}
			}`, m.AddressFamilyType))
		r := models.IOSXERequest{}
		r.HTTPMethod = "PUT"
		r.Path = models.BgpAddressFamilyBasePath(m)
		r.Payload = rb
		err := request.CreateUpdate(&c.Config, &r)
		if err != nil {
			return err
		}
		log.Printf("[INFO] address-family for no-vrf created.")

		return nil
	}

	rb = []byte(fmt.Sprintf(`{
		"Cisco-IOS-XE-bgp:vrf": [
			{
				"name": %q
			}
		]
	}`, *m.VRF))

	r := models.IOSXERequest{}
	r.HTTPMethod = "PATCH"
	r.Key = m.VRF
	r.Payload = rb
	r.Path = models.BgpAddressFamilyVRFBasePath(m)

	err := request.CreateUpdate(&c.Config, &r)
	if err != nil {
		if err.Error() == "patch to a nonexistent resource" {
			log.Printf("[WARN] Unable to add VRF to Address family VRF list. Will attempt to create from scratch.")
			r.HTTPMethod = "PUT"
			r.Path = models.BgpAddressFamilyBasePath(m)
			r.Payload = []byte(fmt.Sprintf(`{
				"Cisco-IOS-XE-bgp:%s": {
					"af-name": "unicast",
							"vrf": [
								{
									"name": %q
								}
					]
				}
			}`, m.AddressFamilyType, *m.VRF))
			err = request.CreateUpdate(&c.Config, &r)
			if err != nil {
				return err
			}
		}
		// return err
	}

	log.Printf("[INFO] address-family for vrf %v created.", *m.VRF)

	return nil
}
