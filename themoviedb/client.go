package themoviedb

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/pkg/errors"
)

const (
	defaultBaseURL = "https://api.themoviedb.org"
	apiVersion     = "3"
	mediaType      = "application/json"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getAPIKey(filePath string) string {
	dat, err := ioutil.ReadFile(filePath)
	check(err)
	return string(dat)
}

type Client struct {
	baseURL    string
	apiVersion string
	apiKey     string
	client     *http.Client
}

func NewClient(filePath string) *Client {
	return &Client{
		baseURL:    defaultBaseURL,
		apiVersion: apiVersion,
		apiKey:     getAPIKey(filePath),
		client:     http.DefaultClient,
	}
}

func (c *Client) doRequest(method, endpoint string, data interface{}, queryparams map[string]string) ([]byte, error) {
	// Encode data if we passed an object
	b := bytes.NewBuffer(nil)
	if data != nil {
		// Create the encoder
		enc := json.NewEncoder(b)
		if err := enc.Encode(data); err != nil {
			return nil, errors.Wrap(err, "json encoding data for doRequest failed")
		}
	}

	// Create the request
	uri := fmt.Sprintf("%s/%s/%s", c.baseURL, c.apiVersion, strings.Trim(endpoint, "/"))
	req, err := http.NewRequest(method, uri, b)

	if queryparams != nil {
		q := req.URL.Query()
		for k, v := range queryparams {
			q.Add(k, v)
		}
		req.URL.RawQuery = q.Encode()
	}

	fmt.Println("URI: ", uri)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("creating %s request to %s failed", method, uri))
	}

	// Set the proper headers
	// req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", mediaType)

	// Do the request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("performing  %s request to %s failed", method, uri))
	}
	defer resp.Body.Close()

	fmt.Println(resp.StatusCode)
	// Check that the response status code was OK
	switch resp.StatusCode {
	case http.StatusOK:
	case http.StatusCreated:
	case http.StatusUnauthorized:
	case http.StatusForbidden:
	case http.StatusNotFound:
	case http.StatusBadRequest:
		return nil, fmt.Errorf("the request in invailid")
	default:
		return nil, fmt.Errorf("bad response code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("decoding response from %s request to %s failed: body -> %s\n", method, uri, string(body)))
	}

	return body, nil

}
