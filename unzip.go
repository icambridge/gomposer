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

func Extract(dirName, zipFile string) {

	r, err := zip.OpenReader(zipFile)
	if err != nil {
		log.Fatal(err)
	}
	defer r.Close()
	for _, f := range r.File {

		fileName := filepath.Base(f.Name)
		if fileName == "." {
			continue
		}
		currentDir := dirName

		rc, err := f.Open()
		if err != nil {
			log.Fatal(err)
		}

		parts := strings.Split(filepath.Dir(f.Name), "/")
		parts = parts[1:]

		for _, k := range parts {
			currentDir = currentDir + "/" + k
			os.Mkdir(currentDir, 0744)
		}
		partCount := len(parts)

		if partCount > 0 && parts[partCount-1] == fileName {
			continue
		}

		of, err := os.Create(currentDir + "/" + fileName)
		if err != nil {
			log.Fatal(err)
		}
		_, err = io.Copy(of, rc)

		if err != nil {
			log.Fatal(err)
		}
		rc.Close()
		of.Close()
	}
}

