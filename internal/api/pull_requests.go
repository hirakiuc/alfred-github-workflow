package api

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

type FetchPullRequestsHandler func(pulls []*github.PullRequest, err error, hasNext bool) bool

func (client *GithubClient) FetchPullRequests(owner string, repo string) ([]*github.PullRequest, error) {
	opt := &github.PullRequestListOptions{
		State: "open",
	}

	var pullRequests []*github.PullRequest
	for {
		fmt.Println("Fetch pulls!")

		pulls, resp, err := client.PullRequests.List(context.Background(), owner, repo, opt)
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
