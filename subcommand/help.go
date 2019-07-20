package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
)

// HelpCommand describe a subcommand to show the subcommands.
type HelpCommand struct {
}

// NewHelpCommand return an instance of this subcommand.
func NewHelpCommand() HelpCommand {
	return HelpCommand{}
}

// Run start this subcommand.
func (cmd HelpCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	subcommands := []struct {
		name string
		desc string
	}{
		{
			name: "user",
			desc: "",
		},
		{
			name: "user/repo",
			desc: "",
		},
		{
			name: ">",
			desc: "",
		},
		{
			name: "my",
			desc: "",
		},
	}

	for _, cmd := range subcommands {
		wf.NewItem(cmd.name).
			Subtitle(cmd.desc).
			Autocomplete(cmd.name).
			Valid(false)
	}
}
