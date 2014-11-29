package gomposer

import (
	"github.com/mcuadros/go-version"
)

type DependencyResolver struct {
	requiredPackages map[string][]*version.ConstraintGroup
	versions         map[string][]string
}

func (dr DependencyResolver) AddRequirement(packageName string, versionRule string) {
	cg := version.NewConstrainGroupFromString(versionRule)
	dr.requiredPackages[packageName] = append(dr.requiredPackages[packageName], cg)
}
func (dr DependencyResolver) AddPackages(packageName string, versions []string) {
	dr.versions[packageName] = versions
}

func (dr DependencyResolver) Resolve() map[string]string {
	m := make(map[string]string)
	for k, _ := range dr.versions {
		m[k] = dr.versions[k][0]
	}
	return m
}
