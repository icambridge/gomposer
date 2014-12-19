package gomposer

import (
	"reflect"
	"testing"

	"os"
)

func Test_PackageReader_Read(t *testing.T) {

	expected := Version{
		Name:    "symfony/framework-standard-edition",
		Require: map[string]string{"php": ">=5.3.3", "symfony/symfony": "2.3.*"},
	}

	f, err := os.Create("composer.json")

	if err != nil {
		t.Error(err)
	}
	b := []byte(`{
    "name": "symfony/framework-standard-edition",
    "require": {
        "php": ">=5.3.3",
        "symfony/symfony": "2.3.*"
    }
}`)
	_,err = f.Write(b)
	if err != nil {
		t.Error(err)
	}
	r := PackageReader{}
	actual, err := r.Read("composer.json")

	if err != nil {
		t.Error(err)
		return
	}

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response actual = %v, expected %v", actual, expected)
	}
}
