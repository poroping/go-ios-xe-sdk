package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/poroping/go-ios-xe-sdk/models"
)

const sviURI = "restconf/data/Cisco-IOS-XE-native:native/interface/Vlan"

func (c *Client) CreateSVI(m models.SVI) error {
	id := m.Vlan.Name

	// exists, _ := c.ReadSVI(m)
	// if exists != nil {
	// 	return c.UpdateSVI(m)
	// }

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s=%d", c.HostURL, sviURI, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	_, err = c.doRequest(req, 0)

	if err != nil {
		return err
	}

	log.Printf("SVI %d created", id)

	return nil
}

func (c *Client) ReadSVI(m models.SVI) (*models.SVI, error) {
	id := m.Vlan.Name

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s=%d", c.HostURL, sviURI, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.SVI{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateSVI(m models.SVI) error {
	id := m.Vlan.Name

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s=%d", c.HostURL, sviURI, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	log.Printf("SVI %d updated", id)

	return nil
}

func (c *Client) DeleteSVI(m models.SVI) error {
	id := m.Vlan.Name

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s=%d", c.HostURL, sviURI, id), nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	log.Printf("SVI %d deleted", id)

	return nil
}

// func (c *Client) ListSVI() (*models.SVIList, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/interface/Vlan/", c.HostURL), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, 200)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := models.SVIList{}
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &res, nil
// }
