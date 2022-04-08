package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/poroping/go-ios-xe-sdk/models"
	"github.com/poroping/go-ios-xe-sdk/request"
)

func (c *CiscoIOSXEClient) CreateL2Vlan(m models.L2VlanList) error {
	id := m.VlanList.ID

	// exists, _ := c.ReadL2Vlan(m)
	// if exists != nil {
	// 	return c.UpdateL2Vlan(m)
	// }

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	// r.Key = &m.ID
	r.Payload = rb
	r.Path = fmt.Sprintf("%s=%v", models.L2VlanPath, id)

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] L2Vlan %d Created", id)

	return nil
}

func (c *CiscoIOSXEClient) ReadL2Vlan(m models.L2VlanList) (*models.L2VlanList, error) {
	id := m.VlanList.ID

	r := models.IOSXERequest{}
	r.HTTPMethod = "GET"
	// r.Key = &m.ID
	r.Path = fmt.Sprintf("%s=%v", models.L2VlanPath, id)

	body, err := request.Read(&c.Config, &r)
	if err != nil {
		return nil, err
	}

	res := models.L2VlanList{}
	if body == nil {
		return &res, nil
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *CiscoIOSXEClient) UpdateL2Vlan(m models.L2VlanList) error {
	id := m.VlanList.ID

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] %v", string(rb))

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	// r.Key = &m.ID
	r.Path = fmt.Sprintf("%s=%v", models.L2VlanPath, id)
	r.Payload = rb

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] L2Vlan %d updated", id)

	return nil
}

func (c *CiscoIOSXEClient) DeleteL2Vlan(m models.L2VlanList) error {
	id := m.VlanList.ID

	r := models.IOSXERequest{}
	r.HTTPMethod = "DELETE"
	// r.Key = &m.ID
	r.Path = fmt.Sprintf("%s=%v", models.L2VlanPath, id)

	err := request.Delete(&c.Config, &r)
	if err != nil {
		return err
	}

	return nil
}

func (c *CiscoIOSXEClient) ListL2Vlan() (*[]models.L2VlanList, error) {
	r := models.IOSXERequest{}
	r.HTTPMethod = "GET"
	// r.Key = &m.Vlan.Name
	r.Path = models.L2VlanPath

	body, err := request.Read(&c.Config, &r)
	if err != nil {
		return nil, err
	}

	res := []models.L2VlanList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
