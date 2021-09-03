package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/poroping/go-ios-xe-sdk/models"
)

func (c *Client) CreateBgpRouter(uri string, m models.BgpRouter) error {
	id := m.Bgp.ID

	exists, _ := c.ReadBgpRouter(uri, m)
	if exists != nil {
		return c.UpdateBgpRouter(uri, m)
	}

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s=%s", c.HostURL, uri, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ReadBgpRouter(uri string, m models.BgpRouter) (*models.BgpRouter, error) {
	id := m.Bgp.ID

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s=%s", c.HostURL, uri, id), nil)
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

func (c *Client) UpdateBgpRouter(uri string, m models.BgpRouter) error {
	id := m.Bgp.ID

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/%s=%s", c.HostURL, uri, id), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteBgpRouter(uri string, m models.BgpRouter) error {
	id := m.Bgp.ID

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s=%s", c.HostURL, uri, id), nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListBgpRouter(uri string) (*models.BgpRouterList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, uri), nil)
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