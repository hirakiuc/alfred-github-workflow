package repo

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/api"
	"github.com/hirakiuc/alfred-github-workflow/cache"
	"github.com/hirakiuc/alfred-github-workflow/icon"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// MilestonesCommand describe a subcommand to fetch milestones.
type MilestonesCommand struct {
	Owner string
	Repo  string

	Query string
	Limit int
}

// NewMilestonesCommand return a MilestonesCommand instance
func NewMilestonesCommand(owner string, repo string, args []string) MilestonesCommand {
	return MilestonesCommand{
		Owner: owner,
		Repo:  repo,
		Query: strings.Join(args, " "),
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

	client, err := api.NewClient(ctx, wf)
	if err != nil {
		return []model.Milestone{}, err
	}

	milestones, err = client.FetchMilestones(ctx, cmd.Owner, cmd.Repo)
	if err != nil {
		return []model.Milestone{}, err
	}

	return store.Store(cmd.Owner, cmd.Repo, milestones)
}

func (cmd MilestonesCommand) Wait() {}

// Run start this subcommand
func (cmd MilestonesCommand) Run(ctx context.Context, wf *aw.Workflow) {
	milestones, err := cmd.fetchMilestones(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	icon, _ := icon.GetIcon(icon.TypeMilestone)

	// Add items
	for _, milestone := range milestones {
		item := wf.NewItem(milestone.GetItemTitle()).
			Subtitle(milestone.GetItemSubtitle()).
			Arg(milestone.HTMLURL).
			Valid(true)

		if icon != nil {
			item.Icon(icon)
		}
	}

	if len(cmd.Query) > 0 {
		wf.Filter(cmd.Query)
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No milestones found.", "")
}
