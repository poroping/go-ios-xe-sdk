package client

import (
	"encoding/json"
	"log"

	"github.com/poroping/go-ios-xe-sdk/models"
	"github.com/poroping/go-ios-xe-sdk/request"
)

func (c *CiscoIOSXEClient) CreateBgpNeighbor(m models.BgpNeighbor) error {
	id := m.Neighbor.ID
	asn := m.Neighbor.ASN
	uri := models.BgpNeighborPath(asn, id)

	// disabled cause can't PATCH certain fields "off"
	// exists, _ := c.ReadBgpNeighbor(asn, m)
	// if exists != nil {
	// 	return c.UpdateBgpNeighbor(asn, m)
	// }

	// log.Printf("Doesn't exist, will create")

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	r.Key = &m.Neighbor.ID
	r.Payload = rb
	r.Path = uri

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Println("[INFO] neighbor created.")

	return nil
}

func (c *CiscoIOSXEClient) ReadBgpNeighbor(m models.BgpNeighbor) (*models.BgpNeighbor, error) {
	id := m.Neighbor.ID
	asn := m.Neighbor.ASN
	uri := models.BgpNeighborPath(asn, id)

	r := models.IOSXERequest{}
	r.HTTPMethod = "GET"
	r.Key = &m.Neighbor.ID
	r.Path = uri

	body, err := request.Read(&c.Config, &r)
	if err != nil {
		return nil, err
	}

	res := models.BgpNeighbor{}
	if body == nil {
		return &res, nil
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *CiscoIOSXEClient) UpdateBgpNeighbor(m models.BgpNeighbor) error {
	id := m.Neighbor.ID
	asn := m.Neighbor.ASN
	uri := models.BgpNeighborPath(asn, id)

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] %v", string(rb))

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	r.Key = &m.Neighbor.ID
	r.Path = uri
	r.Payload = rb

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	return nil
}

func (c *CiscoIOSXEClient) DeleteBgpNeighbor(m models.BgpNeighbor) error {
	id := m.Neighbor.ID
	asn := m.Neighbor.ASN
	uri := models.BgpNeighborPath(asn, id)

	r := models.IOSXERequest{}
	r.HTTPMethod = "DELETE"
	r.Key = &m.Neighbor.ID
	r.Path = uri

	err := request.Delete(&c.Config, &r)
	if err != nil {
		return err
	}

	return nil
}

/*
func (c *CiscoIOSXEClient) ListBgpNeighbor() (*models.BgpNeighborList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/neighbor=", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.BgpNeighborList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
*/
