package repo

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/api"
	"github.com/hirakiuc/alfred-github-workflow/cache"
	"github.com/hirakiuc/alfred-github-workflow/icon"
	"github.com/hirakiuc/alfred-github-workflow/model"
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
)

// BranchesCommand describe a subcommand to fetch branches.
type BranchesCommand struct {
	Owner string
	Repo  string
	Limit int

	subcommand.BaseCommand
}

// NewBranchesCommand return an instance of BranchesCommand
func NewBranchesCommand(owner string, repo string, args []string) BranchesCommand {
	return BranchesCommand{
		Owner: owner,
		Repo:  repo,
		Limit: 100,

		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
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

	client, err := api.NewClient(ctx, wf)
	if err != nil {
		return []model.Branch{}, err
	}

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

	icon, _ := icon.GetIcon(icon.TypeBranch)

	// Add items
	for _, branch := range branches {
		item := wf.NewItem(branch.Name).
			Arg(branch.HTMLURL).
			Icon(icon).
			Valid(true)

		if icon != nil {
			item.Icon(icon)
		}
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No branches found.", "")
}
