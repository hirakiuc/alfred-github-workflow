package repo

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/api"
	"github.com/hirakiuc/alfred-github-workflow/internal/cache"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// BranchesCommand describe a subcommand to fetch branches.
type BranchesCommand struct {
	Owner string
	Repo  string
	Query string
	Limit int
}

// NewBranchesCommand return an instance of BranchesCommand
func NewBranchesCommand(owner string, repo string, query string) BranchesCommand {
	return BranchesCommand{
		Owner: owner,
		Repo:  repo,

		Query: query,
		Limit: 100,
	}
}

func (cmd BranchesCommand) fetchBranches(ctx context.Context, wf *aw.Workflow) ([]model.Branch, error) {
	store := cache.NewBranchesCache(wf)

	branches, err := store.GetCache(cmd.Owner, cmd.Repo)
	if err != nil {
		return []model.Branch{}, err
	}

	if len(branches) != 0 {
		return branches, nil
	}

	client := api.NewClient()
	branches, err = client.FetchBranches(ctx, cmd.Owner, cmd.Repo)
	if err != nil {
		return []model.Branch{}, err
	}

	return store.Store(cmd.Owner, cmd.Repo, branches)
}

// Run start this subcommand.
func (cmd BranchesCommand) Run(ctx context.Context, wf *aw.Workflow) {
	branches, err := cmd.fetchBranches(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	// Add items
	for _, branch := range branches {
		wf.NewItem(branch.Name)
	}

	if len(cmd.Query) > 0 {
		wf.Filter(cmd.Query)
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No branches found.", "")
}