package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// ReposCache describe an instance of repository cache.
type ReposCache struct {
	wf *aw.Workflow
}

// NewReposCache return an instance of repository cache store.
func NewReposCache(wf *aw.Workflow) *ReposCache {
	return &ReposCache{
		wf: wf,
	}
}

func (cache *ReposCache) getCacheKey(owner string) string {
	return fmt.Sprintf("repos-%s", owner)
}

// GetCache return the repository cache.
func (cache *ReposCache) GetCache(owner string) ([]model.Repo, error) {
	cacheKey := cache.getCacheKey(owner)

	store := cache.wf.Cache

	repos := []model.Repo{}
	if !store.Exists(cacheKey) {
		return []model.Repo{}, nil
	}

	if store.Expired(cacheKey, getMaxCacheAge()) {
		return []model.Repo{}, nil
	}

	if err := store.LoadJSON(cacheKey, &repos); err != nil {
		cache.wf.FatalError(err)
		return []model.Repo{}, err
	}

	return repos, nil
}

// Store stores the repos to the cache.
func (cache *ReposCache) Store(owner string, repos []model.Repo) ([]model.Repo, error) {
	cacheKey := cache.getCacheKey(owner)
	store := cache.wf.Cache

	if err := store.StoreJSON(cacheKey, repos); err != nil {
		cache.wf.FatalError(err)
		return []model.Repo{}, err
	}

	return repos, nil
}
