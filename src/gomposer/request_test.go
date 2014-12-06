package gomposer

import (
	"fmt"

	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func getHttpClient(server *httptest.Server) *HttpClient {

	hc, _ := NewHttpClient(server.URL)

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
