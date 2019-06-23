package subcommand

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/api"
	"github.com/hirakiuc/alfred-github-workflow/internal/cache"
	"github.com/hirakiuc/alfred-github-workflow/internal/icon"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// ReposCommand describe a subcommand to fetch repos
type ReposCommand struct {
	Owner string

	Query string
	Limit int
}

// NewReposCommand return a ReposCommand instance.
func NewReposCommand(owner string, args []string) ReposCommand {
	return ReposCommand{
		Owner: owner,
		Query: strings.Join(args, " "),
		Limit: 50,
	}
}

func (cmd ReposCommand) fetchRepos(ctx context.Context, wf *aw.Workflow) ([]model.Repo, error) {
	store := cache.NewReposCache(wf)

	repos, err := store.GetCache(cmd.Owner)
	if err != nil {
		return []model.Repo{}, err
	}

	if len(repos) != 0 {
		return repos, nil
	}

	client, err := api.NewClient(ctx, wf)
	if err != nil {
		return []model.Repo{}, err
	}

	repos, err = client.FetchReposByOwner(ctx, cmd.Owner)
	if err != nil {
		return []model.Repo{}, err
	}

	if len(repos) > 1 {
		return store.Store(cmd.Owner, repos)
	}

	repos, err = client.FetchReposByOrgs(ctx, cmd.Owner)
	if err != nil {
		return []model.Repo{}, err
	}

	return store.Store(cmd.Owner, repos)
}

// Run start this subcommand.
func (cmd ReposCommand) Run(ctx context.Context, wf *aw.Workflow) {
	repos, err := cmd.fetchRepos(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	icon, _ := icon.GetIcon(icon.TypeRepo)

	// Add items
	for _, repo := range repos {
		item := wf.NewItem(repo.Name).
			Subtitle(repo.Description).
			Autocomplete(cmd.Owner + "/" + repo.Name + " ").
			Arg(repo.HTMLURL).
			Valid(true)

		if icon != nil {
			item.Icon(icon)
		}
	}

	if len(cmd.Query) > 0 {
		wf.Filter(cmd.Query)
	}

	// Show a warning in Alfred if there are no repos
	wf.WarnEmpty("No repos found.", "")
}
