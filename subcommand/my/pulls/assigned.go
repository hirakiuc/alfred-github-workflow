package pulls

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/api"
	"github.com/hirakiuc/alfred-github-workflow/cache"
	"github.com/hirakiuc/alfred-github-workflow/icon"
	"github.com/hirakiuc/alfred-github-workflow/model"
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
)

type AssignedCommand struct {
	Limit int

	subcommand.BaseCommand
}

func NewAssignedCommand(args []string) AssignedCommand {
	return AssignedCommand{
		Limit: 100,

		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
	}
}

func fetchPullsAssigned(ctx context.Context, wf *aw.Workflow, client *api.Client, user string) ([]model.Issue, error) {
	store := cache.NewPullsAssignedCache(wf)

	issues, err := store.GetCache(user)
	if err != nil {
		return []model.Issue{}, err
	}
	if len(issues) != 0 {
		return issues, nil
	}

	issues, err = client.FetchPullsAssigned(ctx, user)
	if err != nil {
		return []model.Issue{}, err
	}
	if len(issues) == 0 {
		return []model.Issue{}, nil
	}

	return store.Store(user, issues)
}

func (cmd AssignedCommand) Run(ctx context.Context, wf *aw.Workflow) {
	client, err := api.NewClient(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	user, err := subcommand.FetchAuthorizedUser(ctx, wf, client)
	if err != nil {
		wf.FatalError(err)
		return
	}

	issues, err := fetchPullsAssigned(ctx, wf, client, user.Login)
	if err != nil {
		wf.FatalError(err)
		return
	}

	icon, _ := icon.GetIcon(icon.TypePull)

	// Add Items
	for _, issue := range issues {
		wf.NewItem(issue.GetItemTitle()).
			Subtitle(issue.GetItemSubtitle()).
			Arg(issue.HTMLURL).
			Icon(icon).
			Valid(true)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}

	// Show a warning in Alfred if there are no items.
	wf.WarnEmpty("No pulls found.", "")
}
