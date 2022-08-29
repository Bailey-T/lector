package main

import (
	"context"

	"github.com/drtbz/lector/exporters/azuretable"
	"github.com/drtbz/lector/sources"
	"github.com/drtbz/lector/sources/artifacthub"
	ghsource "github.com/drtbz/lector/sources/github"
	"github.com/google/go-github/v45/github"
	"log"
)

func main() {

	ctx, client := ghsource.PATSetup("GHTOKEN")
	repoList := sources.GetUpstream("_data/repolist.txt")
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
	aztClient := azuretable.ConnectionStringSetup("DefaultEndpointsProtocol=https;AccountName=tbtfstate;AccountKey=qqQPkM11QPVpa4lrB0bbZo+m7BzOHdA7KIn51C1b0eehsuBAS+YjFBLyAOGMuReEU0WU4Z17EE7P+AStY1yrMQ==;EndpointSuffix=core.windows.net")
	resp, err := aztClient.CreateTable(context.TODO(), "fromServiceClient", nil)
	if err != nil {
		panic(err)
	}
	log.Printf("%v", resp)
}
