package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/internal/api"
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

// Run start this subcommand.
func (cmd BranchesCommand) Run(ctx context.Context, wf *aw.Workflow) {
	items := []*github.Branch{}

	client := api.NewClient()
	client.FetchBranches(ctx, cmd.Owner, cmd.Repo, func(branches []*github.Branch, err error, hasNext bool) bool {
		if err != nil {
			return false
		}

		for _, branch := range branches {
			items = append(items, branch)
		}

		if len(items) >= cmd.Limit {
			return false
		}

		return true
	})

	// Add items
	for _, item := range items {
		wf.NewItem(*item.Name)
	}

	if len(cmd.Query) > 0 {
		wf.Filter(cmd.Query)
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No branches found.", "")
}
