package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/poroping/go-ios-xe-sdk/models"
)

func (c *Client) CreateL2Vlan(id int, name interface{}) error {

	m := models.L2Vlan{VlanList: models.VlanList{
		ID: id,
	},
	}

	if name != nil {
		m.VlanList.Name = fmt.Sprintf("%v", name)
	}

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/restconf/api/running/native/vlan/vlan-list/%d", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ReadL2Vlan(id int) (*models.L2Vlan, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/vlan/vlan-list/%d", c.HostURL, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.L2Vlan{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateL2Vlan(id int, name string) error {
	m := models.L2Vlan{VlanList: models.VlanList{
		Name: name,
	},
	}

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/restconf/api/running/native/vlan/vlan-list/%d", c.HostURL, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteL2Vlan(id int) error {
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

func (c *Client) ListL2Vlan() (*models.L2VlanList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/vlan/vlan-list/", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.L2VlanList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
