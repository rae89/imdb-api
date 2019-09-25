package main

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
	path           = "/Users/gravitywaves/Projects/imdb-api/imdb-api-read.txt"
	defaultBaseURL = "https://api.themoviedb.org"
	apiVersion     = "3"
	mediaType      = "application/json"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getAPIKey(path string) string {
	dat, err := ioutil.ReadFile(path)
	check(err)
	return string(dat)
}

type Client struct {
	baseURL    string
	apiVersion string
	apiKey     string
	client     *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		baseURL:    defaultBaseURL,
		apiVersion: apiVersion,
		apiKey:     getAPIKey(path),
		client:     http.DefaultClient,
	}
}

func (c *Client) doRequest(method, endpoint string, data interface{}) ([]byte, error) {
	// Encode data if we passed an object
	b := bytes.NewBuffer(nil)
	if data != nil {
		// Create the encoder
		enc := json.NewEncoder(b)
		if err := enc.Encode(data); err != nil {
			return nil, errors.Wrap(err, fmt.Sprintf("creating %s request to %s failed", method, uri))
		}
	}

	// Create the request
	uri := fmt.Sprintf("%s/%s/%s", c.baseURL, c.apiVersion, strings.Trim(endpoint, "/"))
	req, err := http.NewRequest(method, uri, b)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("creating %s request to %s failed", method, uri))
	}

	// Set the proper headers
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", c.apiKey))
	req.Header.Set("Content-Type", mediaType)

	// Do the request
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("performing  %s request to %s failed", method, uri))
	}
	defer resp.Body.Close()

	// Check that the response status code was OK
	switch resp.StatusCode {
		case http.StatusOK:
		case http.StatusCreated:
		case http.StatusUnauthorized:
			return nil, fmt.Errorf("invalid access token")
		case http.StatusForbidden:
			return nil, fmt.Errorf("unauthorized access to endpoint")
		case http.StatusNotFound:
			return nil, fmt.Errorf("unauthorized access to endpoint")
		case http.StatusBadRequest:
			return nil, fmt.Errorf("the request in invailid")
		default:
			return nil, fmt.Errorf("bad response code: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, fmt.Sprintf("decoding response from %s request to %s failed: body -> %s\n", methodm uri, string(body)))
	}

	return body, nil

}

func main() {

	fmt.Print(string("START"))
}
