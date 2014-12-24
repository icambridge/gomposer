package gomposer

import (
	"sort"
	"encoding/json"
	"fmt"
	"os"
)

type Lock struct {
	Packages    []ComposerPackage `json:"packages"`
	PackagesDev []ComposerPackage `json:"packages-dev"`
}

type LockGenerator struct {
	PackageRepo PackageRepository
}

func (lg LockGenerator) Generate(dependencies map[string]string) (Lock, error) {

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
			return l, err
		}
		l.Packages = append(l.Packages, p.Versions[v])
	}

	return l, nil
}

func DiffLock(new, old Lock) map[string][]ComposerPackage {
	added := []ComposerPackage{}
	removed := []ComposerPackage{}

	oldPackages := map[string]ComposerPackage{}

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

	return map[string][]ComposerPackage{"added": added, "removed": removed}
}

func WriteLock(lock Lock) {

	b, err := json.Marshal(lock)

	if err != nil {
		fmt.Println("error:", err)
	}

	f, err := os.Create("composer.lock")

	if err != nil {
		fmt.Println("error:", err)
	}
	f.Write(b)
}

func ReadLock(filename string) (Lock, error) {

	output := Lock{}
	buf, err := os.Open(filename)


	if err != nil {
		if os.IsNotExist(err) {
			err = nil
			output.Packages = []ComposerPackage{}
			output.PackagesDev = []ComposerPackage{}
		}
		return output, err
	}

	json.NewDecoder(buf).Decode(&output)

	return output, nil
}
