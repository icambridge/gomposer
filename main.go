package main

import (
	"fmt"
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

				fmt.Println(fmt.Sprintf("%v", actual))
				m := make(map[string]*gomposer.PackageInfo)
				hc, _ := gomposer.NewHttpClient("https://packagist.org/packages/")
				pr := gomposer.PackageRepository{Client: hc}

				p := gomposer.Process{PackageRepo: &pr, Packages: m}
				l := p.Process(actual)

				for _, v := range l.Packages {
					fmt.Println(v.Name + " " + v.Version)
				}

			},
		},
	}

	app.Run(os.Args)
}
