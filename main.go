package main

import (
	"fmt"
	"sort"
	"github.com/codegangsta/cli"
	"gomposer"
	"os"
)

func main() {
	app := cli.NewApp()
	app.Name = "gomposer"
	app.Usage = "Composer is go bro"
	app.Commands = []cli.Command{
		{
			Name:  "update",
			Usage: "Updates",
			Action: func(c *cli.Context) {

				r := gomposer.PackageReader{}
				actual, _ := r.Read("composer.json")

				m := make(map[string]*gomposer.PackageInfo)
				hc, _ := gomposer.NewHttpClient("https://packagist.org/packages/")
				pr := gomposer.PackageRepository{Client: hc}

				p := gomposer.Process{PackageRepo: &pr, Packages: m}
				l := p.Process(actual)

				names := []string{}
				version := map[string]string{}
				for _, v := range l.Packages {
					names = append(names, v.Name)
					version[v.Name] = v.Version
				}

				sort.Strings(names)
				for _, name := range names {
					fmt.Println(name + "->" + version[name])
				}

			},
		},
	}

	app.Run(os.Args)
}
