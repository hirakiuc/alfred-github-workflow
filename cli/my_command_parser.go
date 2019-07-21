package cli

import (
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
	"github.com/hirakiuc/alfred-github-workflow/subcommand/my"
	"github.com/hirakiuc/alfred-github-workflow/subcommand/my/pulls"
)

const (
	cmdTypeMyPulls = "pulls"

	cmdTypeMyPullsCreated        = "created"
	cmdTypeMyPullsMentioned      = "mentioned"
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
	case cmdTypeMyPullsCreated:
		return pulls.NewCreatedCommand(opts)
	case cmdTypeMyPullsMentioned:
		return pulls.NewMentionedCommand(opts)
	case cmdTypeMyPullsReviewRequests:
		return pulls.NewReviewRequestsCommand(opts)
	default:
		return pulls.NewHelpCommand(opts)
	}
}
