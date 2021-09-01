package client

import (
	"crypto/tls"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	password   string
	username   string
	userAgent  string
}

func NewClient(host, username, password, userAgent string, insecure bool) (*Client, error) {
	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second, Transport: tr},
		HostURL:    host,
		username:   username,
		password:   password,
		userAgent:  userAgent,
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, sc int) ([]byte, error) {
	req.Header.Set("Content-Type", "application/vnd.yang.data+json")
	req.Header.Set("Accept", "application/vnd.yang.data+json, application/vnd.yang.collection+json")
	req.Header.Set("User-Agent", c.userAgent)
	req.SetBasicAuth(c.username, c.password)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != sc {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
