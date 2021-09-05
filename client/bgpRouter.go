package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/poroping/go-ios-xe-sdk/models"
)

func (c *Client) CreateBgpRouter(m models.BgpRouter) error {
	id := m.Bgp.ID
	uri := GetBgpURI(id)

	// exists, _ := c.ReadBgpRouter(m)
	// if exists != nil {
	// 	return c.UpdateBgpRouter(m)
	// }

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s", c.HostURL, uri), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 0)

	if err != nil {
		return err
	}

	// create ipv4 unicast AF
	// TODO: probs move this to some other func later

	payload := `{
		"Cisco-IOS-XE-bgp:ipv4": {
			"af-name": "unicast"
		}
	}`
	uri += "/address-family/no-vrf/ipv4=unicast"

	req, err = http.NewRequest("PUT", fmt.Sprintf("%s/%s", c.HostURL, uri), strings.NewReader(payload))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 0)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ReadBgpRouter(m models.BgpRouter) (*models.BgpRouter, error) {
	id := m.Bgp.ID
	uri := GetBgpURI(id)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, uri), nil)
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
	uri := GetBgpURI(id)

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/%s", c.HostURL, uri), strings.NewReader(string(rb)))
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
	uri := GetBgpURI(id)

	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/%s", c.HostURL, uri), nil)
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
