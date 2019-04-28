package api

import (
	"context"

	"github.com/google/go-github/github"
)

// FetchIssuesHandler describe a handler interface
type FetchIssuesHandler func(issues []*github.Issue, err error, hasNext bool) bool

// FetchIssues fetch the issues in the repository.
func (client *Client) FetchIssues(ctx context.Context, owner string, repo string, handler FetchIssuesHandler) {
	opt := github.IssueListByRepoOptions{
		State:     "open",
		Sort:      "created",
		Direction: "desc",
	}

	for {
		issues, resp, err := client.github.Issues.ListByRepo(ctx, owner, repo, &opt)
		if err != nil {
			handler([]*github.Issue{}, err, false)
			return
		}

		hasNext := (resp.NextPage != 0)
		if handler(issues, nil, hasNext) != true {
			return
		}

		if hasNext != true {
			return
		}

		opt.Page = resp.NextPage
	}
}
