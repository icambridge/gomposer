package gomposer

import (
	"github.com/mcuadros/go-version"
	"reflect"
	"testing"
)

func Test_DependencyResolver_Resolve(t *testing.T) {

  m := make(map[string]map[string]*version.ConstraintGroup)
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

  m := make(map[string]map[string]*version.ConstraintGroup)
	v := make(map[string][]string)
	dr := DependencyResolver{requiredPackages: m, versions: v}

	expected := map[string]string{"symfony/symfony": "2.3.1"}
	dr.AddPackages("symfony/symfony", []string{"2.3.0", "2.3.1", "2.3.6"})
	dr.AddRequirement("symfony/symfony", ">2.3.0")
	dr.AddRequirement("symfony/symfony", ">2.3.0,<2.3.5")

	actual := dr.Resolve()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected %v", actual, expected)
	}
}

func Test_DependencyResolver_Resolve_Complex(t *testing.T) {

  m := make(map[string]map[string]*version.ConstraintGroup)
	v := make(map[string][]string)
	dr := DependencyResolver{requiredPackages: m, versions: v}

	expected := map[string]string{
		"symfony/symfony":   "2.3.2",
		"doctrine/orm":      "2.3.2",
		"phpunit/phpunit":   "2.3.2",
		"mailgun/php-sdk":   "2.3.2",
		"racksapce/php-sdk": "2.3.2",
		"twigphp/twig":      "2.3.2",
		"behat/behat":       "2.3.2",
	}
	dr.AddPackages("symfony/symfony", []string{"2.3.0", "2.3.1", "2.3.6", "2.3.2", "2.4.1", "2.4.2", "2.4.2", "2.4.3", "2.4.4"})
	dr.AddPackages("doctrine/orm", []string{"2.3.0", "2.3.1", "2.3.6", "2.3.2", "2.4.1", "2.4.2", "2.4.2", "2.4.3", "2.4.4"})
	dr.AddPackages("phpunit/phpunit", []string{"2.3.0", "2.3.1", "2.3.6", "2.3.2", "2.4.1", "2.4.2", "2.4.2", "2.4.3", "2.4.4"})
	dr.AddPackages("mailgun/php-sdk", []string{"2.3.0", "2.3.1", "2.3.6", "2.3.2", "2.4.1", "2.4.2", "2.4.2", "2.4.3", "2.4.4"})
	dr.AddPackages("racksapce/php-sdk", []string{"2.3.0", "2.3.1", "2.3.6", "2.3.2", "2.4.1", "2.4.2", "2.4.2", "2.4.3", "2.4.4"})
	dr.AddPackages("twigphp/twig", []string{"2.3.0", "2.3.1", "2.3.6", "2.3.2", "2.4.1", "2.4.2", "2.4.2", "2.4.3", "2.4.4"})
	dr.AddPackages("behat/behat", []string{"2.3.0", "2.3.1", "2.3.6", "2.3.2", "2.4.1", "2.4.2", "2.4.2", "2.4.3", "2.4.4"})

	dr.AddRequirement("symfony/symfony", ">2.3.0")
	dr.AddRequirement("symfony/symfony", ">2.3.0,<2.3.5")
	dr.AddRequirement("symfony/symfony", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("symfony/symfony", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("symfony/symfony", "2.3.*")
	dr.AddRequirement("symfony/symfony", "~2.3.0")
	dr.AddRequirement("symfony/symfony", ">2.3.1,<2.3.5")
	dr.AddRequirement("symfony/symfony", ">=2.3.0,<2.3.5")
	dr.AddRequirement("symfony/symfony", ">2.3.0,<=2.3.5")
	dr.AddRequirement("symfony/symfony", ">2.3.0,<=2.3.6")

	dr.AddRequirement("doctrine/orm", ">2.3.0")
	dr.AddRequirement("doctrine/orm", ">2.3.0,<2.3.5")
	dr.AddRequirement("doctrine/orm", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("doctrine/orm", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("doctrine/orm", "2.3.*")
	dr.AddRequirement("doctrine/orm", "~2.3.0")
	dr.AddRequirement("doctrine/orm", ">2.3.0,<2.3.5")
	dr.AddRequirement("doctrine/orm", ">=2.3.0,<2.3.5")
	dr.AddRequirement("doctrine/orm", ">2.3.0,<=2.3.5")
	dr.AddRequirement("doctrine/orm", ">2.3.0,<=2.3.6")

	dr.AddRequirement("phpunit/phpunit", ">2.3.0")
	dr.AddRequirement("phpunit/phpunit", ">2.3.0,<2.3.5")
	dr.AddRequirement("phpunit/phpunit", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("phpunit/phpunit", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("phpunit/phpunit", "2.3.*")
	dr.AddRequirement("phpunit/phpunit", "~2.3.0")
	dr.AddRequirement("phpunit/phpunit", ">2.3.1,<2.3.5")
	dr.AddRequirement("phpunit/phpunit", ">=2.3.0,<2.3.5")
	dr.AddRequirement("phpunit/phpunit", ">2.3.0,<=2.3.5")
	dr.AddRequirement("phpunit/phpunit", ">2.3.0,<=2.3.6")

	dr.AddRequirement("mailgun/php-sdk", ">2.3.0")
	dr.AddRequirement("mailgun/php-sdk", ">2.3.0,<2.3.5")
	dr.AddRequirement("mailgun/php-sdk", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("mailgun/php-sdk", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("mailgun/php-sdk", "2.3.*")
	dr.AddRequirement("mailgun/php-sdk", "~2.3.0")
	dr.AddRequirement("mailgun/php-sdk", ">2.3.1,<2.3.5")
	dr.AddRequirement("mailgun/php-sdk", ">=2.3.0,<2.3.5")
	dr.AddRequirement("mailgun/php-sdk", ">2.3.0,<=2.3.5")
	dr.AddRequirement("mailgun/php-sdk", ">2.3.0,<=2.3.6")

	dr.AddRequirement("racksapce/php-sdk", ">2.3.0")
	dr.AddRequirement("racksapce/php-sdk", ">2.3.0,<2.3.5")
	dr.AddRequirement("racksapce/php-sdk", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("racksapce/php-sdk", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("racksapce/php-sdk", "2.3.*")
	dr.AddRequirement("racksapce/php-sdk", "~2.3.0")
	dr.AddRequirement("racksapce/php-sdk", ">2.3.1,<2.3.5")
	dr.AddRequirement("racksapce/php-sdk", ">=2.3.0,<2.3.5")
	dr.AddRequirement("racksapce/php-sdk", ">2.3.0,<=2.3.5")
	dr.AddRequirement("racksapce/php-sdk", ">2.3.0,<=2.3.6")

	dr.AddRequirement("twigphp/twig", ">2.3.0")
	dr.AddRequirement("twigphp/twig", ">2.3.0,<2.3.5")
	dr.AddRequirement("twigphp/twig", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("twigphp/twig", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("twigphp/twig", "2.3.*")
	dr.AddRequirement("twigphp/twig", "~2.3.0")
	dr.AddRequirement("twigphp/twig", ">2.3.1,<2.3.5")
	dr.AddRequirement("twigphp/twig", ">=2.3.0,<2.3.5")
	dr.AddRequirement("twigphp/twig", ">2.3.0,<=2.3.5")
	dr.AddRequirement("twigphp/twig", ">2.3.0,<=2.3.6")

	dr.AddRequirement("behat/behat", ">2.3.0")
	dr.AddRequirement("behat/behat", ">2.3.0,<2.3.5")
	dr.AddRequirement("behat/behat", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("behat/behat", ">=2.3.0,<=2.3.5")
	dr.AddRequirement("behat/behat", "2.3.*")
	dr.AddRequirement("behat/behat", "~2.3.0")
	dr.AddRequirement("behat/behat", ">2.3.1,<2.3.5")
	dr.AddRequirement("behat/behat", ">=2.3.0,<2.3.5")
	dr.AddRequirement("behat/behat", ">2.3.0,<=2.3.5")
	dr.AddRequirement("behat/behat", ">2.3.0,<=2.3.6")

	actual := dr.Resolve()

	if !reflect.DeepEqual(actual, expected) {
		t.Errorf("actual = %v, expected %v", actual, expected)
	}
}
