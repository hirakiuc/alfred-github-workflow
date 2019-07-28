package model

import (
	"fmt"

	"github.com/google/go-github/github"
)

type Release struct {
	Author      string           `json:"author"`
	Name        string           `json:"name"`
	Draft       bool             `json:"draft"`
	Prerelease  bool             `json:"prerelease"`
	HTMLURL     string           `json:"html_url"`
	CreatedAt   github.Timestamp `json:"created_at"`
	PublishedAt github.Timestamp `json:"published_at"`
}

func ConvertReleases(rels []*github.RepositoryRelease) []Release {
	items := []Release{}
	for _, rel := range rels {
		items = append(items, ConvertRelease(rel))
	}
	return items
}

func ConvertRelease(rel *github.RepositoryRelease) Release {
	return Release{
		Author:      authorName(rel.GetAuthor()),
		Name:        rel.GetName(),
		Draft:       rel.GetDraft(),
		Prerelease:  rel.GetPrerelease(),
		HTMLURL:     rel.GetHTMLURL(),
		CreatedAt:   rel.GetCreatedAt(),
		PublishedAt: rel.GetPublishedAt(),
	}
}

func (m Release) Subtitle() string {
	return fmt.Sprintf("%s released this on %s", m.Author, m.PublishedAt.String())
}

func authorName(user *github.User) string {
	if user == nil {
		return ""
	}

	return user.GetName()
}
