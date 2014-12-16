package gomposer

import (
	"net/http"
	"io"
	"os"
	"os/exec"
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

	cmd := exec.Command("unzip", "output.zip", "-d", "vendors/"+v.Name)
	err = cmd.Start()
	if err != nil {
		panic(err)
	}

	//os.Remove("output.zip")
}
