package gomposer

import (
	"github.com/mcuadros/go-version" // TODO remove the need for this to be imported all over the show
)

type Process struct {
	packageRepo *PackageRepository
	packages    map[string]*PackageInfo
	dr          *DependencyResolver
}

// TODO recrusive
func (p Process) Process(packageInfo *Version) *Lock {

	// TODO inject
	m := make(map[string][]*version.ConstraintGroup)
	v := make(map[string][]string)

	if p.dr == nil {
		p.dr = &DependencyResolver{requiredPackages: m, versions: v}
	}

	p.inner(packageInfo.Require)
	// TODO solve this.
	requiredVersions := p.dr.Resolve()
	l := &Lock{Packages: []Version{}}

	for packageName, versionNum := range requiredVersions {
		packageInfo := p.packages[packageName].Versions[versionNum]
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
		packageInfo := p.packages[packageName].Versions[versionNum]
		p.inner(packageInfo.Require)
	}
}

func (p Process) addPackages(packages map[string]*PackageInfo) {

	for packageName, packageInfo := range packages {
		p.packages[packageName] = packageInfo
		versions := make([]string, 0, len(packageInfo.Versions))
		for versionNum := range packageInfo.Versions {
			versions = append(versions, versionNum)
		}
		p.dr.AddPackages(packageName, versions)
	}
}

func (p Process) getRequire(require map[string]string) map[string]*PackageInfo {
	packages := map[string]*PackageInfo{}
	for packageName, rule := range require {

		p.dr.AddRequirement(packageName, rule)

		if _, ok := p.packages[packageName]; ok {
			continue
		}
		packageInfo, err := p.packageRepo.Find(packageName)

		if err != nil {
			// Todo improve error handling
			panic(err)
		}
		packages[packageName] = packageInfo

	}
	return packages
}
