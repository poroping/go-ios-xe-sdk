package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"log"

	"github.com/poroping/go-ios-xe-sdk/models"
)

func (c *Client) CreateVlan(m models.Vlan) error {
	id := m.VlanList.ID

	exists, _ := c.ReadVlan(m)

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/restconf/data/Cisco-IOS-XE-native:native/vlan/vlan-list=%d", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}

	if exists != nil {
		return c.UpdateVlan(m)
	}

	_, err = c.doRequest(req, 201)

	if err != nil {
		return err
	}

	log.Printf("Vlan %d Created", id)
	return nil
}

func (c *Client) ReadVlan(m models.Vlan) (*models.Vlan, error) {
	id := m.VlanList.ID

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/data/Cisco-IOS-XE-native:native/vlan/vlan-list=%d", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.Vlan{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateVlan(m models.Vlan) error {
	id := m.VlanList.ID

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/restconf/data/Cisco-IOS-XE-native:native/vlan/vlan-list=%d", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	log.Printf("VLAN %d already exists, UPDATING", id)

	return nil
}

func (c *Client) DeleteVlan(id int) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/restconf/api/running/native/vlan/vlan-list/%d", c.HostURL, id), nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListVlan() (*models.VlanList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/vlan/vlan-list/", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.VlanList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
