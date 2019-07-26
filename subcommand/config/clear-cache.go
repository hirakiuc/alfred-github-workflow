package config

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
)

type ClearCacheCommand struct {
	subcommand.BaseCommand
}

func NewClearCacheCommand() ClearCacheCommand {
	return ClearCacheCommand{
		subcommand.BaseCommand{
			Args: []string{},
		},
	}
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
