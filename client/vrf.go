package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/poroping/go-ios-xe-sdk/models"
	"github.com/poroping/go-ios-xe-sdk/request"
)

func (c *CiscoIOSXEClient) CreateVRF(m models.VRFDefinition) error {
	id := m.VRFDefinition.Name

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	r.Key = &m.VRFDefinition.Name
	r.Payload = rb
	r.Path = fmt.Sprintf("%s=%v", models.VRFPath, id)

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] VRF %v created", id)

	return nil
}

func (c *CiscoIOSXEClient) ReadVRF(m models.VRFDefinition) (*models.VRFDefinition, error) {
	id := m.VRFDefinition.Name

	r := models.IOSXERequest{}
	r.HTTPMethod = "GET"
	r.Key = &m.VRFDefinition.Name
	r.Path = fmt.Sprintf("%s=%v", models.VRFPath, id)

	body, err := request.Read(&c.Config, &r)
	if err != nil {
		return nil, err
	}

	res := models.VRFDefinition{}
	if body == nil {
		return &res, nil
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] VRF %v read", id)

	return &res, nil
}

func (c *CiscoIOSXEClient) UpdateVRF(m models.VRFDefinition) error {
	id := m.VRFDefinition.Name

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] %v", string(rb))

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	r.Key = &m.VRFDefinition.Name
	r.Path = fmt.Sprintf("%s=%v", models.VRFPath, id)
	r.Payload = rb

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] VRF %v updated", id)

	return nil
}

func (c *CiscoIOSXEClient) DeleteVRF(m models.VRFDefinition) error {
	id := m.VRFDefinition.Name

	r := models.IOSXERequest{}
	r.HTTPMethod = "DELETE"
	r.Key = &m.VRFDefinition.Name
	r.Path = fmt.Sprintf("%s=%v", models.VRFPath, id)

	err := request.Delete(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] VRF %v deleted", id)

	return nil
}

// func (c *CiscoIOSXEClient) ListVRF() (*models.VRFDefinition, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/interface/VRF/", c.HostURL), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, 200)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := models.VRFDefinition{}
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &res, nil
// }
