package gomposer

import (
	"os"

	"encoding/json"
)

type PackageReader struct {
}

func (pr PackageReader) Read(filename string) (ComposerPackage, error) {

	buf, err := os.Open(filename)

	output := ComposerPackage{}
	if err != nil {
		return output, err
	}

	json.NewDecoder(buf).Decode(&output)

	return output, nil
}

