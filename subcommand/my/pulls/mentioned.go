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

type MentionedCommand struct {
	Limit int

	subcommand.BaseCommand
}

func NewMentionedCommand(args []string) MentionedCommand {
	return MentionedCommand{
		Limit: 100,

		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
	}
}

func fetchMentioned(
	ctx context.Context, wf *aw.Workflow, client *api.Client, user string) ([]model.Issue, error) {
	store := cache.NewMentionedCache(wf)

	issues, err := store.GetCache(user)
	if err != nil {
		return []model.Issue{}, err
	}
	if len(issues) != 0 {
		return issues, nil
	}

	issues, err = client.FetchMentioned(ctx, user)
	if err != nil {
		return []model.Issue{}, err
	}
	if len(issues) == 0 {
		return []model.Issue{}, nil
	}

	return store.Store(user, issues)
}

func (cmd MentionedCommand) Run(ctx context.Context, wf *aw.Workflow) {
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

	issues, err := fetchMentioned(ctx, wf, client, user.Login)
	if err != nil {
		wf.FatalError(err)
		return
	}

	pullIcon, _ := icon.GetIcon(icon.TypePull)
	issueIcon, _ := icon.GetIcon(icon.TypeIssue)

	// Add items
	for _, issue := range issues {
		item := wf.NewItem(issue.GetItemTitle()).
			Subtitle(issue.GetItemSubtitle()).
			Arg(issue.HTMLURL).
			Valid(true)

		if issue.IsPullRequest {
			item.Icon(pullIcon)
		} else {
			item.Icon(issueIcon)
		}
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}

	// Show a warning in Alfred if there are no items
	wf.WarnEmpty("No issues/pulls found.", "")
}
