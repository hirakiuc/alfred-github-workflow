package config

import (
	"context"
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/secret"
)

// TokenCommand describe a subcommand to configure the api token.
type TokenCommand struct {
	Token string
}

// NewTokenCommand return an instance of TokenCommand.
func NewTokenCommand(token string) TokenCommand {
	return TokenCommand{
		Token: token,
	}
}

// Run start this subcommand.
func (cmd TokenCommand) Run(ctx context.Context, wf *aw.Workflow) {
	store := secret.NewStore(wf)

	if len(cmd.Token) == 0 {
		token, err := store.GetAPIToken()
		if err != nil {
			wf.FatalError(err)
			return
		}
		if len(token) == 0 {
			wf.WarnEmpty("No token found.", "")
			return
		}

		wf.NewItem(fmt.Sprintf("Found token:%s", token))
		return
	}

	err := store.Store(secret.KeyGithubAPIToken, cmd.Token)
	if err != nil {
		wf.FatalError(err)
		return
	}

	wf.NewItem("Success.")
}
