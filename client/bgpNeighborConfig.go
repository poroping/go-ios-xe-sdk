package client

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	"github.com/poroping/go-ios-xe-sdk/models"
	"github.com/poroping/go-ios-xe-sdk/request"
)

func (c *CiscoIOSXEClient) CreateBgpNeighborConfig(m models.BgpNeighborConfig) error {
	id := m.NeighborConfig.ID
	asn := m.NeighborConfig.ASN
	vrf := m.NeighborConfig.Vrf
	uri := models.BgpNeighborConfigPath(m)

	if vrf == nil {
		neighbor := models.BgpNeighbor{}
		neighbor.Neighbor.ID = id
		neighbor.Neighbor.ASN = asn
		parent_exists, err := c.ReadBgpNeighbor(neighbor)
		if parent_exists == nil {
			return fmt.Errorf("neighbor %q does not exist in default vrf: %s", id, err)
		}
	}

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	r.Key = &id
	r.Payload = rb
	r.Path = uri

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		// Some crazy flow cause API sucks
		// PUT neighbor to 'new' VRF
		// "error-message": "missing element: ipv4-unicast in ..."
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
		// finally, PUT neighbor config
		e := err.Error()
		if strings.HasPrefix(e, "missing element: ipv4-unicast") {
			log.Printf("[WARN] Unable to create Neighbor due to missing dependency. Attempting to resolve.")
			af := models.BgpAddressFamily{}
			af.ASN = asn
			af.AddressFamilyType = m.NeighborConfig.AddressFamilyType
			af.VRF = vrf
			err := c.CreateBgpAddressFamily(af)
			if err != nil {
				return err
			}
			err = request.CreateUpdate(&c.Config, &r)
			if err != nil {
				return err
			}
			return nil
		}
		return err
	}

	return nil
}

func (c *CiscoIOSXEClient) ReadBgpNeighborConfig(m models.BgpNeighborConfig) (*models.BgpNeighborConfig, error) {
	id := m.NeighborConfig.ID
	asn := m.NeighborConfig.ASN
	vrf := m.NeighborConfig.Vrf
	uri := models.BgpNeighborConfigPath(m)

	if vrf == nil {
		neighbor := models.BgpNeighbor{}
		neighbor.Neighbor.ID = id
		neighbor.Neighbor.ASN = asn
		parent_exists, err := c.ReadBgpNeighbor(neighbor)
		if parent_exists == nil {
			return nil, fmt.Errorf("neighbor %q does not exist in default vrf: %s", id, err)
		}
	}

	r := models.IOSXERequest{}
	r.HTTPMethod = "GET"
	r.Key = &id
	r.Path = uri

	body, err := request.Read(&c.Config, &r)
	if err != nil {
		return nil, err
	}

	res := models.BgpNeighborConfig{}
	if body == nil {
		return &res, nil
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *CiscoIOSXEClient) UpdateBgpNeighborConfig(m models.BgpNeighborConfig) error {
	id := m.NeighborConfig.ID
	asn := m.NeighborConfig.ASN
	vrf := m.NeighborConfig.Vrf
	uri := models.BgpNeighborConfigPath(m)

	if vrf == nil {
		neighbor := models.BgpNeighbor{}
		neighbor.Neighbor.ID = id
		neighbor.Neighbor.ASN = asn
		parent_exists, err := c.ReadBgpNeighbor(neighbor)
		if parent_exists == nil {
			return fmt.Errorf("neighbor %q does not exist in default vrf: %s", id, err)
		}
	}

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] %v", string(rb))

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	r.Key = &id
	r.Path = uri
	r.Payload = rb

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	return nil
}

func (c *CiscoIOSXEClient) DeleteBgpNeighborConfig(m models.BgpNeighborConfig) error {
	id := m.NeighborConfig.ID
	uri := models.BgpNeighborConfigPath(m)

	r := models.IOSXERequest{}
	r.HTTPMethod = "DELETE"
	r.Key = &id
	r.Path = uri

	err := request.Delete(&c.Config, &r)
	if err != nil {
		return err
	}

	return nil
}

// func (c *CiscoIOSXEClient) ListBgpNeighborConfig(as int) (*models.BgpNeighborConfigList, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/address-family/no-vrf/ipv4/unicast/neighbor/", c.HostURL, as), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, 200)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := models.BgpNeighborConfigList{}
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &res, nil
// }
