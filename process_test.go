package main

import (
	"fmt"
	"net/http"
	"reflect"
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
		fmt.Fprint(w, `{"package":{"name":"symfony/symfony", "versions": {"2.3.5":{ "name": "symfony/symfony", "version": "2.3.5", "require": {"twigphp/twig": "2.4.1"}}}}}}`)
		apiHitSymfony = true
	})

	mux.HandleFunc("/twigphp/twig.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"package":{"name":"twigphp/twig", "versions": {"2.4.1":{ "name": "twigphp/twig", "version": "2.4.1"}}}}`)

		})
	mux.HandleFunc("/doctrine/doctrine.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"package":{"name":"doctrine/doctrine", "versions": {"2.3.5":{ "name": "doctrine/doctrine", "version": "2.3.5", "require": {"twigphp/twig": "2.4.1"}}}}}}`)
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
	pr := PackageRepository{Client: hc}

	p := &Process{PackageRepo: &pr, Packages: m}
	p.Process(v)

	if apiHitSymfony == false || apiHitDoctrine == false {
		t.Error("Expected Api to be hit")
	}
}

func Test_ProcessHit_All_Required_Including_Vendors_Once_Each(t *testing.T) {

	mux, server := getMuxAndServer()
	apiHitTwig := 0
	mux.HandleFunc("/symfony/symfony.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"package":{"name":"symfony/symfony", "versions": {"2.3.5":{ "name": "symfony/symfony", "version": "2.3.5", "require": {"twigphp/twig": "2.4.1"}}}}}`)
	})

	mux.HandleFunc("/twigphp/twig.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"package":{"name":"twigphp/twig", "versions": {"2.4.1":{ "name": "twigphp/twig", "version": "2.4.1"}}}}`)
		apiHitTwig++
	})
	mux.HandleFunc("/doctrine/doctrine.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"package":{"name":"doctrine/doctrine", "versions": {"2.3.5":{ "name": "doctrine/doctrine", "version": "2.3.5", "require": {"twigphp/twig": "2.4.1"}}}}}`)
	})

	v := &Version{
		Require: map[string]string{
			"symfony/symfony":   "2.3.5",
			"doctrine/doctrine": "2.3.5",
		},
	}
	m := make(map[string]*PackageInfo)
	hc := getHttpClient(server)
	pr := PackageRepository{Client: hc}

	p := &Process{PackageRepo: &pr, Packages: m}
	p.Process(v)

	if apiHitTwig != 1 {
		t.Errorf("Expected Api to be hit to be hit once for twig but %v times", apiHitTwig)
	}
}

func Test_Process_Returns_CorrectDependencies(t *testing.T) {
	mux, server := getMuxAndServer()
	// TODO Clean up this.
	mux.HandleFunc("/symfony/symfony.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"package":{"name":"symfony/symfony", "versions": {"2.3.5":{ "name": "symfony/symfony", "version": "2.3.5", "require": {"twigphp/twig": "2.4.1"}}}}}`)
	})

	mux.HandleFunc("/twigphp/twig.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"package":{"name":"twigphp/twig", "versions": {"2.4.1":{ "name": "twigphp/twig", "version": "2.4.1"}}}}`)
	})

	mux.HandleFunc("/doctrine/doctrine.json", func(w http.ResponseWriter, r *http.Request) {
		if m := "GET"; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		fmt.Fprint(w, `{"package":{"name":"doctrine/doctrine", "versions": {"2.3.5":{ "name": "doctrine/doctrine", "version": "2.3.5", "require": {"twigphp/twig": "2.4.1"}}}}}`)
	})

	v := &Version{
		Require: map[string]string{
			"symfony/symfony":   "2.3.5",
			"doctrine/doctrine": "2.3.5",
		},
	}
	m := make(map[string]*PackageInfo)
	hc := getHttpClient(server)
	pr := PackageRepository{Client: hc}

	p := &Process{PackageRepo: &pr, Packages: m}
	actual := p.Process(v)

	expected := &Lock{
		Packages: []Version{
			Version{
				Name:    "symfony/symfony",
				Version: "2.3.5",
				Require: map[string]string{"twigphp/twig": "2.4.1"},
			},
			Version{
				Name:    "doctrine/doctrine",
				Version: "2.3.5",
				Require: map[string]string{"twigphp/twig": "2.4.1"},
			},
			Version{
				Name:    "twigphp/twig",
				Version: "2.4.1",
			},
		},
	}
	// TODO  look into correctness of this. Failed with
	/* --- FAIL: Test_Process_Returns_CorrectDependencies (0.01 seconds)
	process_test.go:156: Response actual = &{[{doctrine/doctrine  2.3.5 [] {   } {   } map[twigphp/twig:2.4.1] map[] map[] map[]} {twigphp/twig  2.4.1 [] {   } {   } map[] map[] map[] map[]} {symfony/symfony  2.3.5 [] {   } {   } map[twigphp/twig:2.4.1] map[] map[] map[]}] []}, expected &{[{symfony/symfony  2.3.5 [] {   } {   } map[twigphp/twig:2.4.1] map[] map[] map[]} {doctrine/doctrine  2.3.5 [] {   } {   } map[twigphp/twig:2.4.1] map[] map[] map[]} {twigphp/twig  2.4.1 [] {   } {   } map[] map[] map[] map[]}] []}
	FAIL */
	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response actual = %v, expected %v", actual, expected)
	}

}
