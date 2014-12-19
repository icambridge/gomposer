package gomposer

import (
	"encoding/json"
	"os"
	"fmt"
	"os/user"
	"strings"
)

func ReadCache(filename, packageName string) (PackageInfo, error) {
	output := PackageInfo{}
	cached := PackageCache{}
	buf, err := os.Open(filename)

	if err != nil {
		return output, err
	}
	err = json.NewDecoder(buf).Decode(&cached)
	output.Versions = cached.PackageData[packageName]

	return output, nil
}

func WriteCache(packageName string, packageInfo PackageInfo) {
	cache := PackageCache{
		PackageData: map[string]map[string]Version{
			packageName: packageInfo.Versions,
		},
	}
	filename := GetCacheFilename(packageName)
	b, err := json.Marshal(cache)

	if err != nil {
		fmt.Println("error:", err)
	}

	f, err := os.Create(filename)

	f.Write(b)
}

func GetCacheFilename(packageName string) string {
	usr, err := user.Current()
	if err != nil {
		return ""
	}
	prepFilename := strings.NewReplacer("/", "$").Replace(packageName)
	directory := ".composer/cache/repo/https---packagist.org"
	filename := fmt.Sprintf("%s/%s/provider-%s.json", usr.HomeDir, directory, prepFilename)

	return filename
}
