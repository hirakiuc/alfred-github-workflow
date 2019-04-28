package api

import (
	"context"

	"github.com/google/go-github/github"
)

// FetchPullRequestsHandler describe the handler for the FetchPullRequestsHandler method.
type FetchPullRequestsHandler func(pulls []*github.PullRequest, err error, hasNext bool) bool

// FetchPullRequests fetch the pull requests on the repository.
func (client *Client) FetchPullRequests(ctx context.Context, owner string, repo string) ([]*github.PullRequest, error) {
	opt := &github.PullRequestListOptions{
		State: "open",
	}

	var pullRequests []*github.PullRequest
	for {
		pulls, resp, err := client.github.PullRequests.List(ctx, owner, repo, opt)
		if err != nil {
			return []*github.PullRequest{}, err
		}

		pullRequests = append(pullRequests, pulls...)
		if resp.NextPage == 0 {
			break
		}

		opt.Page = resp.NextPage
	}

	return pullRequests, nil
}
