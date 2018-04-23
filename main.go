package main

import (
	"context"

	"github.com/google/go-github/github"
)

func main() {
	client := github.NewClient(nil)

	opt := &github.RepositoryListByOrgOptions{Type: "public"}
	repos, _, err := client.Repositories.ListByOrg(context.Background(), "github", opt)

}
