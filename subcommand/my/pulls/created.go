package pulls

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/api"
	"github.com/hirakiuc/alfred-github-workflow/cache"
	"github.com/hirakiuc/alfred-github-workflow/icon"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

type CreatedCommand struct {
	Query string
	Limit int
}

func NewCreatedCommand(args []string) CreatedCommand {
	return CreatedCommand{
		Query: strings.Join(args, " "),
		Limit: 100,
	}
}

func fetchPullsCreated(ctx context.Context, wf *aw.Workflow, client *api.Client, user string) ([]model.Issue, error) {
	store := cache.NewPullsCreatedCache(wf)

	issues, err := store.GetCache(user)
	if err != nil {
		return []model.Issue{}, err
	}
	if len(issues) != 0 {
		return issues, nil
	}

	issues, err = client.FetchPullsCreated(ctx, user)
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

	user, err := fetchUser(ctx, wf, client)
	if err != nil {
		wf.FatalError(err)
		return
	}

	issues, err := fetchPullsCreated(ctx, wf, client, user.Login)
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

	if len(cmd.Query) > 0 {
		wf.Filter(cmd.Query)
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No pulls found.", "")
}
