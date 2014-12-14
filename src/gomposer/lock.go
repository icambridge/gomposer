package gomposer

import (
	"sort"
)

type LockGenerator struct {
	PackageRepo PackageRepository
}

func (lg LockGenerator) Generate(dependencies map[string]string) (Lock) {

	l := Lock{}
	packages := []string{}
	for k, _ := range dependencies {
		packages = append(packages, k)
	}

	sort.Sort(sort.StringSlice(packages))

	for _, k := range packages {
		p, err := lg.PackageRepo.Find(k)
		v := dependencies[k]
		if err != nil{
			// TODO remove
			panic(err)
		}
		l.Packages = append(l.Packages, p.Versions[v])
	}

	return l
}
