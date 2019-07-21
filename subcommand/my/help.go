package my

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
}