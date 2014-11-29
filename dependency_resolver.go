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
	for packageName, contraintList := range dr.requiredPackages {
		versions := dr.versions[packageName]
		for _, contraint := range contraintList {
			for i, version := range versions {
				if contraint.Match(version) == false {
					versions[i], versions = versions[len(versions)-1], versions[:len(versions)-1]
				}
			}
		}

		m[packageName] = versions[0]


	}
	return m
}
