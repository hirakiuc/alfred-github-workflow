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
	subcmd := getCommand(wf.Args())
	ctx := context.Background()

	subcmd.Run(ctx, wf)

	// Add a "Script Filter" result
	// wf.NewItem("First result!")

	// Send results to Alfred
	wf.SendFeedback()
}

func getCommand(args []string) subcommand.SubCommand {
	if len(args) == 0 {
		return subcommand.NewShowSubCommand()
	}

	slug := args[0]
	components := strings.Split(slug, "/")
	switch len(components) {
	case 1:
		// Fetch the repos which created by the username.
		return subcommand.NewReposCommand(slug, args[1:])
	case 2:
		owner := components[0]
		repo := components[1]
		return getRepoSubCommand(owner, repo, args[1:])
	default:
		return subcommand.NewShowSubCommand()
	}
}

func parseSubCommandArgs(args []string) (string, []string) {
	switch len(args) {
	case 0:
		return "", []string{}
	case 1:
		return args[0], []string{}
	default:
		return args[0], args[1:]
	}
}

func getRepoSubCommand(owner string, repo string, args []string) subcommand.SubCommand {
	cmd, options := parseSubCommandArgs(args)
	query := strings.Join(options, " ")

	switch cmd {
	case "issues":
		return subcommand.NewIssueCommand(owner, repo, query)
	case "pulls":
		return subcommand.NewPullsCommand(owner, repo, query)
	case "branches":
		return subcommand.NewBranchesCommand(owner, repo, query)
	case "milestones":
		return subcommand.NewMilestonesCommand(owner, repo, query)
	default:
		// Show the subcommands
		return subcommand.NewShowRepoSubCommand()
	}
}

func main() {
	// Wrap your entry point with Run() to catch and log panics and
	// show an error in Alfred instead of silently dying
	wf.Run(run)
}
