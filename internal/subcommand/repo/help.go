package repo

import (
	"context"

	aw "github.com/deanishe/awgo"
)

// HelpCommand describe a subcommand to show the repo subcommand.
type HelpCommand struct {
}

// NewHelpCommand return an instance of this subcommand.
func NewHelpCommand() HelpCommand {
	return HelpCommand{}
}

// Run start this subcommand
func (cmd HelpCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	subcommands := []string{
		"branches",
		"issues",
		"milestones",
		"pulls",
	}

	for _, name := range subcommands {
		wf.NewItem(name)
	}
}
