package model

import "github.com/google/go-github/github"

// Issue describe the github issue.
type Issue struct {
	Number int
	Title  string
	State  string
	URL    string
}

// ConvertIssues convert github.Issue to Issue
func ConvertIssues(issues []*github.Issue) []Issue {
	items := []Issue{}
	for _, issue := range issues {
		items = append(items, Issue{
			Number: issue.GetNumber(),
			Title:  issue.GetTitle(),
			State:  issue.GetState(),
			URL:    issue.GetHTMLURL(),
		})
	}
	return items
}
