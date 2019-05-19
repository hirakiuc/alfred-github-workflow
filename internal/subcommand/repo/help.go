package repo

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
)

// HelpCommand describe a subcommand to show the repo subcommand.
type HelpCommand struct {
	Owner string
	Repo  string
}

// NewHelpCommand return an instance of this subcommand.
func NewHelpCommand(owner string, repo string) HelpCommand {
	return HelpCommand{
		Owner: owner,
		Repo:  repo,
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

// Run start this subcommand
func (cmd HelpCommand) Run(_ctx context.Context, wf *aw.Workflow) {
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
