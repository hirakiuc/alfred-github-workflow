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
		items = append(items, ConvertRepo(item))
	}

	return items
}

func ConvertRepo(repo *github.Repository) Repo {
	return Repo{
		Owner:       repo.Owner.GetName(),
		Name:        repo.GetName(),
		Description: repo.GetDescription(),
		HTMLURL:     repo.GetHTMLURL(),
	}
}
