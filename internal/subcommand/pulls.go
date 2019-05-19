package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/api"
	"github.com/hirakiuc/alfred-github-workflow/internal/cache"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// PullsCommand describe a subcommand to fetch pull requests
type PullsCommand struct {
	Owner string
	Repo  string

	Query string
	Limit int
}

// NewPullsCommand return an instance of PullsCommand
func NewPullsCommand(owner string, repo string, query string) PullsCommand {
	return PullsCommand{
		Owner: owner,
		Repo:  repo,
		Query: query,
		Limit: 100,
	}
}

func (cmd PullsCommand) fetchPulls(ctx context.Context, wf *aw.Workflow) ([]model.PullRequest, error) {
	store := cache.NewPullsCache(wf)

	pulls, err := store.GetCache(cmd.Owner, cmd.Repo)
	if err != nil {
		return []model.PullRequest{}, nil
	}

	if len(pulls) != 0 {
		return pulls, nil
	}

	client := api.NewClient(ctx)
	pulls, err = client.FetchPulls(ctx, cmd.Owner, cmd.Repo)
	if err != nil {
		return []model.PullRequest{}, err
	}

	return store.Store(cmd.Owner, cmd.Repo, pulls)
}

// Run start this subcommand.
func (cmd PullsCommand) Run(ctx context.Context, wf *aw.Workflow) {
	pulls, err := cmd.fetchPulls(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	// Add items
	for _, pull := range pulls {
		wf.NewItem(pull.Title)
	}

	if len(cmd.Query) > 0 {
		wf.Filter(cmd.Query)
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No pull requests found.", "")
}
