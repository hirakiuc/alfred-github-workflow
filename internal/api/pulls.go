package api

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// FetchPulls fetch the pull requests in the repository.
func (client *Client) FetchPulls(ctx context.Context, owner string, repo string) ([]model.PullRequest, error) {
	opt := github.PullRequestListOptions{
		State:     "open",
		Sort:      "created",
		Direction: "desc",
	}

	items := []model.PullRequest{}

	for {
		pulls, resp, err := client.github.PullRequests.List(ctx, owner, repo, &opt)
		if err != nil {
			return items, err
		}

		for _, pull := range model.ConvertPullRequests(pulls) {
			items = append(items, pull)
		}

		hasNext := (resp.NextPage != 0)
		if hasNext != true {
			return items, nil
		}

		opt.Page = resp.NextPage
	}
}
