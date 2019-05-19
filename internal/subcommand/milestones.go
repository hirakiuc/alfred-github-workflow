package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/api"
	"github.com/hirakiuc/alfred-github-workflow/internal/cache"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// MilestonesCommand describe a subcommand to fetch milestones.
type MilestonesCommand struct {
	Owner string
	Repo  string

	Query string
	Limit int
}

// NewMilestonesCommand return a MilestonesCommand instance
func NewMilestonesCommand(owner string, repo string, query string) MilestonesCommand {
	return MilestonesCommand{
		Owner: owner,
		Repo:  repo,
		Query: query,
		Limit: 100,
	}
}

func (cmd MilestonesCommand) fetchMilestones(ctx context.Context, wf *aw.Workflow) ([]model.Milestone, error) {
	store := cache.NewMilestonesCache(wf)

	milestones, err := store.GetCache(cmd.Owner, cmd.Repo)
	if err != nil {
		return []model.Milestone{}, err
	}

	if len(milestones) != 0 {
		return milestones, nil
	}

	client := api.NewClient(ctx)
	milestones, err = client.FetchMilestones(ctx, cmd.Owner, cmd.Repo)
	if err != nil {
		return []model.Milestone{}, err
	}

	return store.Store(cmd.Owner, cmd.Repo, milestones)
}

// Run start this subcommand
func (cmd MilestonesCommand) Run(ctx context.Context, wf *aw.Workflow) {
	milestones, err := cmd.fetchMilestones(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	// Add items
	for _, milestone := range milestones {
		wf.NewItem(milestone.Title)
	}

	if len(cmd.Query) > 0 {
		wf.Filter(cmd.Query)
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No milestones found.", "")
}
