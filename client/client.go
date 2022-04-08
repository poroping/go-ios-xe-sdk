package client

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strings"
	"time"

	"github.com/hashicorp/go-retryablehttp"
	"github.com/poroping/go-ios-xe-sdk/config"
	"github.com/poroping/go-ios-xe-sdk/models"
)

type CiscoIOSXEClient struct {
	Config config.Config
}

func NewClient(cfg config.Config) (*CiscoIOSXEClient, error) {
	insecure := cfg.Insecure

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
			r := models.CiscoErrorResp{}
			body, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				return false, err
			}
			err = json.Unmarshal(body, &r)
			if err != nil {
				return false, err
			}
			msg := r.CiscoErrors.CiscoError[0].ErrorMessage
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

	cfg.HTTPCon = retryClient.StandardClient()

	c := CiscoIOSXEClient{
		Config: cfg,
	}

	return &c, nil
}
