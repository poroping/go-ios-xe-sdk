package client

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/poroping/go-ios-xe-sdk/models"
)

func (c *Client) CreateBgpNeighbor(asn int, m models.BgpNeighbor) error {
	id := m.Neighbor.ID
	uri := GetBgpNeighborURI(asn, id)

	// disabled cause can't PATCH certain fields "off"
	// exists, _ := c.ReadBgpNeighbor(asn, m)
	// if exists != nil {
	// 	return c.UpdateBgpNeighbor(asn, m)
	// }

	// log.Printf("Doesn't exist, will create")

	rb, err := json.Marshal(m)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/%s", c.HostURL, uri), strings.NewReader(string(rb)))
	if err != nil {
		return err
	}
	_, err = c.doRequest(req, 0) // set 0 cause PUT create and PUT update are diff 200 codes

	if err != nil {
		return err
	}

	log.Println("[INFO] neighbor created.")

	return nil
}

func (c *Client) ReadBgpNeighbor(asn int, m models.BgpNeighbor) (*models.BgpNeighbor, error) {
	id := m.Neighbor.ID
	uri := GetBgpNeighborURI(asn, id)

	req, err := http.NewRequest("GET", fmt.Sprintf("%s/%s", c.HostURL, uri), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.BgpNeighbor{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) UpdateBgpNeighbor(asn int, m models.BgpNeighbor) error {
	id := m.Neighbor.ID
	uri := GetBgpNeighborURI(asn, id)

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

func (c *Client) DeleteBgpNeighbor(asn int, m models.BgpNeighbor) error {
	id := m.Neighbor.ID
	uri := GetBgpNeighborURI(asn, id)

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

/*
func (c *Client) ListBgpNeighbor() (*models.BgpNeighborList, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/restconf/api/running/native/router/bgp/43892/neighbor=", c.HostURL), nil)
	if err != nil {
		return nil, err
	}

	body, err := c.doRequest(req, 200)
	if err != nil {
		return nil, err
	}

	res := models.BgpNeighborList{}
	err = json.Unmarshal(body, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
*/
