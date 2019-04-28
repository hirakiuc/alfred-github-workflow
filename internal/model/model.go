package model

import "github.com/google/go-github/github"

// Repo describe the github repository.
type Repo struct {
	Owner       string `json:"owner"`
	Name        string `json:"name"`
	Description string `json:"description,omitempty"`
	HTMLURL     string `json:"html_url"`
}

// ConvertRepos convert the github.Repository to Repo.
func ConvertRepos(repos []*github.Repository) []Repo {
	items := []Repo{}
	for _, item := range repos {
		items = append(items, Repo{
			Owner:       item.Owner.GetName(),
			Name:        item.GetName(),
			Description: item.GetDescription(),
			HTMLURL:     item.GetHTMLURL(),
		})
	}
	return items
}

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
