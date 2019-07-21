package cli

import "github.com/hirakiuc/alfred-github-workflow/subcommand"

type RootCommandParser struct {
	tokenizer *Tokenizer
}

func NewRootCommandParser(tokenizer *Tokenizer) RootCommandParser {
	return RootCommandParser{
		tokenizer: tokenizer,
	}
}

func (p RootCommandParser) Parse() subcommand.SubCommand {
	return subcommand.NewHelpCommand()
}
