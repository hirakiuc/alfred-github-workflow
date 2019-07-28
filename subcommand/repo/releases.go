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

type ReleasesCommand struct {
	Owner string
	Repo  string
	Limit int

	subcommand.BaseCommand
}

func NewReleasesCommand(owner string, repo string, args []string) ReleasesCommand {
	return ReleasesCommand{
		Owner: owner,
		Repo:  repo,
		Limit: 100,

		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
	}
}

func (cmd ReleasesCommand) fetchReleases(ctx context.Context, wf *aw.Workflow) ([]model.Release, error) {
	store := cache.NewReleasesCache(wf)

	rels, err := store.GetCache(cmd.Owner, cmd.Repo)
	if err != nil {
		return []model.Release{}, err
	}

	if len(rels) != 0 {
		return rels, nil
	}

	client, err := api.NewClient(ctx, wf)
	if err != nil {
		return []model.Release{}, err
	}

	rels, err = client.FetchReleases(ctx, cmd.Owner, cmd.Repo)
	if err != nil {
		return []model.Release{}, err
	}

	return store.Store(cmd.Owner, cmd.Repo, rels)
}

func (cmd ReleasesCommand) Run(ctx context.Context, wf *aw.Workflow) {
	rels, err := cmd.fetchReleases(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	icon, _ := icon.GetIcon(icon.TypeRelease)

	// Add items
	for _, rel := range rels {
		wf.NewItem(rel.Name).
			Subtitle(rel.Subtitle()).
			Arg(rel.HTMLURL).
			Icon(icon).
			Valid(true)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}

	// Show a warning in Alfred if there are no items.
	wf.WarnEmpty("No releases found.", "")
}
