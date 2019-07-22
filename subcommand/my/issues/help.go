package issues

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
)

type HelpCommand struct {
	Query string
	Limit int
}

func NewHelpCommand(args []string) HelpCommand {
	return HelpCommand{
		Query: strings.Join(args, " "),
		Limit: 50,
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
}
