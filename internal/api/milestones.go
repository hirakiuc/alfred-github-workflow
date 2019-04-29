package api

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// FetchMilestones fetch the milestones in the repository.
func (client *Client) FetchMilestones(ctx context.Context, owner string, repo string) ([]model.Milestone, error) {
	opt := github.MilestoneListOptions{}

	items := []model.Milestone{}

	for {
		milestones, resp, err := client.github.Issues.ListMilestones(ctx, owner, repo, &opt)
		if err != nil {
			return items, err
		}

		for _, milestone := range model.ConvertMilestones(milestones) {
			items = append(items, milestone)
		}

		hasNext := (resp.NextPage != 0)
		if hasNext != true {
			return items, nil
		}

		opt.Page = resp.NextPage
	}
}
