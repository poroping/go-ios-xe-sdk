package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/poroping/go-ios-xe-sdk/models"
)

const contentType = "application/yang-data+json"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	password   string
	username   string
	userAgent  string
}

func NewClient(host, username, password, userAgent string, insecure bool) (*Client, error) {
	retryClient := retryablehttp.NewClient()
	retryClient.RetryMax = 10

	tr := &http.Transport{
		TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
	}

	retryClient.HTTPClient.Transport = tr
	retryClient.HTTPClient.Timeout = 10 * time.Second

	// add custom retry check for config sync in progress (http 409)
	retryClient.CheckRetry = func(ctx context.Context, resp *http.Response, err error) (bool, error) {
		if resp == nil {
			return true, nil
		}
		if resp.StatusCode == 409 {
			r := models.CiscoError{}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return false, err
			}
			err = json.Unmarshal(body, &r)
			if err != nil {
				return false, err
			}
			msg := r.Errors.Error[0].ErrorMessage
			if strings.Contains(*msg, "sync in progress") {
				return true, nil
			}
		}

		return retryablehttp.DefaultRetryPolicy(ctx, resp, err)
	}

	// add custom backoff for config sync in progress (http 409)
	retryClient.Backoff = func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
		if resp != nil {
			if resp.StatusCode == http.StatusConflict {
				return time.Second * time.Duration(attemptNum)
				// return math.Pow(1.5, float64(attemptNum)) * float64(min)  // alternate?
			}
		}

		return retryablehttp.DefaultBackoff(min, max, attemptNum, resp)
	}

	c := Client{
		HTTPClient: retryClient.StandardClient(),
		HostURL:    host,
		username:   username,
		password:   password,
		userAgent:  userAgent,
	}

	return &c, nil
}

func (c *Client) doRequest(req *http.Request, sc int) ([]byte, error) {
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("Accept", contentType)
	req.Header.Set("User-Agent", c.userAgent)
	req.SetBasicAuth(c.username, c.password)

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	log.Printf("Status code: %d", res.StatusCode)
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// ignore 404s - TODO: temp solution as state refresh can break if resource no longer exists
	// prevents checking if res exists and changing to update
	if res.StatusCode == 404 {
		log.Printf("Status code: %d for request: %s\n    Body: %s \n\n", res.StatusCode, req.RequestURI, body)
		emptybody := []byte(`{}`)
		return emptybody, nil
	}

	if res.StatusCode != sc && sc != 0 {
		log.Printf("Status code: %d for request: %s\n    Body: %s \n\n", res.StatusCode, req.RequestURI, body)

		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
