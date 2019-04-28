package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/internal/api"
)

// IssueCommand describe a subcommand to fetch issues
type IssueCommand struct {
	Owner string
	Repo  string
	Limit int
}

// NewIssueCommand return a IssueCommand instance
func NewIssueCommand(owner string, repo string) IssueCommand {
	return IssueCommand{
		Owner: owner,
		Repo:  repo,
	}
}

// Run start this subcommand.
func (cmd IssueCommand) Run(ctx context.Context, wf *aw.Workflow) {
	items := []*github.Issue{}

	client := api.NewClient()
	client.FetchIssues(ctx, cmd.Owner, cmd.Repo, func(issues []*github.Issue, err error, hasNext bool) bool {
		if err != nil {
			return false
		}

		for _, issue := range issues {
			items = append(items, issue)
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
