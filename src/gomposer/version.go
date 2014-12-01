package gomposer

import (
	"github.com/mcuadros/go-version"
)

type Comparer interface {
	GreaterThan(v1 string, v2 string) bool
	LessThan(v1 string, v2 string) bool
	GreaterThanOrEqual(v1 string, v2 string) bool
	LessThanOrEqual(v1 string, v2 string) bool
	Equal(v1 string, v2 string) bool
}

type VersionComparer struct {
}

func (vc VersionComparer) GreaterThan(v1 string, v2 string) bool {
	return version.Compare(v1, v2, ">")
}

func (vc VersionComparer) LessThan(v1 string, v2 string) bool {
	return version.Compare(v1, v2, "<")
}
func (vc VersionComparer) GreaterThanOrEqual(v1 string, v2 string) bool {
	return version.Compare(v1, v2, ">=")
}

func (vc VersionComparer) LessThanOrEqual(v1 string, v2 string) bool {
	return version.Compare(v1, v2, "<=")
}

func (vc VersionComparer) Equal(v1 string, v2 string) bool {
	return version.Compare(v1, v2, "==")
}
