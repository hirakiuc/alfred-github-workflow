package api

import (
	"context"
	"os"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
)

// Client describe an instance of API client.
type Client struct {
	github *github.Client
}

// NewClient return a instance of github client.
func NewClient(ctx context.Context) *Client {
	token := os.Getenv("GITHUB_API_TOKEN")
	if token == "" {
		return &Client{github: github.NewClient(nil)}
	}

	ts := oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	)
	tc := oauth2.NewClient(ctx, ts)

	return &Client{
		github: github.NewClient(tc),
	}
}
