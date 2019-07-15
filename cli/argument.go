package cli

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
	configcmd "github.com/hirakiuc/alfred-github-workflow/subcommand/config"
	mycmd "github.com/hirakiuc/alfred-github-workflow/subcommand/my"
	mypullscmd "github.com/hirakiuc/alfred-github-workflow/subcommand/my/pulls"
	repocmd "github.com/hirakiuc/alfred-github-workflow/subcommand/repo"
)

// Slug...
type Slug struct {
	Owner string
	Repo  string
}

const (
	repoSeparator = "/"

	cmdTypeEmpty   = ""
	cmdTypeConfig  = ">"
	cmdTypeOwner   = "owner"
	cmdTypeRepo    = "repo"
	cmdTypeMy      = "my"
	cmdTypeHelp    = "help"
	cmdTypeInvalid = "invalid"
)

// Command...
type Command struct {
	Args   []string
	Slug   *Slug
	Subcmd subcommand.SubCommand
}

func parseConfigCommandArgs(args []string) (string, []string) {
	// args: {">", ...}
	switch len(args) {
	case 0, 1:
		return cmdTypeHelp, []string{}
	case 2:
		return args[1], []string{}
	default:
		return args[1], args[2:]
	}
}

func (c *Command) createConfigSubCommand() {
	c.Slug = nil

	cmd, opts := parseConfigCommandArgs(c.Args)
	switch cmd {
	case "token":
		token := ""
		if len(opts) > 0 {
			token = opts[0]
		}

		c.Subcmd = configcmd.NewTokenCommand(token)
	case "clear-cache":
		c.Subcmd = configcmd.NewClearCacheCommand()
	default:
		c.Subcmd = configcmd.NewHelpCommand()
	}
}

func (c *Command) createOwnerSubCommand() {
	parts := strings.Split(c.Args[0], repoSeparator)
	c.Slug = &Slug{
		Owner: parts[0],
	}

	c.Subcmd = subcommand.NewReposCommand(c.Slug.Owner, c.Args[1:])
}

func parseSubCommandArgs(args []string) (string, []string) {
	switch len(args) {
	case 0:
		return cmdTypeHelp, []string{}
	case 1:
		return args[0], []string{}
	default:
		return args[0], args[1:]
	}
}

func (c *Command) createRepoSubCommand() {
	parts := strings.Split(c.Args[0], repoSeparator)
	c.Slug = &Slug{
		Owner: parts[0],
		Repo:  parts[1],
	}

	cmd, options := parseSubCommandArgs(c.Args[1:])

	switch cmd {
	case "issues":
		c.Subcmd = repocmd.NewIssueCommand(c.Slug.Owner, c.Slug.Repo, options)
	case "pulls":
		c.Subcmd = repocmd.NewPullsCommand(c.Slug.Owner, c.Slug.Repo, options)
	case "branches":
		c.Subcmd = repocmd.NewBranchesCommand(c.Slug.Owner, c.Slug.Repo, options)
	case "milestones":
		c.Subcmd = repocmd.NewMilestonesCommand(c.Slug.Owner, c.Slug.Repo, options)
	case "projects":
		c.Subcmd = repocmd.NewProjectsCommand(c.Slug.Owner, c.Slug.Repo, options)
	default:
		// Show the subcommands
		c.Subcmd = repocmd.NewHelpCommand(c.Slug.Owner, c.Slug.Repo, c.Args[1:])
	}
}

func (c *Command) createMySubCommand() {
	cmd, opts := parseMyCommandArgs(c.Args)

	switch cmd {
	case "pulls":
		c.createMyPullsCommand()
	default:
		c.Subcmd = mycmd.NewHelpCommand(opts)
	}
}

func parseMyCommandArgs(args []string) (string, []string) {
	// args: {"my", ...}
	switch len(args) {
	case 0, 1:
		return cmdTypeHelp, []string{}
	case 2:
		return args[1], []string{}
	default:
		return args[1], args[2:]
	}
}

func (c *Command) createMyPullsCommand() {
	cmd, opts := parseMyPullsCommandArgs(c.Args)

	switch cmd {
	case "review-requests":
		c.Subcmd = mypullscmd.NewReviewRequestsCommand(opts)
	default:
		c.Subcmd = mypullscmd.NewHelpCommand(opts)
	}
}

func parseMyPullsCommandArgs(args []string) (string, []string) {
	// args: {"my", "pulls", cmd, ...}
	switch len(args) {
	case 0, 1, 2:
		return cmdTypeHelp, []string{}
	case 3:
		return args[2], []string{}
	default:
		return args[2], args[3:]
	}
}

func (c *Command) createHelpCommand() {
	c.Subcmd = subcommand.NewHelpCommand()
}

// Run...
func (c *Command) Run(ctx context.Context, wf *aw.Workflow) {
	c.Subcmd.Run(ctx, wf)
}

func normalizeArgs(args []string) []string {
	ret := []string{}

	for _, component := range args {
		parts := strings.Fields(component)

		for _, arg := range parts {
			v := strings.TrimSpace(arg)
			if v != "" {
				ret = append(ret, v)
			}
		}
	}

	return ret
}

func judgeType(args []string) string {
	if len(args) == 0 {
		return cmdTypeHelp
	}

	word := args[0]
	switch word {
	case "":
		return cmdTypeEmpty
	case ">":
		return cmdTypeConfig
	case "my":
		return cmdTypeMy
	default:
		// continue
	}

	parts := strings.Split(word, repoSeparator)
	switch len(parts) {
	case 1:
		return cmdTypeOwner
	case 2:
		if len(parts[0]) > 0 {
			if len(parts[1]) > 0 {
				return cmdTypeRepo
			}

			return cmdTypeOwner
		}

		return cmdTypeInvalid
	default:
		return cmdTypeInvalid
	}
}

// ParseArgs parses the arguments
func ParseArgs(arguments []string) *Command {
	cmd := &Command{
		Args: normalizeArgs(arguments),
	}

	cmdType := judgeType(cmd.Args)
	switch cmdType {
	case cmdTypeConfig:
		cmd.createConfigSubCommand()
	case cmdTypeOwner:
		cmd.createOwnerSubCommand()
	case cmdTypeRepo:
		cmd.createRepoSubCommand()
	case cmdTypeMy:
		cmd.createMySubCommand()
	default:
		cmd.createHelpCommand()
	}

	return cmd
}
