package gomposer

import (
	"reflect"
	"testing"
)

func Test_PackageReader_Read(t *testing.T) {

	expected := Version{
		Name:    "symfony/framework-standard-edition",
		Require: map[string]string{"php": ">=5.3.3", "symfony/symfony": "2.3.*"},
	}

	r := PackageReader{}
	actual, _ := r.Read("fixtures/composer.json")

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("Response actual = %v, expected %v", actual, expected)
	}
}
