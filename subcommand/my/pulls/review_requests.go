package pulls

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/api"
	"github.com/hirakiuc/alfred-github-workflow/cache"
	"github.com/hirakiuc/alfred-github-workflow/icon"
	"github.com/hirakiuc/alfred-github-workflow/model"
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
)

type ReviewRequestsCommand struct {
	Query string
	Limit int
}

func NewReviewRequestsCommand(args []string) ReviewRequestsCommand {
	return ReviewRequestsCommand{
		Query: strings.Join(args, " "),
		Limit: 100,
	}
}

func fetchReviewRequests(
	ctx context.Context, wf *aw.Workflow, client *api.Client, user string) ([]model.Issue, error) {
	store := cache.NewReviewRequestsCache(wf)

	issues, err := store.GetCache(user)
	if err != nil {
		return []model.Issue{}, err
	}
	if len(issues) != 0 {
		return issues, nil
	}

	issues, err = client.FetchReviewRequests(ctx, user)
	if err != nil {
		return []model.Issue{}, err
	}
	if len(issues) == 0 {
		return []model.Issue{}, nil
	}

	return store.Store(user, issues)
}

func (cmd ReviewRequestsCommand) Run(ctx context.Context, wf *aw.Workflow) {
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

	issues, err := fetchReviewRequests(ctx, wf, client, user.Login)
	if err != nil {
		wf.FatalError(err)
		return
	}

	icon, _ := icon.GetIcon(icon.TypePull)

	// Add Items
	for _, issue := range issues {
		item := wf.NewItem(issue.GetItemTitle()).
			Subtitle(issue.GetItemSubtitle()).
			Arg(issue.HTMLURL).
			Valid(true)

		if icon != nil {
			item.Icon(icon)
		}
	}

	if len(cmd.Query) > 0 {
		wf.Filter(cmd.Query)
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No request found.", "")
}
