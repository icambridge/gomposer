package gomposer

import (
	"sort"
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

	sort.Sort(sort.Reverse(sort.StringSlice(versions)))
	dr.versions[packageName] = versions
}

func (dr DependencyResolver) Resolve() map[string]string {
	output := make(map[string]string)
	for packageName, contraintList := range dr.requiredPackages {
		versions := dr.versions[packageName]
		for _, version := range versions {
			failed := 0
			for _, contraint := range contraintList {
				if contraint.Match(version) != true {
					failed++
					// step over. We can't use since it's failed.
					// Remove from list so we don't check againist it again.
				}
			}
			if failed == 0 {
				output[packageName] = version
				break
			}
		}

	}
	return output
}
