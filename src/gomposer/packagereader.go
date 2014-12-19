package gomposer

import (
	"os"

	"encoding/json"
)

type Reader interface {
	Read(filename string) (*Version, error)
}

type PackageReader struct {
}

func (pr PackageReader) Read(filename string) (Version, error) {

	buf, err := os.Open(filename)

	output := Version{}
	if err != nil {
		return output, err
	}

	json.NewDecoder(buf).Decode(&output)

	return output, nil
}

func ReadLock(filename string) (*Lock, error) {

	output := &Lock{}
	buf, err := os.Open(filename)

	if err != nil {
		return output, err
	}

	json.NewDecoder(buf).Decode(output)

	return output, nil
}
