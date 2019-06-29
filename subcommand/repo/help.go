package repo

import (
	"context"
	"fmt"
	"os"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/api"
	"github.com/hirakiuc/alfred-github-workflow/cache"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// HelpCommand describe a subcommand to show the repo subcommand.
type HelpCommand struct {
	Owner string
	Repo  string

	Query string
	Limit int
}

// NewHelpCommand return an instance of this subcommand.
func NewHelpCommand(owner string, repo string, args []string) HelpCommand {
	return HelpCommand{
		Owner: owner,
		Repo:  repo,
		Query: strings.Join(args, " "),
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

func (cmd HelpCommand) fetchRepos(_ context.Context, wf *aw.Workflow) ([]model.Repo, error) {
	store := cache.NewReposCache(wf)

	repos, err := store.GetCache(cmd.Owner)
	if err != nil {
		return []model.Repo{}, err
	}

	if len(repos) != 0 {
		return repos, nil
	}

	return repos, nil
}

func (cmd HelpCommand) fetchRepo(ctx context.Context, wf *aw.Workflow) (*model.Repo, error) {
	client, err := api.NewClient(ctx, wf)
	if err != nil {
		return nil, err
	}

	repo, err := client.FetchRepo(ctx, cmd.Owner, cmd.Repo)
	if err != nil {
		return nil, err
	}

	return repo, nil
}

func findReposContains(repos []model.Repo, key string) []model.Repo {
	keyword := strings.ToUpper(key)

	ret := []model.Repo{}
	for _, repo := range repos {
		if strings.Contains(strings.ToUpper(repo.Name), keyword) {
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

	if len(cmd.Query) > 0 {
		wf.Filter(cmd.Query)
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
		if len(repos) > 0 {
			// show repos if found some repos.
			for _, repo := range founds {
				wf.NewItem(repo.Name).
					Subtitle(repo.Description).
					Autocomplete(cmd.Owner + "/" + repo.Name + " ").
					Arg(repo.HTMLURL).
					Valid(true)
			}

			if len(cmd.Query) > 0 {
				wf.Filter(cmd.Query)
			}
		} else {
			repo, err := cmd.fetchRepo(ctx, wf)
			if err != nil {
				// show error
				wf.FatalError(err)
				return
			}

			if repo != nil {
				// show subcommands
				cmd.appendSubCommand(wf)
			}
		}
	}

	wf.WarnEmpty("No such repo found.", "")
}
