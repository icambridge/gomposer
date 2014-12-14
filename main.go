package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/icambridge/go-dependency"
	"gomposer"
	"os"
	"sort"
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
				}

				names := []string{}
				for k, _ := range required {
					names = append(names, k)
				}
				sort.Sort(sort.StringSlice(names))

				for _, v := range names {
					fmt.Println(fmt.Sprintf("%s->%s", v, required[v]))
				}
				// TODO convert required into Lock file.
			},
		},
	}

	app.Run(os.Args)
}
