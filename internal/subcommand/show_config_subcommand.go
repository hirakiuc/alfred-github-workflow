package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
)

// ShowConfigSubCommand describe a subcommand to show the config subcommand.
type ShowConfigSubCommand struct {
}

// NewShowConfigSubCommand return an instance of this subcommand.
func NewShowConfigSubCommand() ShowConfigSubCommand {
	return ShowConfigSubCommand{}
}

// Run start this subcommand.
func (cmd ShowConfigSubCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	subcommands := []string{
		"token",
	}

	for _, name := range subcommands {
		wf.NewItem(name)
	}
}
