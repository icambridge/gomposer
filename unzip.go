package main

import (
	"archive/zip"
	"log"
	"os"
	"strings"
	"path/filepath"
	"io"
)

func main() {
	// Iterate through the files in the archive,
	// printing some of their contents.
	pkgName := "alexandresalome/PHP-Selenium"
	dirs := strings.Split(pkgName, "/")

	dirName := "vendors"
	for _, k := range dirs {
		dirName = dirName + "/" + k
		os.Mkdir(dirName, 0744)
	}

	Extract(dirName, "output.zip")
}


