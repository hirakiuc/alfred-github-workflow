package cli

import (
	"context"
	"strings"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/subcommand"
	configcmd "github.com/hirakiuc/alfred-github-workflow/internal/subcommand/config"
	repocmd "github.com/hirakiuc/alfred-github-workflow/internal/subcommand/repo"
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
	cmdTypeInvalid = "invalid"
)

// Command...
type Command struct {
	Args   []string
	Slug   *Slug
	Subcmd subcommand.SubCommand
}

func parseConfigCommandArgs(args []string) (string, []string) {
	switch len(args) {
	case 0:
		return "help", []string{}
	case 1:
		return args[0], []string{}
	default:
		return args[0], args[1:]
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
	default:
		c.Subcmd = configcmd.NewHelpCommand()
	}
}

func (c *Command) createOwnerSubCommand() {
	c.Slug = &Slug{
		Owner: c.Args[0],
	}

	// Remove first argument.
	c.Args = c.Args[1:]

	c.Subcmd = subcommand.NewReposCommand(c.Slug.Owner, c.Args)
}

func parseSubCommandArgs(args []string) (string, []string) {
	switch len(args) {
	case 0:
		return "help", []string{}
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

	// Remove first argument.
	c.Args = c.Args[1:]

	cmd, options := parseSubCommandArgs(c.Args)
	query := strings.Join(options, " ")

	switch cmd {
	case "issues":
		c.Subcmd = repocmd.NewIssueCommand(c.Slug.Owner, c.Slug.Repo, query)
	case "pulls":
		c.Subcmd = repocmd.NewPullsCommand(c.Slug.Owner, c.Slug.Repo, query)
	case "branches":
		c.Subcmd = repocmd.NewBranchesCommand(c.Slug.Owner, c.Slug.Repo, query)
	case "milestones":
		c.Subcmd = repocmd.NewMilestonesCommand(c.Slug.Owner, c.Slug.Repo, query)
	case "projects":
		c.Subcmd = repocmd.NewProjectsCommand(c.Slug.Owner, c.Slug.Repo, query)
	default:
		// Show the subcommands
		c.Subcmd = repocmd.NewHelpCommand(c.Slug.Owner, c.Slug.Repo)
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

func judgeType(word string) string {
	if word == "" {
		return cmdTypeEmpty
	}

	if word == ">" {
		return cmdTypeConfig
	}

	parts := strings.Split(word, repoSeparator)
	switch len(parts) {
	case 1:
		return cmdTypeOwner
	case 2:
		return cmdTypeRepo
	default:
		return cmdTypeInvalid
	}
}

// ParseArgs parses the arguments
func ParseArgs(arguments []string) *Command {
	cmd := &Command{
		Args: normalizeArgs(arguments),
	}

	cmdType := judgeType(cmd.Args[0])
	switch cmdType {
	case cmdTypeConfig:
		cmd.createConfigSubCommand()
	case cmdTypeOwner:
		cmd.createOwnerSubCommand()
	case cmdTypeRepo:
		cmd.createRepoSubCommand()
	default:
		cmd.createHelpCommand()
	}

	return cmd
}
