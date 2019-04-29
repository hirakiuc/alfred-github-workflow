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

// Branch describe the github branch
type Branch struct {
	Name string
}

// ConvertBranches convert github.Branch to Branch
func ConvertBranches(branches []*github.Branch) []Branch {
	items := []Branch{}
	for _, branch := range branches {
		items = append(items, Branch{
			Name: branch.GetName(),
		})
	}
	return items
}

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

// Milestone describe the github milestone.
type Milestone struct {
	Description string
	HTMLURL     string
	Title       string
}

// ConvertMilestones convert github.Milestone to Milestone.
func ConvertMilestones(milestones []*github.Milestone) []Milestone {
	items := []Milestone{}
	for _, milestone := range milestones {
		items = append(items, Milestone{
			Description: milestone.GetDescription(),
			HTMLURL:     milestone.GetHTMLURL(),
			Title:       milestone.GetTitle(),
		})
	}
	return items
}
