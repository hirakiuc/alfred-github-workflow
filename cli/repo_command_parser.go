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
	token := p.tokenizer.NextToken()
	opts := p.tokenizer.RestOfTokens()

	switch token {
	case "branches":
		return repo.NewBranchesCommand(p.Owner, p.Repo, opts)
	case "issues":
		return repo.NewIssueCommand(p.Owner, p.Repo, opts)
	case "milestones":
		return repo.NewMilestonesCommand(p.Owner, p.Repo, opts)
	case "pulls":
		return repo.NewPullsCommand(p.Owner, p.Repo, opts)
	case "new":
		return repo.NewNewCommand(p.Owner, p.Repo, opts)
	case "releases":
		return repo.NewReleasesCommand(p.Owner, p.Repo, opts)
	default:
		return repo.NewHelpCommand(p.Owner, p.Repo, opts)
	}
}
