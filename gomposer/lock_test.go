package gomposer

import (
	"fmt"
	"net/http"
	"reflect"
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

func TestLockGeneratesLockFile_Sorted(t *testing.T) {

	mux, server := getMuxAndServer()
	mux.HandleFunc("/m/e.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"package":{"name":"m\/e", "versions": {"dev-master": {"name":"m\/e", "version": "dev-master"}, "1.0.0": {"name":"m\/e", "version": "1.0.0"},"1.0.1": {"name":"m\/e", "version": "1.0.1"},"1.1.0": {"name":"m\/e", "version": "1.1.0"},"1.1.1": {"name":"m\/e", "version": "1.1.1"}}}}`)

	})

	mux.HandleFunc("/z/e.json", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `{"package":{"name":"z\/e", "versions": {"dev-master": {"name":"z\/e", "version": "dev-master"}, "1.0.0": {"name":"m\/e", "version": "1.0.0"},"1.0.1": {"name":"m\/e", "version": "1.0.1"},"1.1.0": {"name":"m\/e", "version": "1.1.0"},"1.1.1": {"name":"m\/e", "version": "1.1.1"}}}}`)

	})
	hc := getHttpClient(server)
	packageRepo := PackageRepository{Client: hc}

	lockGenerator := LockGenerator{
		PackageRepo: packageRepo,
	}

	dependencies := map[string]string{
		`m/e`: "dev-master",
		`z/e`: "dev-master",
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
	expectedSecondName := `z/e`
	if lock.Packages[1].Name != expectedSecondName || lock.Packages[1].Version != expected {
		t.Errorf("Expected name = %s got = %s", expectedSecondName, lock.Packages[1].Name)
	}
}

func Test_DiffLock_Returns_Added_Packages(t *testing.T) {
	expected := Version{Name: "behat/behat", Version: "1.1.2"}
	new := Lock{
		Packages: []Version{
			expected,
		},
	}

	old := Lock{
		Packages: []Version{
			Version{Name: "behat/behat", Version: "1.1.1"},
		},
	}

	output := DiffLock(new, old)

	if actualCount, expectedCount := len(output["added"]), 1; actualCount != expectedCount {
		t.Errorf("Expected only %v item got %v", expectedCount, actualCount)
		return
	}

	if !reflect.DeepEqual(output["added"][0], expected) {
		t.Errorf("Expected %v, got %v", expected, output["added"][0])
		return
	}
}

func Test_DiffLock_Returns_Removed_Packages(t *testing.T) {
	expected := Version{Name: "behat/behat", Version: "1.1.2"}
	new := Lock{
		Packages: []Version{},
	}

	old := Lock{
		Packages: []Version{
			expected,
		},
	}

	output := DiffLock(new, old)

	if actualCount, expectedCount := len(output["removed"]), 1; actualCount != expectedCount {
		t.Errorf("Expected only %v item got %v", expectedCount, actualCount)
		return
	}

	if !reflect.DeepEqual(output["removed"][0], expected) {
		t.Errorf("Expected %v, got %v", expected, output["removed"][0])
		return
	}
}
func Test_DiffLock_Returns_Removed_Packages_When_Replaced(t *testing.T) {
	expected := Version{Name: "behat/behat", Version: "1.1.2"}
	new := Lock{
		Packages: []Version{
			Version{Name: "behat/behat", Version: "1.1.1"},
		},
	}

	old := Lock{
		Packages: []Version{
			expected,
		},
	}

	output := DiffLock(new, old)

	if actualCount, expectedCount := len(output["removed"]), 1; actualCount != expectedCount {
		t.Errorf("Expected only %v item got %v", expectedCount, actualCount)
		return
	}

	if !reflect.DeepEqual(output["removed"][0], expected) {
		t.Errorf("Expected %v, got %v", expected, output["removed"][0])
		return
	}
}
