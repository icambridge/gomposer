package gomposer

import (
	"fmt"

	"net/http"
	"net/http/httptest"
	// "reflect"
	"testing"
)

func getMuxAndServer() (*http.ServeMux, *httptest.Server) {

	mux := http.NewServeMux()
	server := httptest.NewServer(mux)

	return mux, server
}

func TestPackageRepository_Find(t *testing.T) {

	mux, server := getMuxAndServer()
	apiHit := false
	mux.HandleFunc("/phpunit/phpunit.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"name":"phpunit/phpunit"}`)
		apiHit = true
	})

	hc := getHttpClient(server)
	packageRepo := PackageRepository{client: hc}
	pkg, err := packageRepo.Find("phpunit/phpunit")

	if err != nil {
		t.Errorf("Didn't expect an error but got '%s'", err)
	}

	if apiHit == false {
		t.Errorf("Didn't hit api")
	}
	expectedName := "phpunit/phpunit"
	if pkg.Name != expectedName {
		t.Errorf("Expected '%s' but got '%s'", expectedName, pkg.Name)
	}

	server.Close()
}
