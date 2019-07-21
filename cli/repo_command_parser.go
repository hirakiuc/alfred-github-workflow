package cli

import (
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
	"github.com/hirakiuc/alfred-github-workflow/subcommand/repo"
)

type RepoCommandParser struct {
	tokenizer *Tokenizer
	Repo      string
	Owner     string
}

func NewRepoCommandParser(tokenizer *Tokenizer, owner string, repo string) RepoCommandParser {
	return RepoCommandParser{
		tokenizer: tokenizer,
		Repo:      repo,
		Owner:     owner,
	}
}

func (p RepoCommandParser) Parse() subcommand.SubCommand {
	// owner/repo, ...
	opts := p.tokenizer.RestOfTokens()
	if len(opts) == 0 {
		return repo.NewHelpCommand(p.Owner, p.Repo, []string{})
	}

	switch opts[0] {
	case "branches":
		return repo.NewBranchesCommand(p.Owner, p.Repo, opts)
	case "issues":
		return repo.NewIssueCommand(p.Owner, p.Repo, opts)
	case "milestones":
		return repo.NewMilestonesCommand(p.Owner, p.Repo, opts)
	case "pulls":
		return repo.NewPullsCommand(p.Owner, p.Repo, opts)
	default:
		return repo.NewHelpCommand(p.Owner, p.Repo, opts)
	}
}
