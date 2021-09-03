package client

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/poroping/go-ios-xe-sdk/models"
)

func (c *Client) CreateBgpNeighborConfig(asn int, m models.BgpNeighborConfig) error {
	id := m.NeighborConfig.ID
	uri := GetBgpNeighborConfigURI(asn, id, nil)

	neighbor := models.BgpNeighbor{}
	neighbor.Neighbor.ID = id

	parent_exists, _ := c.ReadBgpNeighbor(asn, neighbor)
	if parent_exists == nil {
		return fmt.Errorf("neighbor %q does not exist", id)
	}

	exists, _ := c.ReadBgpNeighborConfig(asn, m)
	if exists != nil {
		return c.UpdateBgpNeighborConfig(asn, m)
	}

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s", c.HostURL, uri), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 201)

	if err != nil {
		return err
	}

	return nil
}

func (c *Client) ReadBgpNeighborConfig(asn int, m models.BgpNeighborConfig) (*models.BgpNeighborConfig, error) {
	id := m.NeighborConfig.ID
	uri := GetBgpNeighborConfigURI(asn, id, nil)

	neighbor := models.BgpNeighbor{}
	neighbor.Neighbor.ID = id

	parent_exists, _ := c.ReadBgpNeighbor(asn, neighbor)
	if parent_exists == nil {
		return nil, fmt.Errorf("neighbor %q does not exist", id)
	}

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, uri), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.BgpNeighborConfig{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateBgpNeighborConfig(asn int, m models.BgpNeighborConfig) error {
	id := m.NeighborConfig.ID
	uri := GetBgpNeighborConfigURI(asn, id, nil)

	neighbor := models.BgpNeighbor{}
	neighbor.Neighbor.ID = id

	parent_exists, _ := c.ReadBgpNeighbor(asn, neighbor)
	if parent_exists == nil {
		return fmt.Errorf("neighbor %q does not exist", id)
	}

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

func (c *Client) DeleteBgpNeighborConfig(asn int, m models.BgpNeighborConfig) error {
	id := m.NeighborConfig.ID
	uri := GetBgpNeighborConfigURI(asn, id, nil)

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

func (c *Client) ListBgpNeighborConfig(as int) (*models.BgpNeighborConfigList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/%d/address-family/no-vrf/ipv4/unicast/neighbor/", c.HostURL, as), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.BgpNeighborConfigList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
