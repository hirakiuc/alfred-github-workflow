package repo

import (
	"context"
	"fmt"
	"os"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/api"
	"github.com/hirakiuc/alfred-github-workflow/internal/cache"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// HelpCommand describe a subcommand to show the repo subcommand.
type HelpCommand struct {
	Owner string
	Repo  string

	Limit int
}

// NewHelpCommand return an instance of this subcommand.
func NewHelpCommand(owner string, repo string) HelpCommand {
	return HelpCommand{
		Owner: owner,
		Repo:  repo,
		Limit: 50,
	}
}

func (cmd HelpCommand) command(name string) string {
	return cmd.Owner + "/" + cmd.Repo + " " + name + " "
}

func (cmd HelpCommand) htmlURL(name string) string {
	components := []string{
		"https://github.com",
		cmd.Owner,
		cmd.Repo,
		name,
	}

	return strings.Join(components, "/")
}

func (cmd HelpCommand) fetchRepos(ctx context.Context, wf *aw.Workflow) ([]model.Repo, error) {
	store := cache.NewReposCache(wf)

	repos, err := store.GetCache(cmd.Owner)
	if err != nil {
		return []model.Repo{}, err
	}

	if len(repos) != 0 {
		return repos, nil
	}

	client := api.NewClient(ctx)
	repos, err = client.FetchReposByOwner(ctx, cmd.Owner)
	if err != nil {
		return []model.Repo{}, err
	}

	return store.Store(cmd.Owner, repos)
}

func findReposContains(repos []model.Repo, key string) []model.Repo {
	keyword := strings.ToUpper(key)

	ret := []model.Repo{}
	for _, repo := range repos {
		if strings.Contains(strings.ToUpper(repo.Name), keyword) == true {
			ret = append(ret, repo)
		}
	}

	return ret
}

func (cmd HelpCommand) appendSubCommand(wf *aw.Workflow) {
	// Show subcommands if the repo found.
	subcommands := []struct {
		name string
		desc string
	}{
		{
			"branches",
			"Show branches",
		},
		{
			"issues",
			"Show issues",
		},
		{
			"milestones",
			"Show milestones",
		},
		{
			"pulls",
			"Show pull requests",
		},
	}

	for _, sub := range subcommands {
		wf.NewItem(sub.name).
			Subtitle(sub.desc).
			Autocomplete(cmd.command(sub.name)).
			Arg(cmd.htmlURL(sub.name)).
			Valid(true)
	}
}

// Run start this subcommand
func (cmd HelpCommand) Run(ctx context.Context, wf *aw.Workflow) {
	repos, err := cmd.fetchRepos(ctx, wf)
	if err != nil {
		wf.FatalError(err)
		return
	}

	founds := findReposContains(repos, cmd.Repo)
	for _, repo := range founds {
		fmt.Fprintf(os.Stderr, "found:%s\n", repo.Name)
	}

	if len(founds) == 1 {
		repo := founds[0]
		if strings.ToUpper(repo.Name) == strings.ToUpper(cmd.Repo) {
			// show subcommand if found exactly one repo.
			cmd.appendSubCommand(wf)
		} else {
			// show the repo if found but not exactly same name.
			wf.NewItem(repo.Name).
				Subtitle(repo.Description).
				Autocomplete(cmd.Owner + "/" + repo.Name + " ").
				Arg(repo.HTMLURL).
				Valid(true)
		}
	} else {
		// show repos if found some repos.
		for _, repo := range founds {
			wf.NewItem(repo.Name).
				Subtitle(repo.Description).
				Autocomplete(cmd.Owner + "/" + repo.Name + " ").
				Arg(repo.HTMLURL).
				Valid(true)
		}
	}

	wf.WarnEmpty("No such repo found.", "")
}
