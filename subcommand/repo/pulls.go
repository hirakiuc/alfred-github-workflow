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

// PullsCommand describe a subcommand to fetch pull requests
type PullsCommand struct {
	Owner string
	Repo  string

	Query string
	Limit int
}

// NewPullsCommand return an instance of PullsCommand
func NewPullsCommand(owner string, repo string, args []string) PullsCommand {
	return PullsCommand{
		Owner: owner,
		Repo:  repo,
		Query: strings.Join(args, " "),
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

	client, err := api.NewClient(ctx, wf)
	if err != nil {
		return []model.PullRequest{}, err
	}

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

	icon, _ := icon.GetIcon(icon.TypePull)

	// Add items
	for _, pull := range pulls {
		item := wf.NewItem(pull.GetItemTitle()).
			Subtitle(pull.GetItemSubtitle()).
			Arg(pull.HTMLURL).
			Valid(true)

		if icon != nil {
			item.Icon(icon)
		}
	}

	if len(cmd.Query) > 0 {
		wf.Filter(cmd.Query)
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No pull requests found.", "")
}
