package gomposer

import (
	"fmt"
	"github.com/mcuadros/go-version"
	"sort"
	// "strings"
)

type DependencyResolver struct {
	requiredPackages map[string]map[string]*version.ConstraintGroup
	versions         map[string][]string
	replacedPackages map[string]int
}

func (dr DependencyResolver) AddReplacement(packageName string) {
	dr.replacedPackages[packageName] = 1
}

func (dr DependencyResolver) AddRequirement(packageName string, versionRule string) {

	cg := version.NewConstrainGroupFromString(versionRule)
	_, ok := dr.requiredPackages[packageName]
	if !ok {
		dr.requiredPackages[packageName] = map[string]*version.ConstraintGroup{}
	}
	dr.requiredPackages[packageName][versionRule] = cg
}

// TODO rename
func (dr DependencyResolver) AddPackages(packageName string, versions []string) {

	sort.Sort(sort.Reverse(VersionSlice(versions)))
	dr.versions[packageName] = versions
}

func (dr DependencyResolver) Resolve() map[string]string {
	output := make(map[string]string)
	for packageName, contraintList := range dr.requiredPackages {
		versions := dr.versions[packageName]
		_, ok := dr.replacedPackages[packageName]

		if ok {
			continue
		}

		// TODO use key to create new slice based on where we bailed
		for _, version := range versions {

			failed := 0
			if len(contraintList) == 0 {
				fmt.Println("Fails")
			}

			for _, contraint := range contraintList {
				if contraint.Match(version) != true {
					failed++
					// if failed check to see if
					break
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
