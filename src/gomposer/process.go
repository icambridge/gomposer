package gomposer

import (
	"github.com/mcuadros/go-version" // TODO remove the need for this to be imported all over the show
	"sort"
	"strings"
)

type Process struct {
	PackageRepo *PackageRepository
	Packages    map[string]PackageInfo
	dr          *DependencyResolver
}

// TODO recrusive
func (p Process) Process(packageInfo *Version) *Lock {

	// TODO inject
	m := make(map[string]map[string]*version.ConstraintGroup)
	v := make(map[string][]string)
	rp := make(map[string]int)

	if p.dr == nil {
		p.dr = &DependencyResolver{requiredPackages: m, versions: v, replacedPackages: rp}
	}

	p.inner(packageInfo.Require)
	// TODO solve this.
	// Just process all dependencies anyways and then solve.
	requiredVersions := p.dr.Resolve()
	return p.generateLock(requiredVersions)
}

func (p Process) generateLock(requiredVersions map[string]string) *Lock {
	l := &Lock{Packages: []Version{}}

	names := []string{}
	for packageName, _ := range requiredVersions {
		names = append(names, packageName)
	}

	sort.Strings(names)

	for _, name := range names {
		versionNum := requiredVersions[name]
		packageInfo := p.Packages[name].Versions[versionNum]
		l.Packages = append(l.Packages, packageInfo)
	}

	return l
}

func (p Process) inner(require map[string]string) {
	foundPackages := p.getRequire(require)
	if len(foundPackages) == 0 {
		return
	}
	p.addPackages(foundPackages)
	requiredVersions := p.dr.Resolve()

	for packageName, versionNum := range requiredVersions {
		packageInfo := p.Packages[packageName].Versions[versionNum]
		p.inner(packageInfo.Require)
	}
}

func (p Process) addPackages(packages map[string]PackageInfo) {

	for packageName, packageInfo := range packages {
		p.Packages[packageName] = packageInfo
		versions := make([]string, 0, len(packageInfo.Versions))
		for versionNum, version := range packageInfo.Versions {

			for replacedPackage, _ := range version.Replace {
				p.dr.AddReplacement(replacedPackage)
			}
			versions = append(versions, versionNum)
		}
		p.dr.AddPackages(packageName, versions)

	}
}

func (p Process) getRequire(require map[string]string) map[string]PackageInfo {
	packages := map[string]PackageInfo{}
	for packageName, rule := range require {

		p.dr.AddRequirement(packageName, rule)

		if _, ok := p.Packages[packageName]; ok {
			continue
		}

		if packageName == "php" || strings.HasPrefix(packageName, "ext-") || strings.HasPrefix(packageName, "lib-") {
			continue
		}

		packageInfo, err := p.PackageRepo.Find(packageName)

		if err != nil {
			// Todo improve error handling
			panic(err)
		}
		packages[packageName] = packageInfo

	}
	return packages
}
