package subcommand

import (
	"context"

	aw "github.com/deanishe/awgo"
)

// SubCommand describe a sub command.
type SubCommand interface {
	// Run the subcommand
	Run(ctx context.Context, wf *aw.Workflow)

	// Wait finish the task.
	Wait()
}
