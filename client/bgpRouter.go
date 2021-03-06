package client

import (
	"encoding/json"
	"log"

	"github.com/poroping/go-ios-xe-sdk/models"
	"github.com/poroping/go-ios-xe-sdk/request"
)

func (c *CiscoIOSXEClient) CreateBgpRouter(m models.BgpRouter) error {
	id := m.Bgp.ID
	uri := models.BgpPath(id)

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	// r.Key = &m.Bgp.ID
	r.Payload = rb
	r.Path = uri

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	return nil
}

func (c *CiscoIOSXEClient) ReadBgpRouter(m models.BgpRouter) (*models.BgpRouter, error) {
	id := m.Bgp.ID
	uri := models.BgpPath(id)

	r := models.IOSXERequest{}
	r.HTTPMethod = "GET"
	// r.Key = &m.Bgp.ID
	r.Path = uri

	body, err := request.Read(&c.Config, &r)
	if err != nil {
		return nil, err
	}

	res := models.BgpRouter{}
	if body == nil {
		return &res, nil
	}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *CiscoIOSXEClient) UpdateBgpRouter(m models.BgpRouter) error {
	id := m.Bgp.ID
	uri := models.BgpPath(id)

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	log.Printf("[DEBUG] %v", string(rb))

	r := models.IOSXERequest{}
	r.HTTPMethod = "PUT"
	// r.Key = &m.Bgp.ID
	r.Path = uri
	r.Payload = rb

	err = request.CreateUpdate(&c.Config, &r)
	if err != nil {
		return err
	}

	return nil
}

func (c *CiscoIOSXEClient) DeleteBgpRouter(m models.BgpRouter) error {
	id := m.Bgp.ID
	uri := models.BgpPath(id)

	r := models.IOSXERequest{}
	r.HTTPMethod = "DELETE"
	// r.Key = &m.Bgp.ID
	r.Path = uri

	err := request.Delete(&c.Config, &r)
	if err != nil {
		return err
	}

	return nil
}

// func (c *CiscoIOSXEClient) ListBgpRouter(uri string) (*models.BgpRouterList, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, uri), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, 200)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := models.BgpRouterList{}
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &res, nil
// }
