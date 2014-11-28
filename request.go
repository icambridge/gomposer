package gomposer

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

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

	resp, err := hc.client.Do(req)

	if err != nil {
		return err
	}

	json.NewDecoder(resp.Body).Decode(output)

	return nil
}
