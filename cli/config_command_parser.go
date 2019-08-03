package cli

import (
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
	"github.com/hirakiuc/alfred-github-workflow/subcommand/config"
)

const (
	cmdTypeConfigToken      = "token"
	cmdTypeConfigClearCache = "clear-cache"
)

type ConfigCommandParser struct {
	tokenizer *Tokenizer
}

func NewConfigCommandParser(tokenizer *Tokenizer) ConfigCommandParser {
	return ConfigCommandParser{
		tokenizer: tokenizer,
	}
}

func (p ConfigCommandParser) Parse() subcommand.SubCommand {
	// '>', ...
	token := p.tokenizer.NextToken()
	opts := p.tokenizer.RestOfTokens()

	switch token {
	case cmdTypeConfigToken:
		token := ""
		if len(opts) > 0 {
			token = opts[0]
		}
		return config.NewTokenCommand(token)
	case cmdTypeConfigClearCache:
		return config.NewClearCacheCommand()
	default:
		options := append([]string{token}, opts...)
		return config.NewHelpCommand(options)
	}
}
