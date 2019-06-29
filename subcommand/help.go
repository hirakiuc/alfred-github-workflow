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
	subcommands := []string{
		"user",
		"user/repo",
		">",
	}

	for _, name := range subcommands {
		wf.NewItem(name)
	}
}
