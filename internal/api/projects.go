package api

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
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

		for _, project := range model.ConvertProjects(projects) {
			items = append(items, project)
		}

		hasNext := (resp.NextPage != 0)
		if hasNext != true {
			return items, nil
		}

		opt.Page = resp.NextPage
	}
}
