package cli

import (
	"strings"

	"github.com/hirakiuc/alfred-github-workflow/subcommand"
)

const (
	repoSeparator = "/"

	// Token types
	emptyToken   = ""
	configToken  = ">"
	myToken      = "my"
	ownerToken   = "owner"
	repoToken    = "owner/repo"
	invalidToken = "---invalid---"
)

type Parser struct {
	tokenizer *Tokenizer
}

type SubCommandParser interface {
	Parse() subcommand.SubCommand
}

func NewParser() *Parser {
	return &Parser{
		tokenizer: NewTokenizer(),
	}
}

func ParseArgs(args []string) subcommand.SubCommand {
	parser := NewParser()
	return parser.Parse(args)
}

func judgeOwnerOrSlug(token string) string {
	parts := strings.Split(token, repoSeparator)
	switch len(parts) {
	case 1:
		return ownerToken
	case 2:
		if len(parts[0]) > 0 {
			if len(parts[1]) > 0 {
				return repoToken
			}
			return ownerToken
		}
		return invalidToken
	default:
		return invalidToken
	}
}

func (parser *Parser) Parse(args []string) subcommand.SubCommand {
	parser.tokenizer.Tokenize(args)

	token := parser.tokenizer.NextToken()
	subCommandParser := parser.createSubcommandParser(token)

	return subCommandParser.Parse()
}

func (parser *Parser) createSubcommandParser(token string) SubCommandParser {
	switch token {
	case configToken:
		return NewConfigCommandParser(parser.tokenizer)
	case myToken:
		return NewMyCommandParser(parser.tokenizer)
	case emptyToken:
		return NewRootCommandParser(parser.tokenizer)
	default:
		// check the token is user name only or slug
		tokenType := judgeOwnerOrSlug(token)
		switch tokenType {
		case ownerToken:
			return NewOwnerCommandParser(parser.tokenizer, token)
		case repoToken:
			parts := strings.Split(token, repoSeparator)
			return NewRepoCommandParser(parser.tokenizer, parts[0], parts[1])
		default:
			return NewRootCommandParser(parser.tokenizer)
		}
	}
}
