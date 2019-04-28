package api

import (
	"context"

	"github.com/google/go-github/github"
)

// FetchReposHandler describe a handler interface
type FetchReposHandler func(repos []*github.Repository, err error, hasNext bool) bool

// FetchReposByUserWithHandler fetch the repos.
func (client *Client) FetchReposByUserWithHandler(ctx context.Context, owner string, handler FetchReposHandler) {
	opt := &github.RepositoryListOptions{
		Visibility: "public",
	}

	for {
		repos, resp, err := client.github.Repositories.List(context.Background(), owner, opt)
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
