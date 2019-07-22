package cli

import (
	"github.com/hirakiuc/alfred-github-workflow/subcommand"
	"github.com/hirakiuc/alfred-github-workflow/subcommand/my"
	"github.com/hirakiuc/alfred-github-workflow/subcommand/my/issues"
	"github.com/hirakiuc/alfred-github-workflow/subcommand/my/pulls"
)

const (
	cmdTypeMyPulls  = "pulls"
	cmdTypeMyIssues = "issues"

	cmdTypeMyPullsAssigned       = "assigned"
	cmdTypeMyPullsCreated        = "created"
	cmdTypeMyPullsMentioned      = "mentioned"
	cmdTypeMyPullsReviewRequests = "review-requests"

	cmdTypeMyIssuesCreated   = "created"
	cmdTypeMyIssuesAssigned  = "assigned"
	cmdTypeMyIssuesMentioned = "mentioned"
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
	case cmdTypeMyIssues:
		return p.parseIssuesCommand()
	default:
		return my.NewHelpCommand(opts)
	}
}

func (p *MyCommandParser) parsePullsCommand() subcommand.SubCommand {
	// 'my', 'pulls', cmd, ...
	token := p.tokenizer.NextToken()
	opts := p.tokenizer.RestOfTokens()

	switch token {
	case cmdTypeMyPullsAssigned:
		return pulls.NewAssignedCommand(opts)
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

func (p *MyCommandParser) parseIssuesCommand() subcommand.SubCommand {
	// 'my', 'issues', cmd, ...
	token := p.tokenizer.NextToken()
	opts := p.tokenizer.RestOfTokens()

	switch token {
	case cmdTypeMyIssuesCreated:
		return issues.NewCreatedCommand(opts)
	case cmdTypeMyIssuesAssigned:
		return issues.NewAssignedCommand(opts)
	case cmdTypeMyIssuesMentioned:
		return issues.NewMentionedCommand(opts)
	default:
		return issues.NewHelpCommand(opts)
	}
}
