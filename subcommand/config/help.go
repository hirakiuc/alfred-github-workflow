package config

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
)

// HelpCommand describe a subcommand to show the config subcommand.
type HelpCommand struct {
	subcommand.BaseCommand
}

// NewHelpCommand return an instance of this subcommand.
func NewHelpCommand() HelpCommand {
	return HelpCommand{
		subcommand.BaseCommand{
			Args: []string{},
		},
	}
}

// Run start this subcommand.
func (cmd HelpCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	subcommands := []struct {
		name string
		desc string
	}{
		{
			name: "token",
			desc: "Configure github token.",
		},
		{
			name: "clear-cache",
			desc: "Clear caches.",
		},
	}

	for _, cmd := range subcommands {
		wf.NewItem(cmd.name).
			Subtitle(cmd.desc).
			Autocomplete("> " + cmd.name).
			Valid(true)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}
}
