package model

import "github.com/google/go-github/github"

// PullRequest describe the github pull request.
type PullRequest struct {
	Number  int
	State   string
	Title   string
	HTMLURL string
}

// ConvertPullRequests convert github.PullRequest to PullRequest
func ConvertPullRequests(pulls []*github.PullRequest) []PullRequest {
	items := []PullRequest{}
	for _, pull := range pulls {
		items = append(items, PullRequest{
			Number:  pull.GetNumber(),
			State:   pull.GetState(),
			Title:   pull.GetTitle(),
			HTMLURL: pull.GetHTMLURL(),
		})
	}
	return items
}
