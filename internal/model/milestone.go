package model

import (
	"fmt"
	"time"

	"github.com/google/go-github/github"
)

// Milestone describe the github milestone.
type Milestone struct {
	Description  string
	HTMLURL      string
	Title        string
	State        string
	DueOn        time.Time
	ClosedAt     time.Time
	OpenIssues   int
	ClosedIssues int
}

func (m Milestone) formatDueDate() string {
	if m.DueOn.IsZero() {
		return "No due date"
	}

	return formatTime(m.DueOn)
}

// GetItemTitle return title string
func (m Milestone) GetItemTitle() string {
	return m.Title
}

// GetProgress return the progress of the milestone
func (m Milestone) GetProgress() float64 {
	total := m.OpenIssues + m.ClosedIssues
	if total == 0 {
		return float64(0)
	}

	return float64(m.ClosedIssues) / float64(total) * 100
}

// GetItemSubtitle return the subtitle string.
func (m Milestone) GetItemSubtitle() string {
	switch m.State {
	case "open":
		return fmt.Sprintf("%s, progress:%.0f (open: %d, closed: %d)",
			m.formatDueDate(), m.GetProgress(), m.OpenIssues, m.ClosedIssues)
	case "closed":
		return fmt.Sprintf("Closed on %s (open: %d, closed: %d)", formatTime(m.ClosedAt), m.OpenIssues, m.ClosedIssues)
	default:
		return fmt.Sprintf("%s progress:%.0f (open: %d, closed: %d)", m.State, m.GetProgress(), m.OpenIssues, m.ClosedIssues)
	}
}

// ConvertMilestones convert github.Milestone to Milestone.
func ConvertMilestones(milestones []*github.Milestone) []Milestone {
	items := []Milestone{}
	for _, milestone := range milestones {
		items = append(items, Milestone{
			Description:  milestone.GetDescription(),
			HTMLURL:      milestone.GetHTMLURL(),
			Title:        milestone.GetTitle(),
			State:        milestone.GetState(),
			DueOn:        milestone.GetDueOn(),
			ClosedAt:     milestone.GetClosedAt(),
			OpenIssues:   milestone.GetOpenIssues(),
			ClosedIssues: milestone.GetClosedIssues(),
		})
	}
	return items
}
