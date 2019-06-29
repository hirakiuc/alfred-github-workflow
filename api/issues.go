package api

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

func filterIssues(issues []*github.Issue) []*github.Issue {
	results := []*github.Issue{}

	for _, issue := range issues {
		if issue.GetPullRequestLinks() == nil {
			results = append(results, issue)
		}
	}

	return results
}

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

		// Remove PullRequests
		issues = filterIssues(issues)

		items = append(items, model.ConvertIssues(issues)...)

		if resp.NextPage == 0 {
			return items, nil
		}

		opt.Page = resp.NextPage
	}
}
