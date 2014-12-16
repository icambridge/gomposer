package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/icambridge/go-dependency"
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

				d := gomposer.ToDependency(actual)

				hc, _ := gomposer.NewHttpClient("https://packagist.org/packages/")
				pr := gomposer.PackageRepository{Client: hc}

				repo := dependency.GetNewRepo(pr)

				ads := dependency.GetPackageNames(d)
				repo.GetAll(ads)

				s := dependency.NewSolver(repo.Dependencies, repo.Replaces)
				required, err := s.Solve(d)

				if err != nil {
					fmt.Println(err)
					return
				}

				// TODO convert required into Lock file.

				lockGenerator := gomposer.LockGenerator{
					PackageRepo: pr,
				}
				os.Mkdir("vendors", 0777)
				lock := lockGenerator.Generate(required)
				gomposer.Download(lock.Packages[0])
			},
		},
	}

	app.Run(os.Args)
}
