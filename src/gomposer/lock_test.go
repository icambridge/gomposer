package gomposer

import (
	"fmt"
	"net/http"
	"testing"
)

func TestLockGeneratesLockFile(t *testing.T) {

	mux, server := getMuxAndServer()
	mux.HandleFunc("/m/e.json", func(w http.ResponseWriter, r *http.Request) {
			fmt.Fprint(w, `{"package":{"name":"m\/e", "versions": {"dev-master": {"name":"m\/e", "version": "dev-master"}, "1.0.0": {"name":"m\/e", "version": "1.0.0"},"1.0.1": {"name":"m\/e", "version": "1.0.1"},"1.1.0": {"name":"m\/e", "version": "1.1.0"},"1.1.1": {"name":"m\/e", "version": "1.1.1"}}}}`)

		})

	hc := getHttpClient(server)
	packageRepo := PackageRepository{Client: hc}

	lockGenerator := LockGenerator{
		PackageRepo: packageRepo,
	}

	dependencies := map[string]string{
		`m/e`: "dev-master",
	}
	lock := lockGenerator.Generate(dependencies)

	if len(lock.Packages) == 0 {
		t.Errorf("Expected packages in lock file")
		return
	}
	expected := "dev-master"

	if lock.Packages[0].Name != `m/e` || lock.Packages[0].Version != expected {
		t.Errorf("Expected version = %s got = %s", expected, lock.Packages[0].Version)
	}
}
