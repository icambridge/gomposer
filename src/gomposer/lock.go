package gomposer

import (
	"sort"
)

type Lock struct {
	Packages    []Version `json:"packages"`
	PackagesDev []Version `json:"packages-dev"`
}

type LockGenerator struct {
	PackageRepo PackageRepository
}

func (lg LockGenerator) Generate(dependencies map[string]string) Lock {

	l := Lock{}
	packages := []string{}
	for k, _ := range dependencies {
		packages = append(packages, k)
	}

	sort.Sort(sort.StringSlice(packages))

	for _, k := range packages {
		p, err := lg.PackageRepo.Find(k)
		v := dependencies[k]
		if err != nil {
			// TODO remove
			panic(err)
		}
		l.Packages = append(l.Packages, p.Versions[v])
	}

	return l
}

func DiffLock(new, old Lock) map[string][]Version {
	added := []Version{}
	removed := []Version{}

	oldPackages := map[string]Version{}

	for _, pkgInfo := range old.Packages {
		oldPackages[pkgInfo.Name] = pkgInfo
	}

	for _, nPkgInfo := range new.Packages {
		oPkgInfo, found := oldPackages[nPkgInfo.Name]

		if !found || nPkgInfo.Version != oPkgInfo.Version {
			if found {
				removed = append(removed, oPkgInfo)
			}
			added = append(added, nPkgInfo)
		}

		if found {
			delete(oldPackages, nPkgInfo.Name)
		}
	}

	for _, pkgInfo := range oldPackages {
		removed = append(removed, pkgInfo)
	}

	return map[string][]Version{"added": added, "removed": removed}
}
