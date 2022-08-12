package main

import (
	"log"

	"github.com/google/go-github/v45/github"

	"github.com/drtbz/lector/sources"
	"github.com/drtbz/lector/sources/gh"
)

func main() {

	ctx, client := gh.Setup()

	repoList := sources.ReadRepos("repolist.txt")



	// just the latest release
	for _, v := range repoList {
		re, _, err := gh.LatestRelease(client, ctx, v.OwnerName(), v.RepoName())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("%v/%v: %v", v.OwnerName(), v.RepoName(), github.Stringify(re.Name))
	}
	// get all pages of results
	// listReleases(client, ctx, "Azure", "AKS")

}

