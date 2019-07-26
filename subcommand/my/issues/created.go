package issues

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/api"
	"github.com/hirakiuc/alfred-github-workflow/cache"
	"github.com/hirakiuc/alfred-github-workflow/icon"
	"github.com/hirakiuc/alfred-github-workflow/model"
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
)

type CreatedCommand struct {
	Limit int

	subcommand.BaseCommand
}

func NewCreatedCommand(args []string) CreatedCommand {
	return CreatedCommand{
		Limit: 50,

		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
	}
}

func fetchIssuesCreated(ctx context.Context, wf *aw.Workflow, client *api.Client, user string) ([]model.Issue, error) {
	store := cache.NewIssuesCreatedCache(wf)

	issues, err := store.GetCache(user)
	if err != nil {
		return []model.Issue{}, err
	}
	if len(issues) != 0 {
		return issues, nil
	}

	issues, err = client.FetchIssuesCreated(ctx, user)
	if err != nil {
		return []model.Issue{}, err
	}
	if len(issues) == 0 {
		return []model.Issue{}, nil
	}

	return store.Store(user, issues)
}

func (cmd CreatedCommand) Run(ctx context.Context, wf *aw.Workflow) {
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

	issues, err := fetchIssuesCreated(ctx, wf, client, user.Login)
	if err != nil {
		wf.FatalError(err)
		return
	}

	icon, _ := icon.GetIcon(icon.TypeIssue)

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

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No issues found.", "")
}
