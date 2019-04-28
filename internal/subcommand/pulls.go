package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/internal/api"
)

// PullsCommand describe a subcommand to fetch pull requests
type PullsCommand struct {
	Owner string
	Repo  string
	Limit int
}

// NewPullsCommand return an instance of PullsCommand
func NewPullsCommand(owner string, repo string) PullsCommand {
	return PullsCommand{
		Owner: owner,
		Repo:  repo,
	}
}

// Run start this subcommand.
func (cmd PullsCommand) Run(ctx context.Context, wf *aw.Workflow) {
	items := []*github.PullRequest{}

	client := api.NewClient()
	client.FetchPulls(ctx, cmd.Owner, cmd.Repo, func(pulls []*github.PullRequest, err error, hasNext bool) bool {
		if err != nil {
			return false
		}

		for _, pull := range pulls {
			items = append(items, pull)
		}

		if len(items) >= cmd.Limit {
			return false
		}

		return true
	})

	// Add items
	for _, item := range items {
		wf.NewItem(*item.Title)
	}
}
