package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"

	"github.com/poroping/go-ios-xe-sdk/config"
	"github.com/poroping/go-ios-xe-sdk/models"
)

const contentType string = "application/yang-data+json"

func CreateUpdate(c *config.Config, r *models.IOSXERequest) error {
	req, err := newRequest(*c, r)
	if err != nil {
		return err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// 201 on create 204 on update

	if res.StatusCode == http.StatusNoContent || res.StatusCode == http.StatusCreated {
		return nil
	}

	err = ciscoErrorCheck(body, res)
	if err != nil {
		return err
	}

	err = fmt.Errorf("uncaught error.\nStatus Code: %v\nBody: %v", res.StatusCode, string(body))

	return err
}

func Read(c *config.Config, r *models.IOSXERequest) ([]byte, error) {
	req, err := newRequest(*c, r)
	if err != nil {
		return nil, err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		return nil, err
	}

	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// 200 on success
	if res.StatusCode == http.StatusOK {
		return body, nil
	}

	if res.StatusCode == http.StatusNotFound {
		return nil, nil
	}

	err = ciscoErrorCheck(body, res)
	if err != nil {
		return nil, err
	}

	err = fmt.Errorf("uncaught error.\nStatus Code: %v\nBody: %v", res.StatusCode, string(body))

	return nil, err
}

func Delete(c *config.Config, r *models.IOSXERequest) (err error) {
	req, err := newRequest(*c, r)
	if err != nil {
		return err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}

	// 204 on success
	// 404 if not present

	if res.StatusCode == http.StatusNoContent {
		// success
		return nil
	}

	if res.StatusCode == http.StatusNotFound {
		// not present, check if restconf, if not return error
		return restconfCheck(*c)
	}

	err = ciscoErrorCheck(body, res)
	if err != nil {
		return err
	}

	err = fmt.Errorf("[ERROR] Uncaught error.\nStatus Code: %v\nBody: %v", res.StatusCode, string(body))

	return
}

// func Generic(c *config.Config, r *models.IOSXERequest) (*string, error) {
// 	req, err := newRequest(*c, r)
// 	if err != nil {
// 		return nil, err
// 	}

// 	res, err := c.HTTPCon.Do(req)
// 	if err != nil {
// 		// handle 404s (search for resource on delete at least)
// 		return nil, err
// 	}

// 	log.Printf("[DEBUG] Status code: %d", res.StatusCode)

// 	defer res.Body.Close()

// 	body, err := ioutil.ReadAll(res.Body)
// 	if err != nil {
// 		return nil, err
// 	}

// 	response := interface{}{}
// 	err = json.Unmarshal(body, &response)
// 	if err != nil {
// 		log.Printf("[ERROR] Error reading response body during CREATE/UPDATE %s", err)
// 		return nil, err
// 	}

// 	err = ciscoErrorCheck(body, response)
// 	if err != nil {
// 		return nil, err
// 	}

// 	resp := string(body)

// 	return &resp, err
// }

// Currently only checks for errors, don't trust false, needs work.
// func Exists(c *config.Config, r *models.IOSXERequest) bool {
// 	_, err := Read(c, r)
// 	return err == nil
// }

// Currently only checks for errors, don't trust false, needs work.
// func ExistsWithResponse(c *config.Config, r *models.IOSXERequest) (*interface{}, bool) {
// 	resp, err := Read(c, r)
// 	if err != nil {
// 		return resp, false
// 	}
// 	return resp, true
// }

func ciscoErrorCheck(body []byte, res *http.Response) error {
	if body == nil || string(body) == "" {
		return noBodyErrorCheck(res)
	}
	response := &models.CiscoErrorResp{}
	err := json.Unmarshal(body, response)
	if err != nil {
		return fmt.Errorf("[ERROR] Unable to parse response into error message. %v \n %w", string(body), err)
	}

	parsedError := parseCiscoError(response)
	if parsedError != nil {
		return parsedError
	}

	return nil
}

func parseCiscoError(e *models.CiscoErrorResp) error {
	if e != nil {
		err0 := e.CiscoErrors.CiscoError[0]
		if err0.ErrorMessage != nil {
			errMsg := err0.ErrorMessage
			return fmt.Errorf("%v", *errMsg)
		}
	}
	return nil
}

func noBodyErrorCheck(res *http.Response) (err error) {
	switch hs := res.StatusCode; hs {
	case http.StatusBadRequest:
		err = fmt.Errorf("[ERROR] Bad Request (%d)", res.StatusCode)
	case http.StatusUnauthorized:
		err = fmt.Errorf("[ERROR] Not Authorized (%d)", res.StatusCode)
	case http.StatusForbidden:
		err = fmt.Errorf("[ERROR] Forbidden (%d)", res.StatusCode)
	case http.StatusNotFound:
		err = fmt.Errorf("[ERROR] Resource Not Found (%d)", res.StatusCode)
	case http.StatusMethodNotAllowed:
		err = fmt.Errorf("[ERROR] Method Not Allowed (%d)", res.StatusCode)
	case http.StatusConflict:
		err = fmt.Errorf("[ERROR] Confict (%d)", res.StatusCode)
	case http.StatusInternalServerError:
		err = fmt.Errorf("[ERROR] Internal Server Error (%d)", res.StatusCode)
	default:
		err = fmt.Errorf("[ERROR] Unknown Error without body (%d)", res.StatusCode)
	}

	return
}

// Returns an error if HEAD to "/restconf/data/Cisco-IOS-XE-native:native" does not respond with StatusOK
func restconfCheck(c config.Config) error {
	r := models.IOSXERequest{}
	r.HTTPMethod = "HEAD"
	r.Path = models.BasePath
	req, err := newRequest(c, &r)
	if err != nil {
		return err
	}

	res, err := c.HTTPCon.Do(req)
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return fmt.Errorf("[ERROR] Device does not appear to have restconf running.\nCheck the host and/or if restconf is enabled")
	}

	log.Print("[INFO] Host appears to be Cisco IOS XE running restconf")

	return nil
}

func newRequest(c config.Config, r *models.IOSXERequest) (*http.Request, error) {
	headers := buildHeaders(c, r)
	body := bytes.NewBuffer(r.Payload)
	if body != nil {
		if body.String() != "" {
			log.Printf("[DEBUG] BODY: %s", body.String())
		}
	}
	// set URL w/ url queries
	url := buildURL(c, r)

	req, err := http.NewRequest(r.HTTPMethod, url.String(), body)
	if err != nil {
		return nil, err
	}
	req.Header = *headers
	req.SetBasicAuth(c.Username, c.Password)

	req.URL = url

	return req, nil
}

func buildHeaders(c config.Config, r *models.IOSXERequest) *http.Header {
	headers := http.Header{}
	headers.Set("Content-Type", contentType)
	headers.Set("Accept", "application/yang-data+json,application/vnd.yang.collection+json,application/yang-patch+json")
	headers.Set("User-Agent", c.UserAgent)

	return &headers
}

func buildURL(c config.Config, r *models.IOSXERequest) *url.URL {
	u := url.URL{}
	u.Scheme = "https"
	u.Host = c.Host
	q := url.Values{}
	// q := marshalParams(&r.Params)
	u.Path = r.Path

	u.RawQuery = q.Encode()

	return &u
}
