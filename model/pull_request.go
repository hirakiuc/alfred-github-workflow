package model

import (
	"fmt"
	"time"

	"github.com/google/go-github/github"
)

// PullRequest describe the github pull request.
type PullRequest struct {
	Number    int
	State     string
	Title     string
	HTMLURL   string
	User      string
	MergedBy  string
	CreatedAt time.Time
	ClosedAt  time.Time
	MergedAt  time.Time
}

// GetItemTitle return a title string.
func (pull PullRequest) GetItemTitle() string {
	return pull.Title
}

// GetItemSubtitle return a subtitle string.
func (pull PullRequest) GetItemSubtitle() string {
	switch pull.State {
	case "open":
		return fmt.Sprintf("#%d opened on %s by %s", pull.Number, formatTime(pull.CreatedAt), pull.User)
	case "closed":
		return fmt.Sprintf("#%d by %s was merged on %s", pull.Number, pull.User, pull.MergedAt)
	default:
		return fmt.Sprintf("#%d state:%s by %s", pull.Number, pull.State, pull.User)
	}
}

// ConvertPullRequests convert github.PullRequest to PullRequest
func ConvertPullRequests(pulls []*github.PullRequest) []PullRequest {
	items := []PullRequest{}
	for _, pull := range pulls {
		items = append(items, PullRequest{
			Number:    pull.GetNumber(),
			State:     pull.GetState(),
			Title:     pull.GetTitle(),
			HTMLURL:   pull.GetHTMLURL(),
			User:      pull.GetUser().GetName(),
			MergedBy:  pull.GetUser().GetName(),
			CreatedAt: pull.GetCreatedAt(),
			ClosedAt:  pull.GetClosedAt(),
			MergedAt:  pull.GetMergedAt(),
		})
	}
	return items
}
