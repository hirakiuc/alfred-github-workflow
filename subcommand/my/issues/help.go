package issues

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

func (cmd HelpCommand) Run(ctx context.Context, wf *aw.Workflow) {
	subcommands := []struct {
		name string
		desc string
	}{
		{
			name: "created",
			desc: "",
		},
		{
			name: "assigned",
			desc: "",
		},
		{
			name: "mentioned",
			desc: "",
		},
	}

	for _, cmd := range subcommands {
		wf.NewItem(cmd.name).
			Subtitle(cmd.desc).
			Autocomplete("my issues " + cmd.name).
			Valid(true)
	}

	if cmd.HasQuery() {
		wf.Filter(cmd.Query())
	}
}
