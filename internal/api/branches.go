package api

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// FetchBranches fetch the branches in the repository.
func (client *Client) FetchBranches(ctx context.Context, owner string, repo string) ([]model.Branch, error) {
	opt := github.ListOptions{}

	items := []model.Branch{}

	for {
		branches, resp, err := client.github.Repositories.ListBranches(ctx, owner, repo, &opt)
		if err != nil {
			return items, err
		}

		for _, branch := range branches {
			items = append(items, model.ConvertBranch(owner, repo, branch))
		}

		if resp.NextPage == 0 {
			return items, nil
		}

		opt.Page = resp.NextPage
	}
}
