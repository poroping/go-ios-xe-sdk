package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/poroping/go-ios-xe-sdk/models"
	"github.com/poroping/go-ios-xe-sdk/request"
)

func (c *CiscoIOSXEClient) CreatePortChannelSubinterface(m models.PortChannelSubinterface) error {
	id := m.PortChannelSubinterface.Name

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	r.Key = &m.PortChannelSubinterface.Name
	r.Payload = rb
	r.Path = fmt.Sprintf("%s=%v", models.PortChannelSubinterfacePath, id)

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] PortChannelSubinterface %v created", id)

	return nil
}

func (c *CiscoIOSXEClient) ReadPortChannelSubinterface(m models.PortChannelSubinterface) (*models.PortChannelSubinterface, error) {
	id := m.PortChannelSubinterface.Name

	r := models.IOSXERequest{}
	r.HTTPMethod = "GET"
	r.Key = &m.PortChannelSubinterface.Name
	r.Path = fmt.Sprintf("%s=%v", models.PortChannelSubinterfacePath, id)

	body, err := request.Read(&c.Config, &r)
	if err != nil {
		return nil, err
	}

	res := models.PortChannelSubinterface{}
	if body == nil {
		return &res, nil
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] PortChannelSubinterface %v read", id)

	return &res, nil
}

func (c *CiscoIOSXEClient) UpdatePortChannelSubinterface(m models.PortChannelSubinterface) error {
	id := m.PortChannelSubinterface.Name

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] %v", string(rb))

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	r.Key = &m.PortChannelSubinterface.Name
	r.Path = fmt.Sprintf("%s=%v", models.PortChannelSubinterfacePath, id)
	r.Payload = rb

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] PortChannelSubinterface %v updated", id)

	return nil
}

func (c *CiscoIOSXEClient) DeletePortChannelSubinterface(m models.PortChannelSubinterface) error {
	id := m.PortChannelSubinterface.Name

	r := models.IOSXERequest{}
	r.HTTPMethod = "DELETE"
	r.Key = &m.PortChannelSubinterface.Name
	r.Path = fmt.Sprintf("%s=%v", models.PortChannelSubinterfacePath, id)

	err := request.Delete(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] PortChannelSubinterface %v deleted", id)

	return nil
}

// func (c *CiscoIOSXEClient) ListPortChannelSubinterface() (*models.PortChannelSubinterfaceList, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/interface/PortChannelSubinterface/", c.HostURL), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, 200)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := models.PortChannelSubinterfaceList{}
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &res, nil
// }
