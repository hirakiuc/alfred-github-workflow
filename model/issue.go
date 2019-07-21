package model

import (
	"fmt"
	"time"

	"github.com/google/go-github/github"
)

const (
	// IssueStateOpen describe the opened status
	IssueStateOpen = "open"
	// IssueStatusClosed describe the closed status
	IssueStatusClosed = "closed"
)

// Issue describe the github issue.
type Issue struct {
	Number    int
	Title     string
	State     string
	HTMLURL   string
	Reporter  string
	CreatedAt time.Time
	ClosedAt  time.Time

	IsPullRequest bool
}

func formatTime(t time.Time) string {
	return t.Format("2019/01/01")
}

// GetItemTitle return the title text for alfred item.
func (i Issue) GetItemTitle() string {
	return i.Title
}

// GetItemSubtitle return the subtitle text for alfred item.
func (i Issue) GetItemSubtitle() string {
	switch i.State {
	case IssueStateOpen:
		return fmt.Sprintf("#%d opened on %s by %s", i.Number, formatTime(i.CreatedAt), i.Reporter)
	case IssueStatusClosed:
		return fmt.Sprintf("#%d by %s was closed on %s", i.Number, i.Reporter, formatTime(i.ClosedAt))
	default:
		return fmt.Sprintf("#%d opened on %s by %s. state:%s", i.Number, formatTime(i.CreatedAt), i.Reporter, i.State)
	}
}

// ConvertIssues convert array of github.Issue to array of Issue
func ConvertIssues(issues []*github.Issue) []Issue {
	items := []Issue{}
	for _, issue := range issues {
		items = append(items, ConvertIssue(issue))
	}
	return items
}

// ConvertIssue convert github.Issue to Issue
func ConvertIssue(issue *github.Issue) Issue {
	return Issue{
		Number:    issue.GetNumber(),
		Title:     issue.GetTitle(),
		State:     issue.GetState(),
		HTMLURL:   issue.GetHTMLURL(),
		Reporter:  issue.GetUser().GetName(),
		CreatedAt: issue.GetCreatedAt(),
		ClosedAt:  issue.GetClosedAt(),

		IsPullRequest: (issue.GetPullRequestLinks() != nil),
	}
}
