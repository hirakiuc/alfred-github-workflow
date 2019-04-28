package api

import (
	"context"

	"github.com/google/go-github/github"
)

// FetchBranchesHandler describe a handler interface
type FetchBranchesHandler func(branches []*github.Branch, err error, hasNext bool) bool

// FetchBranches fetch the branches in the repository.
func (client *Client) FetchBranches(ctx context.Context, owner string, repo string, handler FetchBranchesHandler) {
	opt := github.ListOptions{}

	for {
		branches, resp, err := client.github.Repositories.ListBranches(ctx, owner, repo, &opt)
		if err != nil {
			handler([]*github.Branch{}, err, false)
			return
		}

		hasNext := (resp.NextPage != 0)
		if handler(branches, nil, hasNext) != true {
			return
		}

		if hasNext != true {
			return
		}

		opt.Page = resp.NextPage
	}
}
