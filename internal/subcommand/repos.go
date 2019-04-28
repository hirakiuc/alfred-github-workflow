package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/internal/api"
)

// ReposCommand describe a subcommand to fetch repos
type ReposCommand struct {
	Name  string
	Limit int
}

// NewReposCommand return a ReposCommand instance.
func NewReposCommand(name string) ReposCommand {
	return ReposCommand{
		Name:  name,
		Limit: 50,
	}
}

// Run start this subcommand.
func (cmd ReposCommand) Run(ctx context.Context, wf *aw.Workflow) {
	items := []*github.Repository{}

	client := api.NewClient()
	client.FetchReposByUserWithHandler(ctx, cmd.Name, func(repos []*github.Repository, err error, hasNext bool) bool {
		if err != nil {
			return false
		}

		for _, repo := range repos {
			items = append(items, repo)
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
}
