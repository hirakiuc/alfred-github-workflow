package config

import (
	"context"

	aw "github.com/deanishe/awgo"
)

// HelpCommand describe a subcommand to show the config subcommand.
type HelpCommand struct {
}

// NewHelpCommand return an instance of this subcommand.
func NewHelpCommand() HelpCommand {
	return HelpCommand{}
}

func (cmd HelpCommand) Wait() {}

// Run start this subcommand.
func (cmd HelpCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	subcommands := []string{
		"token",
	}

	for _, name := range subcommands {
		wf.NewItem(name)
	}
}
