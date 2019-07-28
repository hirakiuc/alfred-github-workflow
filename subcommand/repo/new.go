package repo

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
)

type NewCommand struct {
	Owner string
	Repo  string

	subcommand.BaseCommand
}

// NewNewCommand return an instance of this subcommand.
func NewNewCommand(owner string, repo string, args []string) NewCommand {
	return NewCommand{
		Owner: owner,
		Repo:  repo,

		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
	}
}

func (cmd NewCommand) pullURL() string {
	return strings.Join([]string{
		"https://github.com",
		cmd.Owner,
		cmd.Repo,
		"pulls",
		"new",
	}, "/")
}

func (cmd NewCommand) issuesURL() string {
	return strings.Join([]string{
		"https://github.com",
		cmd.Owner,
		cmd.Repo,
		"issues",
		"new",
	}, "/")
}

func (cmd NewCommand) command(name string) string {
	return cmd.Owner + "/" + cmd.Repo + " new " + name + " "
}

func (cmd NewCommand) Run(ctx context.Context, wf *aw.Workflow) {
	subcommands := []struct {
		name string
		desc string
		url  string
	}{
		{
			name: "pull",
			desc: "create new pull requests",
			url:  cmd.pullURL(),
		},
		{
			name: "issues",
			desc: "create new issue",
			url:  cmd.issuesURL(),
		},
	}

	for _, sub := range subcommands {
		wf.NewItem(sub.name).
			Subtitle(sub.desc).
			Autocomplete(cmd.command(sub.name)).
			Arg(sub.url).
			Valid(true)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}
}
