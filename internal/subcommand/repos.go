package subcommand

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/api"
	"github.com/hirakiuc/alfred-github-workflow/internal/cache"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// ReposCommand describe a subcommand to fetch repos
type ReposCommand struct {
	Owner string

	Query string
	Limit int
}

// NewReposCommand return a ReposCommand instance.
func NewReposCommand(owner string, args []string) ReposCommand {
	return ReposCommand{
		Owner: owner,
		Query: strings.Join(args, " "),
		Limit: 50,
	}
}

func (cmd ReposCommand) fetchRepos(ctx context.Context, wf *aw.Workflow) ([]model.Repo, error) {
	store := cache.NewReposCache(wf)

	repos, err := store.GetCache(cmd.Owner)
	if err != nil {
		return []model.Repo{}, err
	}

	if len(repos) != 0 {
		return repos, nil
	}

	client := api.NewClient()
	repos, err = client.FetchReposByOwner(ctx, cmd.Owner)
	if err != nil {
		return []model.Repo{}, err
	}

	return store.Store(cmd.Owner, repos)
}

// Run start this subcommand.
func (cmd ReposCommand) Run(ctx context.Context, wf *aw.Workflow) {
	repos, err := cmd.fetchRepos(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	// Add items
	for _, repo := range repos {
		wf.NewItem(repo.Name)
	}

	if len(cmd.Query) > 0 {
		wf.Filter(cmd.Query)
	}
}
