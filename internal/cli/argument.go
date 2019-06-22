package cli

import (
	"strings"

	"github.com/hirakiuc/alfred-github-workflow/internal/subcommand"
	configcmd "github.com/hirakiuc/alfred-github-workflow/internal/subcommand/config"
	repocmd "github.com/hirakiuc/alfred-github-workflow/internal/subcommand/repo"
)

// Slug...
type Slug struct {
	Owner string
	Repo  string
}

// Args...
type Args struct {
	Args   []string
	Slug   *Slug
	Subcmd subcommand.SubCommand
}

const (
	repoSeparator = "/"

	cmdTypeEmpty   = ""
	cmdTypeConfig  = ">"
	cmdTypeOwner   = "owner"
	cmdTypeRepo    = "repo"
	cmdTypeInvalid = "invalid"
)

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

func getConfigSubCommand(args []string) subcommand.SubCommand {
	cmd, opts := parseConfigCommandArgs(args)

	switch cmd {
	case "token":
		token := ""
		if len(opts) > 0 {
			token = opts[0]
		}

		return configcmd.NewTokenCommand(token)
	default:
		return configcmd.NewHelpCommand()
	}
}

func getOwnerSubCommand(owner string, args []string) subcommand.SubCommand {
	return subcommand.NewReposCommand(owner, args)
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

func getRepoSubCommand(owner string, repo string, args []string) subcommand.SubCommand {
	cmd, options := parseSubCommandArgs(args)
	query := strings.Join(options, " ")

	switch cmd {
	case "issues":
		return repocmd.NewIssueCommand(owner, repo, query)
	case "pulls":
		return repocmd.NewPullsCommand(owner, repo, query)
	case "branches":
		return repocmd.NewBranchesCommand(owner, repo, query)
	case "milestones":
		return repocmd.NewMilestonesCommand(owner, repo, query)
	case "projects":
		return repocmd.NewProjectsCommand(owner, repo, query)
	default:
		// Show the subcommands
		return repocmd.NewHelpCommand(owner, repo)
	}
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
func ParseArgs(arguments []string) *Args {
	args := Args{
		Args: normalizeArgs(arguments),
	}

	cmdType := judgeType(args.Args[0])
	switch cmdType {
	case cmdTypeConfig:
		args.Slug = nil
		args.Subcmd = getConfigSubCommand(args.Args)
	case cmdTypeOwner:
		args.Slug = &Slug{Owner: args.Args[0]}
		args.Subcmd = getOwnerSubCommand(args.Args[0], args.Args[1:])
	case cmdTypeRepo:
		slug := strings.Split(args.Args[0], repoSeparator)
		args.Slug = &Slug{Owner: slug[0], Repo: slug[1]}
		args.Subcmd = getRepoSubCommand(slug[0], slug[1], args.Args[1:])
	default:
		args.Subcmd = subcommand.NewHelpCommand()
	}

	return &args
}
