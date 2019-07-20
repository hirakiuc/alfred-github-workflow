package api

import (
	"context"
	"strings"

	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// SearchPulls fetch the pull requests which the user was requested to review.
func (client *Client) SearchPulls(ctx context.Context, user string) ([]model.Issue, error) {
	queryParams := []string{
		"is:pr",
		"review-",
	}
	q := strings.Join(queryParams, "+")

	opt := github.SearchOptions{}

	items := []model.Issue{}

	for {
		issueResult, resp, err := client.github.Search.Issues(ctx, q, &opt)
		if err != nil {
			return items, err
		}

		// items = append(items, model.ConvertIssueSearchResult(issueResult)...)
		for _, issue := range issueResult.Issues {
			v := issue
			items = append(items, model.ConvertIssue(&v))
		}

		if resp.NextPage == 0 {
			return items, nil
		}

		opt.Page = resp.NextPage
	}
}
