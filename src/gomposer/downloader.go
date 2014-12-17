package gomposer

import (
	"net/http"
	"io"
	"os"
	"archive/zip"
	"log"
	"strings"
	"fmt"
	"path/filepath"
)
//
//type Downloader struct {
//
//}

func Download(v Version) {

	fmt.Println(fmt.Sprintf("Downloading %s", v.Name))
	s := GenerateRandomString(10)
	filename := s + "." + v.Dist.Type
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	resp, err := http.Get(v.Dist.Url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	_, err = io.Copy(out, resp.Body)


	dirs := strings.Split(v.Name, "/")

	dirName := "vendors"
	for _, k := range dirs {
		dirName = dirName + "/" + k
		os.Mkdir(dirName, 0744)
	}

	Extract(dirName, filename)

	//os.Remove(filename)
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
