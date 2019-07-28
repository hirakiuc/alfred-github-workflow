package api

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

func (client *Client) FetchReleases(ctx context.Context, owner string, repo string) ([]model.Release, error) {
	opt := github.ListOptions{}

	items := []model.Release{}

	for {
		rels, resp, err := client.github.Repositories.ListReleases(ctx, owner, repo, &opt)
		if err != nil {
			return items, err
		}

		items = append(items, model.ConvertReleases(rels)...)
		if resp.NextPage == 0 {
			return items, nil
		}

		opt.Page = resp.NextPage
	}
}
