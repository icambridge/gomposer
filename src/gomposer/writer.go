package gomposer

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"strings"
)

func WriteLock(lock Lock) {

	b, err := json.Marshal(lock)

	if err != nil {
		fmt.Println("error:", err)
	}

	f, err := os.Create("composer.lock")

	f.Write(b)
}

func WriteCache(packageName string, versions map[string]Version) {
	cache := PackageCache{
		PackageData: map[string]map[string]Version{
			packageName: versions,
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
