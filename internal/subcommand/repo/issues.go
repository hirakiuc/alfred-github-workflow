package repo

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/api"
	"github.com/hirakiuc/alfred-github-workflow/internal/cache"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// IssueCommand describe a subcommand to fetch issues
type IssueCommand struct {
	Owner string
	Repo  string

	Query string
	Limit int
}

// NewIssueCommand return a IssueCommand instance
func NewIssueCommand(owner string, repo string, query string) IssueCommand {
	return IssueCommand{
		Owner: owner,
		Repo:  repo,
		Query: query,
		Limit: 100,
	}
}

func (cmd IssueCommand) fetchIssues(ctx context.Context, wf *aw.Workflow) ([]model.Issue, error) {
	store := cache.NewIssuesCache(wf)

	issues, err := store.GetCache(cmd.Owner, cmd.Repo)
	if err != nil {
		return []model.Issue{}, err
	}

	if len(issues) != 0 {
		return issues, nil
	}

	client, err := api.NewClient(ctx, wf)
	if err != nil {
		return []model.Issue{}, err
	}

	issues, err = client.FetchIssues(ctx, cmd.Owner, cmd.Repo)
	if err != nil {
		return []model.Issue{}, err
	}

	return store.Store(cmd.Owner, cmd.Repo, issues)
}

// Run start this subcommand.
func (cmd IssueCommand) Run(ctx context.Context, wf *aw.Workflow) {
	issues, err := cmd.fetchIssues(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	// Add items
	for _, issue := range issues {
		wf.NewItem(issue.GetItemTitle()).
			Subtitle(issue.GetItemSubtitle()).
			Arg(issue.HTMLURL).
			Valid(true)
	}

	if len(cmd.Query) > 0 {
		wf.Filter(cmd.Query)
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No issues found.", "")
}
