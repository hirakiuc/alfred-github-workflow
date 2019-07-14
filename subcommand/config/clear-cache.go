package config

import (
	"context"

	aw "github.com/deanishe/awgo"
)

type ClearCacheCommand struct {
}

func NewClearCacheCommand() ClearCacheCommand {
	return ClearCacheCommand{}
}

func (cmd ClearCacheCommand) Run(ctx context.Context, wf *aw.Workflow) {
	err := wf.ClearCache()
	if err != nil {
		wf.FatalError(err)
		return
	}

	wf.NewItem("Success").
		Subtitle("All cache cleared successfully.").
		Valid(true)
}
