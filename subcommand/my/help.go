package my

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
)

type HelpCommand struct {
	Limit int

	subcommand.BaseCommand
}

func NewHelpCommand(args []string) HelpCommand {
	return HelpCommand{
		Limit: 50,

		BaseCommand: subcommand.BaseCommand{
			Args: args,
		},
	}
}

// Run start this saubcommand
func (cmd HelpCommand) Run(ctx context.Context, wf *aw.Workflow) {
	subcommands := []struct {
		name string
		desc string
	}{
		{
			name: "pulls",
			desc: "",
		},
	}

	for _, cmd := range subcommands {
		wf.NewItem(cmd.name).
			Subtitle(cmd.desc).
			Autocomplete("my " + cmd.name + " ").
			Valid(false)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}
}
