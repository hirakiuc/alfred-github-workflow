package api

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// FetchIssues fetch the issues in the repository.
func (client *Client) FetchIssues(ctx context.Context, owner string, repo string) ([]model.Issue, error) {
	opt := github.IssueListByRepoOptions{
		State:     "open",
		Sort:      "created",
		Direction: "desc",
	}

	items := []model.Issue{}

	for {
		issues, resp, err := client.github.Issues.ListByRepo(ctx, owner, repo, &opt)
		if err != nil {
			return items, err
		}

		for _, issue := range model.ConvertIssues(issues) {
			items = append(items, issue)
		}

		hasNext := (resp.NextPage != 0)
		if hasNext != true {
			return items, nil
		}

		opt.Page = resp.NextPage
	}
}
