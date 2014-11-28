package gomposer

import (
	"fmt"

	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func getHttpClient(server *httptest.Server) HttpClient {

	url, _ := url.Parse(server.URL)
	httpClient := http.DefaultClient
	hc := HttpClient{BaseURL: url, client: httpClient}

	return hc
}

func TestHttpClient_Do(t *testing.T) {

	mux, server := getMuxAndServer()

	type Foo struct {
		Bar string
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := "POST"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"Bar":"drink"}`)
	})

	hc := getHttpClient(server)
	body := &Foo{}
	hc.Request("POST", "/", body)

	expected := &Foo{"drink"}

	if !reflect.DeepEqual(body, expected) {
		t.Errorf("Response body = %v, expected %v", body, expected)
	}
}
