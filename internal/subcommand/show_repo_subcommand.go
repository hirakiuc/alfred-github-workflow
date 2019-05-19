package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
)

// ShowRepoSubCommand describe a subcommand to show the repo subcommand.
type ShowRepoSubCommand struct {
}

// NewShowRepoSubCommand return an instance of this subcommand.
func NewShowRepoSubCommand() ShowRepoSubCommand {
	return ShowRepoSubCommand{}
}

// Run start this subcommand
func (cmd ShowRepoSubCommand) Run(_ctx context.Context, wf *aw.Workflow) {
	subcommands := []string{
		"branches",
		"issues",
		"milestones",
		"pulls",
		"projects",
	}

	for _, name := range subcommands {
		wf.NewItem(name)
	}
}
