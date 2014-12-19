package gomposer

import (
	"github.com/mcuadros/go-version"
)

type VersionSlice []string

func (p VersionSlice) Len() int           { return len(p) }
func (p VersionSlice) Less(i, j int) bool { return version.Compare(p[i], p[j], "<") }
func (p VersionSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
