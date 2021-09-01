package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/poroping/go-ios-xe-sdk/models"
)

func (c *Client) CreateL3Vlan(m models.L3Vlan) error {
	id := m.Vlan.Name

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/restconf/data/Cisco-IOS-XE-native:native/interface/Vlan=%d", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ReadL3Vlan(m models.L3Vlan) (*models.L3Vlan, error) {
	id := m.Vlan.Name

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/data/Cisco-IOS-XE-native:native/interface/Vlan=%d", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.L3Vlan{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateL3Vlan(m models.L3Vlan) error {
	id := m.Vlan.Name

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/restconf/data/Cisco-IOS-XE-native:native/interface/Vlan=%d", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteL3Vlan(m models.L3Vlan) error {
	id := m.Vlan.Name

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/restconf/data/Cisco-IOS-XE-native:native/interface/Vlan=%d", c.HostURL, id), nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	return nil
}

// func (c *Client) ListL3Vlan() (*models.L3VlanList, error) {
// 	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/interface/Vlan/", c.HostURL), nil)
// 	if err != nil {
// 		return nil, err
// 	}

// 	body, err := c.doRequest(req, 200)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res := models.L3VlanList{}
// 	err = json.Unmarshal(body, &res)
// 	if err != nil {
// 		return nil, err
// 	}

// 	return &res, nil
// }
