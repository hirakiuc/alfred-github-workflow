package api

import (
	"context"
	"fmt"

	"github.com/google/go-github/github"
)

// FetchReposHandler describe a handler interface
type FetchReposHandler func(repos []*github.Repository, err error, hasNext bool) bool

// FetchReposByUser fetch and return the all of repos.
func (client *Client) FetchReposByUser(ctx context.Context, user string) ([]*github.Repository, error) {
	opt := &github.RepositoryListOptions{Visibility: "public"}

	var repositories []*github.Repository
	for {
		fmt.Println("Fetch repositories!")

		repos, resp, err := client.github.Repositories.List(ctx, user, opt)
		if err != nil {
			return []*github.Repository{}, err
		}
		repositories = append(repositories, repos...)
		if resp.NextPage == 0 {
			break
		}

		opt.Page = resp.NextPage
	}

	return repositories, nil
}

// FetchReposByUserWithHandler fetch and return the repos.
func (client *Client) FetchReposByUserWithHandler(ctx context.Context, user string, handler FetchReposHandler) {
	opt := &github.RepositoryListOptions{
		Visibility: "public",
	}

	for {
		fmt.Println("Fetch repositories with handler!")

		repos, resp, err := client.github.Repositories.List(context.Background(), user, opt)
		if err != nil {
			handler([]*github.Repository{}, err, false)
			return
		}

		hasNext := (resp.NextPage != 0)
		if handler(repos, nil, hasNext) != true {
			return
		}

		if hasNext != true {
			return
		}

		opt.Page = resp.NextPage
	}
}
