package gomposer

import (
    "os"

    "encoding/json"
)

type Reader interface {
	Read(filename string) *Version
}

type PackageReader struct {
}

func (pr PackageReader) Read(filename string) (*Version, error) {

    buf, err := os.Open(filename)

    if err != nil {
        return nil, err
    }

    output := &Version{}

    json.NewDecoder(buf).Decode(output)

	return output, nil
}
