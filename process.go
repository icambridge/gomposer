package gomposer

type Process struct {
    packageRepo *PackageRepository
}

func (p Process) Process(packageInfo *Version) {

    for packageName, _ := range packageInfo.Require {
        p.packageRepo.Find(packageName)
    }

}
