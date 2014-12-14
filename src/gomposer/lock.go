package gomposer

type LockGenerator struct {
	PackageRepo PackageRepository
}

func (lg LockGenerator) Generate(dependencies map[string]string) (Lock) {

	l := Lock{}
	p := []Version{}
	for k, v := range dependencies {
		p, err := lg.PackageRepo.Find(k)
		if err != nil{
			// TODO remove
			panic(err)
		}
		p = append(p, p.Versions[v])
	}

	return l
}
