package gomposer

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"github.com/icambridge/cartel"
)

type DownloadOutput struct {
	
}

func (do DownloadOutput) Value() interface {} {
	return true
}

type DownloadTask struct {
	Version ComposerPackage
}

func (dt DownloadTask) Execute() cartel.OutputValue {
	v := dt.Version
	fmt.Println(fmt.Sprintf("Downloading %s", v.Name))
	s := GenerateRandomString(10)
	filename := os.TempDir() + "/" + s + "." + v.Dist.Type
	out, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	defer out.Close()

	buf := new(bytes.Buffer)
	req, err := http.NewRequest("GET", v.Dist.Url, buf)

	client := http.DefaultClient
	resp, err := client.Do(req)

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

	os.Remove(filename)
	
	return DownloadOutput{}
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

func Remove(vendorDir string, v ComposerPackage) {
	os.RemoveAll(vendorDir + "/" + v.Name)

	parts := strings.Split(v.Name, "/")
	os.Remove(vendorDir + "/" + parts[0])
}
