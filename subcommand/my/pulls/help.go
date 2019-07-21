package pulls

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
		{
			name: "review-requests",
			desc: "",
		},
	}

	for _, cmd := range subcommands {
		wf.NewItem(cmd.name).
			Subtitle(cmd.desc).
			Autocomplete("my pulls " + cmd.name).
			Valid(false)
	}
}
