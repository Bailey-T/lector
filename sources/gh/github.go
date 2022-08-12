package gh

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/google/go-github/v45/github"
	"golang.org/x/oauth2"
)

func Setup() (context.Context, *github.Client) {
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

// func latestRelease(client *github.Client, ctx context.Context, owner,repo string) (*github.RepositoryRelease, *github.Response, error) {
// 	release, resp, err := client.Repositories.GetLatestRelease(ctx, owner, repo)
// 	return release, resp, err
// }

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
		fmt.Printf("FirstPage: %v ", resp.FirstPage)
		fmt.Printf("LastPage: %v ", resp.LastPage)
		fmt.Printf("NextPage: %v \n ", resp.NextPage)
		if resp.NextPage == 0 {
			break
		}
		opt.Page = resp.NextPage
	}

	for _, v := range allRelease {
		if !*v.Prerelease {
			fmt.Printf("%v\n", github.Stringify(v.Name))
		}
	}
}

func LatestRelease(client *github.Client, ctx context.Context, owner, repo string)(release *github.RepositoryRelease, resp *github.Response, err error) {
	release, resp, err = client.Repositories.GetLatestRelease(ctx, owner, repo)
	return
}
