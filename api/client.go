package api

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/google/go-github/github"
	"github.com/hirakiuc/alfred-github-workflow/secret"
	"golang.org/x/oauth2"
)

// Client describe an instance of API client.
type Client struct {
	github *github.Client
}

func fetchAPIToken(wf *aw.Workflow) (string, error) {
	store := secret.NewStore(wf)
	return store.GetAPIToken()
}

// NewClient return a instance of github client.
func NewClient(ctx context.Context, wf *aw.Workflow) (*Client, error) {
	token, err := fetchAPIToken(wf)
	if err != nil {
		return nil, err
	}

	if token == "" {
		return &Client{github: github.NewClient(nil)}, nil
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &Client{
		github: github.NewClient(tc),
	}, nil
}
