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
	output := make(map[string]string)
	for packageName, contraintList := range dr.requiredPackages {
		versions := dr.versions[packageName]
		notValid := []string{}
		for _, contraint := range contraintList {
			for _, version := range versions {
				if contraint.Match(version) == false {
					notValid = append(notValid, version)
				}
			}
		}
		m := make(map[string]int)

		for _,version := range notValid {
			m[version]++
		}
		var validVersions []string
		for _, version := range versions {
			if m[version] > 0 {
				m[version]--
				continue
			}
			validVersions = append(validVersions, version)
		}
		output[packageName] = validVersions[0]


	}
	return output
}
