package main

import (
    "bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func NewHttpClient(baseUrl string) (*HttpClient, error) {

	url, err := url.Parse(baseUrl)

	if err != nil {
		return nil, err
	}

	httpClient := http.DefaultClient
	hc := HttpClient{BaseURL: url, client: httpClient}

	return &hc, nil
}

type Client interface {
	Request(method string, url string) error
}

type HttpClient struct {
	BaseURL *url.URL

	client *http.Client
}

func (hc HttpClient) Request(method string, uri string, output interface{}) error {

	if method != "GET" && method != "POST" && method != "DELETE" && method != "HEAD" && method != "PUT" {
		return errors.New(fmt.Sprintf("Invalid http method '%s'", method))
	}

	u, err := url.Parse(strings.TrimRight(hc.BaseURL.String(), "/") + uri)

	if err != nil {
		return err
	}
	buf := new(bytes.Buffer)
	req, err := http.NewRequest(method, u.String(), buf)

	if err != nil {
		return err
	}

	resp, err := hc.client.Do(req)


	if err != nil {
		return err
	}


	err = json.NewDecoder(resp.Body).Decode(output)

	return err
}
