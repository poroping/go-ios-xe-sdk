package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/poroping/go-ios-xe-sdk/models"
	"github.com/poroping/go-ios-xe-sdk/request"
)

func (c *CiscoIOSXEClient) CreateVlan(m models.Vlan) error {
	id := m.Vlan.Name

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	r.Key = &m.Vlan.Name
	r.Payload = rb
	r.Path = fmt.Sprintf("%s=%v", models.VlanPath, id)

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Vlan %v created", id)

	return nil
}

func (c *CiscoIOSXEClient) ReadVlan(m models.Vlan) (*models.Vlan, error) {
	id := m.Vlan.Name

	r := models.IOSXERequest{}
	r.HTTPMethod = "GET"
	r.Key = &m.Vlan.Name
	r.Path = fmt.Sprintf("%s=%v", models.VlanPath, id)

	body, err := request.Read(&c.Config, &r)
	if err != nil {
		return nil, err
	}

	res := models.Vlan{}
	if body == nil {
		return &res, nil
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] Vlan %v read", id)

	return &res, nil
}

func (c *CiscoIOSXEClient) UpdateVlan(m models.Vlan) error {
	id := m.Vlan.Name

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] %v", string(rb))

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	r.Key = &m.Vlan.Name
	r.Path = fmt.Sprintf("%s=%v", models.VlanPath, id)
	r.Payload = rb

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Vlan %v updated", id)

	return nil
}

func (c *CiscoIOSXEClient) DeleteVlan(m models.Vlan) error {
	id := m.Vlan.Name

	r := models.IOSXERequest{}
	r.HTTPMethod = "DELETE"
	r.Key = &m.Vlan.Name
	r.Path = fmt.Sprintf("%s=%v", models.VlanPath, id)

	err := request.Delete(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] Vlan %v deleted", id)

	return nil
}

// func (c *CiscoIOSXEClient) ListVlan() (*models.VlanList, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/interface/Vlan/", c.HostURL), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, 200)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := models.VlanList{}
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &res, nil
// }
