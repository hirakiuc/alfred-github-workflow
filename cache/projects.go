package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// ProjectsCache describe an instance of project cache.
type ProjectsCache struct {
	*BaseCache
}

// NewProjectsCache return an instance of project cache store.
func NewProjectsCache(wf *aw.Workflow) *ProjectsCache {
	return &ProjectsCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *ProjectsCache) getCacheKey(owner string, repo string) string {
	return fmt.Sprintf("projects-%s-%s", owner, repo)
}

// GetCache return the project cache.
func (cache *ProjectsCache) GetCache(owner string, repo string) ([]model.Project, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	projects := []model.Project{}
	err := cache.getRawCache(cacheKey, getMaxCacheAge(), &projects)
	if err != nil {
		return []model.Project{}, err
	}
	if projects == nil {
		return []model.Project{}, nil
	}

	return projects, nil
}

// Store stores the projects to the cache.
func (cache *ProjectsCache) Store(owner string, repo string, projects []model.Project) ([]model.Project, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	_, err := cache.storeRawData(cacheKey, projects)
	if err != nil {
		return projects, err
	}

	return projects, nil
}
