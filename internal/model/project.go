package model

import "github.com/google/go-github/github"

// Project describe the github project.
type Project struct {
	Name string
	URL  string
}

// ConvertProjects convert github.Project to Project
func ConvertProjects(projects []*github.Project) []Project {
	items := []Project{}
	for _, project := range projects {
		items = append(items, Project{
			Name: project.GetName(),
			URL:  project.GetURL(),
		})
	}
	return items
}
