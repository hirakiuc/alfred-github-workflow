package cli

import (
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
	"github.com/hirakiuc/alfred-github-workflow/subcommand/my"
	"github.com/hirakiuc/alfred-github-workflow/subcommand/my/pulls"
)

const (
	cmdTypeMyPulls = "pulls"

	cmdTypeMyPullsReviewRequests = "review-requests"
)

type MyCommandParser struct {
	tokenizer *Tokenizer
}

func NewMyCommandParser(tokenizer *Tokenizer) MyCommandParser {
	return MyCommandParser{
		tokenizer: tokenizer,
	}
}

func (p MyCommandParser) Parse() subcommand.SubCommand {
	// 'my', cmd, ...
	token := p.tokenizer.NextToken()
	opts := p.tokenizer.RestOfTokens()

	switch token {
	case cmdTypeMyPulls:
		return p.parsePullsCommand()
	default:
		return my.NewHelpCommand(opts)
	}
}

func (p *MyCommandParser) parsePullsCommand() subcommand.SubCommand {
	// 'my', 'pulls', cmd, ...
	token := p.tokenizer.NextToken()
	opts := p.tokenizer.RestOfTokens()

	switch token {
	case cmdTypeMyPullsReviewRequests:
		return pulls.NewReviewRequestsCommand(opts)
	default:
		return pulls.NewHelpCommand(opts)
	}
}
