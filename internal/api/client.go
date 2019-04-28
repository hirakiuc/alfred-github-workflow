package api

import "github.com/google/go-github/github"

// Client describe an instance of API client.
type Client struct {
	github *github.Client
}

// NewClient return a instance of github client.
func NewClient() *Client {
	return &Client{
		github: github.NewClient(nil),
	}
}
