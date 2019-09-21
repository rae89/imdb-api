package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const (
	path           = "/Users/gravitywaves/Projects/imdb-api/imdb-api-read.txt"
	defaultBaseURL = "http://www.omdbapi.com"
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
	baseURL string
	apiKey  string
	client  *http.Client
}

func NewClient(apiKey string) *Client {
	return &Client{
		baseURL: defaultBaseURL,
		apiKey:  getAPIKey(path),
		client:  http.DefaultClient,
	}
}

func main() {

	fmt.Print(string("START"))
}
