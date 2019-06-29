package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// ProjectsCache describe an instance of project cache.
type ProjectsCache struct {
	wf *aw.Workflow
}

// NewProjectsCache return an instance of project cache store.
func NewProjectsCache(wf *aw.Workflow) *ProjectsCache {
	return &ProjectsCache{
		wf: wf,
	}
}

func (cache *ProjectsCache) getCacheKey(owner string, repo string) string {
	return fmt.Sprintf("projects-%s-%s", owner, repo)
}

// GetCache return the project cache.
func (cache *ProjectsCache) GetCache(owner string, repo string) ([]model.Project, error) {
	cacheKey := cache.getCacheKey(owner, repo)
	store := cache.wf.Cache

	projects := []model.Project{}
	if !store.Exists(cacheKey) {
		return projects, nil
	}

	if store.Expired(cacheKey, getMaxCacheAge()) {
		return projects, nil
	}

	if err := store.LoadJSON(cacheKey, &projects); err != nil {
		cache.wf.FatalError(err)
		return []model.Project{}, err
	}

	return projects, nil
}

// Store stores the projects to the cache.
func (cache *ProjectsCache) Store(owner string, repo string, projects []model.Project) ([]model.Project, error) {
	cacheKey := cache.getCacheKey(owner, repo)
	store := cache.wf.Cache

	if err := store.StoreJSON(cacheKey, projects); err != nil {
		cache.wf.FatalError(err)
		return []model.Project{}, err
	}

	return projects, nil
}
