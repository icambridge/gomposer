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
// TODO rename
func (dr DependencyResolver) AddPackages(packageName string, versions []string) {
	dr.versions[packageName] = versions
}

func (dr DependencyResolver) Resolve() map[string]string {
	output := make(map[string]string)
	for packageName, contraintList := range dr.requiredPackages {
		versions := dr.versions[packageName]
		notValid := []string{}
		for _, version := range versions {
			for _, contraint := range contraintList {
				if contraint.Match(version) != true {
					notValid = append(notValid, version)
				}
			}
			// TODO clean
			m := make(map[string]int)
			for _, version := range notValid {
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

			versions = validVersions
			notValid = []string{}
		}
		//
		if len(versions) > 0 {

			output[packageName] = versions[0]
		}

	}
	return output
}
