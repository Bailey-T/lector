package main

import (
	"log"

	"github.com/google/go-github/v45/github"

	"github.com/drtbz/lector/sources"
	"github.com/drtbz/lector/sources/gh"
	"github.com/drtbz/lector/sources/artifacthub"
)

func main() {

	ctx, client := gh.Setup()

	repoList := sources.ReadRepos("repolist.txt")
	artifacthub.GetPackage()
	// just the latest release
	for _, v := range repoList {
		re, _, err := gh.LatestRelease(client, ctx, v.Owner(), v.Repo())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v/%v: %v", v.Owner(), v.Repo(), github.Stringify(re.Name))
	}
	// get all pages of results
	// listReleases(client, ctx, "Azure", "AKS")

}