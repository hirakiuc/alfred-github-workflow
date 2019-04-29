package model

import "github.com/google/go-github/github"

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
