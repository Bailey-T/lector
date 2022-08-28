package main

import (
	//"context"
	"log"
	//"os"
	"github.com/drtbz/lector/sources"
	aztexport "github.com/drtbz/lector/exporters/azuretable"
	"github.com/drtbz/lector/sources/artifacthub"
	ghsource "github.com/drtbz/lector/sources/github"
	"github.com/google/go-github/v45/github"
)

func main() {

	ctx, client := ghsource.PATSetup()
	repoList := sources.GetUpstream("repolist.txt")
	for _, v := range repoList {
		if v.Source() == "github" {
			re, _, err := ghsource.LatestRelease(client, ctx, v.Owner(), v.Repo())
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%v/%v: %v", v.Owner(), v.Repo(), github.Stringify(re.Name))
		}
		if v.Source() == "artifacthub" {
			re, err := artifacthub.GetHelmPackage(v.Owner(), v.Repo())
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%v: Chart: %v, App: %v \n", re["name"], re["version"], re["app_version"])

		}
	}
}
