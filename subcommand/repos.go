package subcommand

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/api"
	"github.com/hirakiuc/alfred-github-workflow/cache"
	"github.com/hirakiuc/alfred-github-workflow/icon"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// ReposCommand describe a subcommand to fetch repos
type ReposCommand struct {
	Owner string

	Query string
	Limit int

	runner *api.AsyncCall
	ch     chan []model.Repo
}

// NewReposCommand return a ReposCommand instance.
func NewReposCommand(owner string, args []string) ReposCommand {
	return ReposCommand{
		Owner: owner,
		Query: strings.Join(args, " "),
		Limit: 50,

		runner: api.NewAsyncCall(),
		ch:     make(chan []model.Repo, 1),
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

	ch := cmd.ch
	cb := func() error {
		client, err := api.NewClient(ctx, wf)
		if err != nil {
			ch <- []model.Repo{}
			return err
		}

		repos, err = client.FetchReposByOwner(ctx, cmd.Owner)
		if err != nil {
			ch <- []model.Repo{}
			return err
		}

		if len(repos) > 1 {
			repos, err := store.Store(cmd.Owner, repos)
			ch <- repos
			return err
		}

		repos, err = client.FetchReposByOrgs(ctx, cmd.Owner)
		if err != nil {
			ch <- repos
			return err
		}

		repos, err := store.Store(cmd.Owner, repos)
		ch <- repos
		return err
	}

	err = cmd.runner.RunWithTimeout(cb)
	if err != nil {
		return []model.Repo{}, err
	}

	return <-ch, nil
}

func (cmd ReposCommand) Wait() {
	cmd.runner.Wait()
	close(cmd.ch)
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
