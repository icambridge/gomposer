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
				fmt.Println("Starting update process")
				r := gomposer.PackageReader{}
				actual, _ := r.Read("composer.json")


				fmt.Pritnln("Solving dependencies")
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
				fmt.Println("Downloading dependencies")
				lockGenerator := gomposer.LockGenerator{
					PackageRepo: pr,
				}
				os.Mkdir("vendors", 0777)
				lock := lockGenerator.Generate(required)
				for _, p := range lock.Packages {
					gomposer.Download(p)
				}
				gomposer.WriteLock(lock)
			},
		},
	}

	app.Run(os.Args)
}
