package gomposer

import (
	"github.com/icambridge/go-dependency"
	"os"
	"strings"
	"time"
)

type PackageRepository struct {
	Client   *HttpClient
}

func (r *PackageRepository) Find(packageName string) (PackageInfo, error) {

	packageName = strings.ToLower(packageName)

	filename := GetCacheFilename(packageName)
	then := time.Now().AddDate(0, -1, 0)
	output := PackageDetail{}
	if fi, found := os.Stat(filename); os.IsNotExist(found) || fi.ModTime().Before(then) {

		err := r.Client.Request("GET", "/"+packageName+".json", &output)
		// TODO remove cache in tests
		WriteCache(packageName, output.PackageData)

		return output.PackageData, err
	}

	cache, err := ReadCache(filename, packageName)
	output.PackageData = cache
	return output.PackageData, err
}

func (r PackageRepository) Get(packageName string) (map[string]dependency.Dependency, error) {

	m := map[string]dependency.Dependency{}
	packageInfo, err := r.Find(packageName)

	if err != nil {
		return m, err
	}

	for k, v := range packageInfo.Versions {

		m[k] = ToDependency(v)

	}
	r = r
	return m, nil

}

func ToDependency(pi ComposerPackage) dependency.Dependency {
	requires := map[string]string{}

	for reqPackageName, reqPackageVersion := range pi.Require {
		if !IsPackagist(reqPackageName) {
			continue
		}
		requires[reqPackageName] = reqPackageVersion
	}

	return dependency.Dependency{
		Name:     pi.Name,
		Version:  pi.Version,
		Requires: requires,
		Replaces: pi.Replace,
	}
}

func IsPackagist(name string) bool {
	return strings.Contains(name, "/")
}

type PackageDetail struct {
	PackageData PackageInfo `json:"package"`
}
type PackageCache struct {
	PackageData map[string]map[string]ComposerPackage `json:"packages"`
}

type PackageInfo struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Time        string             `json:"time"`
	Maintainers []Maintainer       `json:"maintainers"`
	Versions    map[string]ComposerPackage `json:"versions"`
	Type        string             `json:"type"`
	Repository  string             `json:"repository"`
	Downloads   Downloads          `json:"downloads"`
	Favers      int                `json:"favers"`
}

type Maintainer struct {
	Name string `json:"name"`
}

type Author struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type Source struct {
	Type      string `json:"type"`
	Url       string `json:"url"`
	Reference string `json:"reference"`
	Shasum    string `json:"shasum"`
}

type Downloads struct {
	Total   int `json:"total"`
	Monthly int `json:"monthly"`
	Daily   int `json:daily`
}

type Autoload struct {
	ClassMap []string          `json:"classmap"`
	Files    []string          `json:"files"`
	Psr0     map[string]string `json:"psr-0"`
	Psr4     map[string]string `json:"psr-4"`
}

type ComposerPackage struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	Keywords          []string          `json:"keywords"`
	Homepage          string            `json:"homepage"`
	Version string `json:"version"`
	VersionNormalized string            `json:"version_normalized"`
	License           []string          `json:"license"`
	Authors []Author `json:"authors"`
	Source  Source   `json:"source"`
	Dist    Source   `json:"dist"`
	Type              string            `json:"type"`
	Time              string            `json:"time"`
	Require    map[string]string `json:"require"`
	Replace    map[string]string `json:"replace"`
	RequireDev map[string]string `json:"require-dev"`
	Suggest    map[string]string `json:"suggest"`
}
