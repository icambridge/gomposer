package gomposer

import (
	// "reflect"
    "fmt"
    "net/http"
    // "net/http/httptest"
	"testing"
)

func Test_ProcessHit_All_Required(t *testing.T) {

	mux, server := getMuxAndServer()
	apiHitSymfony := false
    apiHitDoctrine := false
	mux.HandleFunc("/symfony/symfony.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"name":"symfony/symfony"}`)
		apiHitSymfony = true
	})


    mux.HandleFunc("/doctrine/doctrine.json", func(w http.ResponseWriter, r *http.Request) {
        if m := "GET"; m != r.Method {
            t.Errorf("Request method = %v, expected %v", r.Method, m)
        }
        fmt.Fprint(w, `{"name":"symfony/symfony"}`)
        apiHitDoctrine = true
        })

	v := &Version{
		Require: map[string]string{
            "symfony/symfony": "2.3.5",
            "doctrine/doctrine": "2.3.5",
		},
	}


    hc := getHttpClient(server)
    pr := PackageRepository{client: hc}

	p := &Process{packageRepo: &pr}
	p.Process(v)

    if apiHitSymfony == false || apiHitDoctrine == false {
        t.Error("Expected Api to be hit")
    }
}
