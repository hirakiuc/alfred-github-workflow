package main

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/subcommand"
)

// Workflow is the main API
var wf *aw.Workflow

func init() {
	// Create a new Workflow using default settings.
	// Critical settings are provided by Alfred via environment variables.
	// so this *will* die in flames if not run in an Alfred-like environment.
	wf = aw.New()
}

// Your workflow starts here
func run() {
	subcmd := getSubCommand(wf.Args())
	ctx := context.Background()

	subcmd.Run(ctx, wf)

	// Add a "Script Filter" result
	// wf.NewItem("First result!")

	// Send results to Alfred
	wf.SendFeedback()
}

func getSubCommand(args []string) subcommand.SubCommand {
	if len(args) == 0 {
		return subcommand.NewShowSubCommand()
	}

	slug := args[0]
	components := strings.Split(slug, "/")
	switch len(components) {
	case 1:
		// Fetch the repos which created by the username.
		return subcommand.NewReposCommand(slug)
	// case 2:
	// Show the subcommands
	// case 3:
	// Invoke the subcommands

	default:
		return subcommand.NewShowSubCommand()
	}
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
