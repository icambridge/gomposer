package gomposer


import (
	"github.com/mcuadros/go-version" // TODO remove the need for this to be imported all over the show
)

type Process struct {
	packageRepo *PackageRepository
	packages map[string]*PackageInfo
}

// TODO recrusive
func (p Process) Process(packageInfo *Version) {


	// TODO inject
	m := make(map[string][]*version.ConstraintGroup)
	v := make(map[string][]string)
	dr := DependencyResolver{requiredPackages: m, versions: v}

	for packageName, rule := range packageInfo.Require {
		dr.AddRequirement(packageName, rule)
	}

	p.getRequire(packageInfo.Require)

	// TODO move to seperate function

	for packageName, packageInfo := range p.packages {

		versions := make([]string, 0, len(packageInfo.Versions))
		for versionNum := range packageInfo.Versions {
			versions = append(versions, versionNum)
		}
		dr.AddPackages(packageName, versions)
	}

	requiredVersions := dr.Resolve()
	for packageName, versionNum := range requiredVersions {
			packageInfo := p.packages[packageName].Versions[versionNum]
			p.Process(&packageInfo)
	}
}

func (p Process) getRequire(require map[string]string) {
	for packageName, _ := range require {
		// todo check if we've already got this package
		packageInfo, err := p.packageRepo.Find(packageName)

		if err != nil {
			// Todo improve error handling
			panic(err)
		}
		p.packages[packageName] = packageInfo
	}
}
