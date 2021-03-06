package api

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

func (client *Client) FetchPullsCreated(ctx context.Context, user string) ([]model.Issue, error) {
	queryParams := []string{
		"is:pr",
		"is:open",
		fmt.Sprintf("author:%s", user),
		"archived:false",
	}
	q := strings.Join(queryParams, " ")

	opt := github.SearchOptions{
		Sort:  "updated",
		Order: "desc",
	}

	return client.SearchIssues(ctx, q, &opt)
}
