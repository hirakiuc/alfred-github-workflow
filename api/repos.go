package api

import (
	"context"

	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// FetchReposByOwner fetch the repos.
func (client *Client) FetchReposByOwner(ctx context.Context, owner string) ([]model.Repo, error) {
	opt := &github.RepositoryListOptions{
		Visibility: "all",
	}

	items := []model.Repo{}

	for {
		repos, resp, err := client.github.Repositories.List(ctx, owner, opt)
		if err != nil {
			return items, err
		}

		items = append(items, model.ConvertRepos(repos)...)

		if resp.NextPage == 0 {
			return items, nil
		}

		opt.Page = resp.NextPage
	}
}

// FetchReposByOrgs fetch the repository in the organization.
func (client *Client) FetchReposByOrgs(ctx context.Context, org string) ([]model.Repo, error) {
	opt := &github.RepositoryListByOrgOptions{Type: "all"}

	items := []model.Repo{}

	for {
		repos, resp, err := client.github.Repositories.ListByOrg(ctx, org, opt)
		if err != nil {
			return items, err
		}

		items = append(items, model.ConvertRepos(repos)...)

		if resp.NextPage == 0 {
			return items, nil
		}

		opt.Page = resp.NextPage
	}
}

// FetchRepo fetch the target repository
func (client *Client) FetchRepo(ctx context.Context, owner string, repo string) (*model.Repo, error) {
	r, _, err := client.github.Repositories.Get(ctx, owner, repo)
	if err != nil {
		return nil, err
	}

	v := model.ConvertRepo(r)

	return &v, nil
}
