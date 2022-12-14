package sources

import (
	"bufio"
	"log"
	"os"
	"strings"
)

type Upstream struct {
	// Declares the fields for an upstream source to check
	owner  string // Azure, Hashicorp, Elastic etc.
	repo   string // AKS, Terraform, MetricBeat etc.
	source string // Github, ArtifactHub
}

func GetUpstream(filePath string) (repoList []Upstream) {
	// open file
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read the file line by line using scanner
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		st := strings.ToLower(strings.Replace(scanner.Text(), "https://", "", 1))
		if strings.Contains(st, "github") {
			// go over each line and split on the slash
			sp := strings.Split(strings.Replace(st, "github.com/", "", 1), "/")
			// append as a repo object to the array
			repoList = append(repoList, Upstream{owner: sp[0], repo: sp[1], source: "github"})
		}
		if strings.Contains(st, "artifacthub") {
			sp := strings.Split(strings.Replace(st, "artifacthub.io/", "", 1), "/")
			repoList = append(repoList, Upstream{owner: sp[2], repo: sp[3], source: "artifacthub"})
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return repoList
}

func (u Upstream) Owner() (owner string) {
	return u.owner
}

func (u Upstream) Repo() (repo string) {
	return u.repo
}

func (u Upstream) Source() (repo string) {
	return u.source
}
