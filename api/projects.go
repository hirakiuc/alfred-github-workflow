package api

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// FetchProjects fetch the project cards in the repository.
func (client *Client) FetchProjects(ctx context.Context, owner string, repo string) ([]model.Project, error) {
	opt := github.ProjectListOptions{}

	items := []model.Project{}

	for {
		projects, resp, err := client.github.Repositories.ListProjects(ctx, owner, repo, &opt)
		if err != nil {
			return items, err
		}

		items = append(items, model.ConvertProjects(projects)...)

		if resp.NextPage == 0 {
			return items, nil
		}

		opt.Page = resp.NextPage
	}
}
