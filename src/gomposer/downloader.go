package gomposer

import (
	"net/http"
	"io"
	"os"
)
//
//type Downloader struct {
//
//}

func Download(v Version) {

	out, err := os.Create("output."+ v.Dist.Type)
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
	if err != nil {
		panic(err)
	}
}
