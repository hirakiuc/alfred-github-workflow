package model

import "github.com/google/go-github/github"

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
