package main

import (
	"context"
	"log"
	"os"

	"github.com/drtbz/lector/sources"
	"github.com/drtbz/lector/sources/artifacthub"
	//"github.com/drtbz/lector/sources/gh"
	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func main() {

	ctx, client := GitHubSetup()
	repoList := sources.GetUpstream("repolist.txt")
	for _, v := range repoList {
		if v.Source() == "github" {
			re, _, err := LatestRelease(client, ctx, v.Owner(), v.Repo())
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

func GitHubSetup() (context.Context, *github.Client) {
	token, tkex := os.LookupEnv("GHTOKEN")

	if !tkex {
		log.Fatal("Couldn't get token from ENV")
	}

	ctx := context.Background()
	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	client := github.NewClient(tc)
	return ctx, client
}

func LatestRelease(client *github.Client, ctx context.Context, owner, repo string) (release *github.RepositoryRelease, resp *github.Response, err error) {
	release, resp, err = client.Repositories.GetLatestRelease(ctx, owner, repo)
	return
}

func ListReleases(client *github.Client, ctx context.Context, owner, repo string, prerelease ...bool, ) {
	opt := &github.ListOptions{
		Page:    0,
		PerPage: 50,
	}
	var allRelease []*github.RepositoryRelease
	for {
		repos, resp, err := client.Repositories.ListReleases(ctx, owner, repo, opt)
		if err != nil {
			log.Fatalf("Error: %v", err)
		}
		allRelease = append(allRelease, repos...)
		log.Printf("FirstPage: %v ", resp.FirstPage)
		log.Printf("LastPage: %v ", resp.LastPage)
		log.Printf("NextPage: %v \n ", resp.NextPage)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	for _, v := range allRelease {
		if !*v.Prerelease {
			log.Printf("%v\n", github.Stringify(v.Name))
		}
	}
}