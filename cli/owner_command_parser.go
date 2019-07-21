package cli

import "github.com/hirakiuc/alfred-github-workflow/subcommand"

type OwnerCommandParser struct {
	tokenizer *Tokenizer
	Owner     string
}

func NewOwnerCommandParser(tokenizer *Tokenizer, owner string) OwnerCommandParser {
	return OwnerCommandParser{
		tokenizer: tokenizer,
		Owner:     owner,
	}
}

func (p OwnerCommandParser) Parse() subcommand.SubCommand {
	// owner, ...
	opts := p.tokenizer.RestOfTokens()
	return subcommand.NewReposCommand(p.Owner, opts)
}
