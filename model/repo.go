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

func (repo Repo) BranchURL() string {
	return repo.HTMLURL + "/branches"
}

func (repo Repo) IssuesURL() string {
	return repo.HTMLURL + "/issues"
}

func (repo Repo) MilestoneURL() string {
	return repo.HTMLURL + "/milestones"
}

func (repo Repo) PullsURL() string {
	return repo.HTMLURL + "/pulls"
}

func (repo Repo) ProjectsURL() string {
	return repo.HTMLURL + "/projects"
}

func (repo Repo) WikiURL() string {
	return repo.HTMLURL + "/wiki"
}

func (repo Repo) SecurityURL() string {
	return repo.HTMLURL + "/network/alerts"
}

func (repo Repo) InsightsURL() string {
	return repo.HTMLURL + "/pulse"
}

func (repo Repo) SettingsURL() string {
	return repo.HTMLURL + "/settings"
}

func (repo Repo) CloneURL() string {
	return repo.HTMLURL + ".git"
}
