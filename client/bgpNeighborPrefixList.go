package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/poroping/go-ios-xe-sdk/models"
)

func (c *Client) CreateBgpNeighborPrefixList(id string, as int, inout string, prefixlist string) error {
	m := models.BgpNeighborPrefixList{}
	m.PrefixList.Inout = inout
	m.PrefixList.PrefixListName = prefixlist
	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/address-family/no-vrf/ipv4/unicast/neighbor/%s/prefix-list/%s/", c.HostURL, as, id, inout), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ReadBgpNeighborPrefixList(id string, as int, inout string) (*models.BgpNeighborPrefixList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/address-family/no-vrf/ipv4/unicast/neighbor/%s/prefix-list/%s/", c.HostURL, as, id, inout), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.BgpNeighborPrefixList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateBgpNeighborPrefixList(id string, as int, inout string, prefixlist string) error {
	m := models.BgpNeighborPrefixList{}
	m.PrefixList.Inout = inout
	m.PrefixList.PrefixListName = prefixlist
	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PATCH", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/address-family/no-vrf/ipv4/unicast/neighbor/%s/prefix-list/%s/", c.HostURL, as, id, inout), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) DeleteBgpNeighborPrefixList(id string, as int, inout string) error {
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/address-family/no-vrf/ipv4/unicast/neighbor/%s/prefix-list/%s/", c.HostURL, as, id, inout), nil)
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 204)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ListBgpNeighborPrefixList(id string, as int) (*models.BgpNeighborPrefixListList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/address-family/no-vrf/ipv4/unicast/neighbor/%s/prefix-list/", c.HostURL, as, id), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.BgpNeighborPrefixListList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
