package gomposer

import (
	// "reflect"
	"fmt"
	"net/http"
	// "net/http/httptest"
	"testing"
)

func Test_ProcessHit_All_Directly_Required(t *testing.T) {

	mux, server := getMuxAndServer()
	apiHitSymfony := false
	apiHitDoctrine := false
	mux.HandleFunc("/symfony/symfony.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
        fmt.Fprint(w, `{"name":"symfony/symfony", "versions": {"2.3.5":{ "name": "symfony/symfony", "require": {"twigphp/twig": "2.4.1"}}}}`)
		apiHitSymfony = true
	})

	mux.HandleFunc("/doctrine/doctrine.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
        fmt.Fprint(w, `{"name":"doctrine/doctrine", "versions": {"2.3.5":{ "name": "doctrine/doctrine", "require": {"twigphp/twig": "2.4.1"}}}}`)
		apiHitDoctrine = true
	})

	v := &Version{
		Require: map[string]string{
			"symfony/symfony":   "2.3.5",
			"doctrine/doctrine": "2.3.5",
		},
	}

    m := make(map[string]*PackageInfo)
    hc := getHttpClient(server)
    pr := PackageRepository{client: hc}

    p := &Process{packageRepo: &pr, packages: m}
	p.Process(v)

	if apiHitSymfony == false || apiHitDoctrine == false {
		t.Error("Expected Api to be hit")
	}
}

func Test_ProcessHit_All_Required_Including_Vendors(t *testing.T) {

	mux, server := getMuxAndServer()
	apiHitSymfony := false
    apiHitDoctrine := false
    apiHitTwig := false
	mux.HandleFunc("/symfony/symfony.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"name":"symfony/symfony", "versions": {"2.3.5":{ "name": "symfony/symfony", "require": {"twigphp/twig": "2.4.1"}}}}`)
		apiHitSymfony = true
	})

	mux.HandleFunc("/twigphp/twig.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"name":"twigphp/twig", "versions": {"2.4.1":{ "name": "twigphp/twig"}}}`)
        apiHitTwig = true
	})
	mux.HandleFunc("/doctrine/doctrine.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"name":"doctrine/doctrine", "versions": {"2.3.5":{ "name": "doctrine/doctrine", "require": {"twigphp/twig": "2.4.1"}}}}`)
		apiHitDoctrine = true
	})

	v := &Version{
		Require: map[string]string{
			"symfony/symfony":   "2.3.5",
			"doctrine/doctrine": "2.3.5",
		},
	}
    m := make(map[string]*PackageInfo)
	hc := getHttpClient(server)
	pr := PackageRepository{client: hc}

	p := &Process{packageRepo: &pr, packages: m}
	p.Process(v)

	if apiHitSymfony == false || apiHitDoctrine == false || apiHitTwig == false {
		t.Error("Expected Api to be hit")
	}
}
