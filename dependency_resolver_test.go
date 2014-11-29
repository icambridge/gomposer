package gomposer

import (
	"github.com/mcuadros/go-version"
	"reflect"
	"testing"
)

func Test_DependencyResolver_Resolve(t *testing.T) {
	m := make(map[string][]*version.ConstraintGroup)
	v := make(map[string][]string)
	dr := DependencyResolver{requiredPackages: m, versions: v}

	expected := map[string]string{"symfony/symfony": "2.3.0"}
	dr.AddPackages("symfony/symfony", []string{"2.3.0"})
	dr.AddRequirement("symfony/symfony", "2.3.0")

	actual := dr.Resolve()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected %v", actual, expected)
	}
}

func Test_DependencyResolver_Resolve_Multiple(t *testing.T) {
	m := make(map[string][]*version.ConstraintGroup)
	v := make(map[string][]string)
	dr := DependencyResolver{requiredPackages: m, versions: v}

	expected := map[string]string{"symfony/symfony": "2.3.1"}
	dr.AddPackages("symfony/symfony", []string{"2.3.0","2.3.1","2.3.6"})
	dr.AddRequirement("symfony/symfony", ">2.3.0")
	dr.AddRequirement("symfony/symfony", ">2.3.0,<2.3.5")

	actual := dr.Resolve()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected %v", actual, expected)
	}
}
