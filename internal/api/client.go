package api

import (
	"github.com/google/go-github/github"
)

type GithubClient struct {
	*github.Client
}

func GetClient() *GithubClient {
	return &GithubClient{
		github.NewClient(nil),
	}
}
