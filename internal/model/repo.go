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
