package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/poroping/go-ios-xe-sdk/models"
)

const bgpRouterURI = "restconf/data/Cisco-IOS-XE-native:native/router/Cisco-IOS-XE-bgp:bgp"

func (c *Client) CreateBgpRouter(m models.BgpRouter) error {
	id := m.Bgp.ID

	exists, _ := c.ReadBgpRouter(m)
	if exists != nil {
		return c.UpdateBgpRouter(m)
	}

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s=%s", c.HostURL, bgpRouterURI, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ReadBgpRouter(m models.BgpRouter) (*models.BgpRouter, error) {
	id := m.Bgp.ID

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s=%s", c.HostURL, bgpRouterURI, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.BgpRouter{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateBgpRouter(m models.BgpRouter) error {
	id := m.Bgp.ID

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/%s=%s", c.HostURL, bgpRouterURI, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteBgpRouter(m models.BgpRouter) error {
	id := m.Bgp.ID

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s=%s", c.HostURL, bgpRouterURI, id), nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListBgpRouter() (*models.BgpRouterList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, bgpRouterURI), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.BgpRouterList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}