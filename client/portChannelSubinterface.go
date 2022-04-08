package client

import (
	"encoding/json"
	"fmt"
	"log"

	"github.com/poroping/go-ios-xe-sdk/models"
	"github.com/poroping/go-ios-xe-sdk/request"
)

func (c *CiscoIOSXEClient) CreatePortChannel(m models.PortChannel) error {
	id := m.PortChannel.Name

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	r.Key = &m.PortChannel.Name
	r.Payload = rb
	r.Path = fmt.Sprintf("%s=%v", models.PortChannelPath, id)

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] PortChannel %v created", id)

	return nil
}

func (c *CiscoIOSXEClient) ReadPortChannel(m models.PortChannel) (*models.PortChannel, error) {
	id := m.PortChannel.Name

	r := models.IOSXERequest{}
	r.HTTPMethod = "GET"
	r.Key = &m.PortChannel.Name
	r.Path = fmt.Sprintf("%s=%v", models.PortChannelPath, id)

	body, err := request.Read(&c.Config, &r)
	if err != nil {
		return nil, err
	}

	res := models.PortChannel{}
	if body == nil {
		return &res, nil
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	log.Printf("[INFO] PortChannel %v read", id)

	return &res, nil
}

func (c *CiscoIOSXEClient) UpdatePortChannel(m models.PortChannel) error {
	id := m.PortChannel.Name

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] %v", string(rb))

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	r.Key = &m.PortChannel.Name
	r.Path = fmt.Sprintf("%s=%v", models.PortChannelPath, id)
	r.Payload = rb

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] PortChannel %v updated", id)

	return nil
}

func (c *CiscoIOSXEClient) DeletePortChannel(m models.PortChannel) error {
	id := m.PortChannel.Name

	r := models.IOSXERequest{}
	r.HTTPMethod = "DELETE"
	r.Key = &m.PortChannel.Name
	r.Path = fmt.Sprintf("%s=%v", models.PortChannelPath, id)

	err := request.Delete(&c.Config, &r)
	if err != nil {
		return err
	}

	log.Printf("[INFO] PortChannel %v deleted", id)

	return nil
}

// func (c *CiscoIOSXEClient) ListPortChannel() (*models.PortChannelList, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/interface/PortChannel/", c.HostURL), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, 200)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := models.PortChannelList{}
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &res, nil
// }
