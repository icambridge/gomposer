package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/icambridge/go-dependency"
	"github.com/icambridge/gomposer/gomposer"
	"os"
)

func main() {

	app := cli.NewApp()
	app.Name = "gomposer"
	app.Usage = "Composer is go bro"
	app.Commands = []cli.Command{
		{
			Name:   "update",
			Usage:  "Updates",
			Action: Update,
		},
		{
			Name:   "install",
			Usage:  "Installs",
			Action: Install,
		},
	}

	app.Run(os.Args)
}

func Install(c *cli.Context) {

	lockFile := "composer.lock"
	if _, found := os.Stat(lockFile); os.IsNotExist(found) {
		Update(c)
		return
	}

	fmt.Println("Reading lock")
	lock, err := gomposer.ReadLock(lockFile)
	if err != nil {
		panic(err)
	}
	Download(lock.Packages)
}

func Update(c *cli.Context) {
	r := gomposer.PackageReader{}
	actual, _ := r.Read("composer.json")

	fmt.Println("Solving dependencies")
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
	newLock := lockGenerator.Generate(required)
	lockFile := "composer.lock"

	oldLock, err := gomposer.ReadLock(lockFile)
	if err != nil {
		panic(err)
	}

	diff := gomposer.DiffLock(newLock, oldLock)

	if len(diff["removed"]) != 0 {
		Remove(diff["removed"])
	}

	if len(diff["added"]) == 0 {
		fmt.Println("Nothing to update")
		return
	}

	Download(diff["added"])
	gomposer.WriteLock(newLock)
}

func Remove(packages []gomposer.Version) {
	fmt.Println("Removing outdated dependencies")
	for _, p := range packages {

		fmt.Println(fmt.Sprintf("Removing %s", p.Name))
		gomposer.Remove("vendors", p)
	}
}

func Download(packages []gomposer.Version) {

	os.Mkdir("vendors", 0777)
	fmt.Println("Downloading dependencies")
	for _, p := range packages {
		gomposer.Download(p)
	}
}
