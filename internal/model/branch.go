package model

import (
	"strings"

	"github.com/google/go-github/github"
)

// Branch describe the github branch
type Branch struct {
	Owner   string
	Repo    string
	Name    string
	HTMLURL string
}

func htmlURL(owner string, repo string, name string) string {
	components := []string{
		"https://github.com",
		owner,
		repo,
		"tree",
		name,
	}

	return strings.Join(components, "/")
}

// ConvertBranch convert github.Branch to Branch
func ConvertBranch(owner string, repo string, branch *github.Branch) Branch {
	return Branch{
		Owner:   owner,
		Repo:    repo,
		Name:    branch.GetName(),
		HTMLURL: htmlURL(owner, repo, branch.GetName()),
	}
}
