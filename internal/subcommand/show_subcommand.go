package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
)

// ShowSubCommand describe a subcommand to show the subcommands.
type ShowSubCommand struct {
}

// NewShowSubCommand return an instance of this subcommand.
func NewShowSubCommand() ShowSubCommand {
	return ShowSubCommand{}
}

// Run start this subcommand.
func (cmd ShowSubCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	subcommands := []string{
		"user",
		"user/repo",
	}

	for _, name := range subcommands {
		wf.NewItem(name)
	}
}
