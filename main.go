package main

import (
	"log"

	"github.com/drtbz/lector/exporters/azuretable"
	"github.com/drtbz/lector/sources"
	"github.com/drtbz/lector/sources/artifacthub"
	ghsource "github.com/drtbz/lector/sources/github"
	"github.com/google/go-github/v45/github"
)

func main() {

	serviceClient := azuretable.ConnectionStringSetup("DefaultEndpointsProtocol=https;AccountName=drtbzlector;AccountKey=BUikGQr8ujKUv+Q5dvhd7NC5jvDK9QSxGdzdPnXthuPcR89cxP+CYrg9POQcKs5WbY1Et4GBAMGA+AStQxhnLw==;EndpointSuffix=core.windows.net")
	if azuretable.Setup("lector", serviceClient) != nil {
		log.Fatalf("Error Setting up Azure Tables")
	}
	aztClient := serviceClient.NewClient("lector")

	ctx, client := ghsource.PATSetup("GHTOKEN")
	repoList := sources.GetUpstream("_data/repolist.txt")
	for _, v := range repoList {
		if v.Source() == "github" {
			re, _, err := ghsource.LatestRelease(client, ctx, v.Owner(), v.Repo())
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%v/%v: %v", v.Owner(), v.Repo(), github.Stringify(re.Name))
			//azuretable.NewGHEntity(aztClient, v, github.Stringify(re.Name))
			azuretable.Query(aztClient, v)
		}
		if v.Source() == "artifacthub" {
			re, err := artifacthub.GetHelmPackage(v.Owner(), v.Repo())
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("%v: Chart: %v, App: %v \n", re["name"], re["version"], re["app_version"])
			//azuretable.NewAHEntity(aztClient, v, fmt.Sprintf("%v", re["version"]), fmt.Sprintf("%v", re["app_version"]))
			azuretable.Query(aztClient, v)

		}
	}

}
