package gomposer

import (
	"encoding/json"
	"fmt"
	"strings"
	"github.com/icambridge/go-dependency"
	"os"
//	"io/ioutil"
)

// TODO reanme
type PackageRepository struct {
	Client   *HttpClient
	Packages map[string]map[string]Version
}

// TODO remove
func (r PackageRepository) Find(packageName string) (*PackageInfo, error) {
	packageName = strings.ToLower(packageName)
	prepFilename := strings.NewReplacer("/", "$").Replace(packageName)
//	err := nil
	filename := fmt.Sprintf("/Users/icambridge/.composer/cache/repo/https---packagist.org/provider-%s.json", prepFilename)


	output := &PackageDetail{}
	if _, found := os.Stat(filename); os.IsNotExist(found) {

		err := r.Client.Request("GET", "/"+packageName+".json", output)

		return &output.PackageData, err
	}

	cached := &PackageCache{}
	buf, err := os.Open(filename)

		if err != nil {
			return nil, err
		}

	err = json.NewDecoder(buf).Decode(cached)
	output.PackageData = PackageInfo{Versions: cached.PackageData[packageName]}
	return &output.PackageData, err
}

func (r PackageRepository) Get(packageName string) (map[string]dependency.Dependency, error) {

	if r.Packages == nil {
		r.Packages = make(map[string]map[string]Version)
	}

	start, ok := r.Packages[packageName]

	if !ok {
		packageInfo, err := r.Find(packageName)

		if err != nil {
			return nil, err
		}

		start = packageInfo.Versions
	}
	m := map[string]dependency.Dependency{}
	for k, v := range start {

		m[k] = ToDependency(&v)

	}

	return m, nil

}

func ToDependency(pi *Version) dependency.Dependency {
	requires := map[string]string{}

	for reqPackageName, reqPackageVersion := range pi.Require {
		if !IsPackagist(reqPackageName) {
			continue
		}
		requires[reqPackageName] = reqPackageVersion
	}
	// todo move to function
//	for reqPackageName, reqPackageVersion := range pi.RequireDev {
//		if !IsPackagist(reqPackageName) {
//			continue
//		}
//		requires[reqPackageName] = reqPackageVersion
//	}
	return dependency.Dependency{
			Name: pi.Name,
			Version: pi.Version,
			Requires: requires,
			Replaces: pi.Replace,
		}
}

func IsPackagist(name string) bool {
	return strings.Contains(name, "/")
}

type Lock struct {
	Packages    []Version `json:"packages"`
	PackagesDev []Version `json:"packages-dev"`
}

type PackageDetail struct {
	PackageData PackageInfo `json:"package"`
}
type PackageCache struct {
	PackageData map[string]map[string]Version `json:"packages"`
}

type PackageInfo struct {
	Name        string             `json:"name"`
	Description string             `json:"description"`
	Time        string             `json:"time"`
	Maintainers []Maintainer       `json:"maintainers"`
	Versions    map[string]Version `json:"versions"`
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

// TODO rename since this going to be the name composer.json
type Version struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	// Keywords          []string          `json:"keywords"`
	// Homepage          string            `json:"homepage"`
	Version string `json:"version"`
	// VersionNormalized string            `json:"version_normalized"`
	//	License           []string          `json:"license"`
	Authors []Author `json:"authors"`
	Source  Source   `json:"source"`
	Dist    Source   `json:"dist"`
	// Type              string            `json:"type"`
	// Time              string            `json:"time"`
	// Autoload          Autoload          `json:"autoload"`
	Require    map[string]string `json:"require"`
	Replace    map[string]string `json:"replace"`
	RequireDev map[string]string `json:"require-dev"`
	Suggest    map[string]string `json:"suggest"`
}
