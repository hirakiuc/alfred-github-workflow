package api

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// FetchReposByOwner fetch the repos.
func (client *Client) FetchReposByOwner(ctx context.Context, owner string) ([]model.Repo, error) {
	opt := &github.RepositoryListOptions{
		Visibility: "public",
	}

	items := []model.Repo{}

	for {
		repos, resp, err := client.github.Repositories.List(ctx, owner, opt)
		if err != nil {
			return items, err
		}

		for _, repo := range model.ConvertRepos(repos) {
			items = append(items, repo)
		}

		hasNext := (resp.NextPage != 0)
		if hasNext != true {
			return items, nil
		}

		opt.Page = resp.NextPage
	}
}
