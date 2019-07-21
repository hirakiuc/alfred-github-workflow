package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

func (client *Client) SearchIssues(
	ctx context.Context, query string, opt *github.SearchOptions) ([]model.Issue, error) {
	items := []model.Issue{}

	for {
		result, resp, err := client.github.Search.Issues(ctx, query, opt)
		if err != nil {
			return items, err
		}

		for _, issue := range result.Issues {
			v := issue
			items = append(items, model.ConvertIssue(&v))
		}

		if resp.NextPage == 0 {
			return items, nil
		}

		opt.Page = resp.NextPage
	}
}

func (client *Client) FetchReviewRequests(ctx context.Context, user string) ([]model.Issue, error) {
	queryParams := []string{
		"is:pr",
		"is:open",
		fmt.Sprintf("review-requested:%s", user),
		"archived:false",
	}
	q := strings.Join(queryParams, " ")

	opt := github.SearchOptions{
		Sort:  "updated",
		Order: "desc",
	}

	return client.SearchIssues(ctx, q, &opt)
}
