package api

import (
	"context"

	"github.com/google/go-github/github"
)

// FetchPullsHandler describe a handler interface
type FetchPullsHandler func(pulls []*github.PullRequest, err error, hasNext bool) bool

// FetchPulls fetch the pull requests in the repository.
func (client *Client) FetchPulls(ctx context.Context, owner string, repo string, handler FetchPullsHandler) {
	opt := github.PullRequestListOptions{
		State:     "open",
		Sort:      "created",
		Direction: "desc",
	}

	for {
		pulls, resp, err := client.github.PullRequests.List(ctx, owner, repo, &opt)
		if err != nil {
			handler([]*github.PullRequest{}, err, false)
			return
		}

		hasNext := (resp.NextPage != 0)
		if handler(pulls, nil, hasNext) != true {
			return
		}

		if hasNext != true {
			return
		}

		opt.Page = resp.NextPage
	}
}
