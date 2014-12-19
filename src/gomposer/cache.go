package gomposer

import (
	"encoding/json"
	"os"
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
