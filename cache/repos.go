package cache

import (
	"fmt"
	"time"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// ReposCache describe an instance of repository cache.
type ReposCache struct {
	*BaseCache
}

// NewReposCache return an instance of repository cache store.
func NewReposCache(wf *aw.Workflow) *ReposCache {
	return &ReposCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *ReposCache) getCacheKey(owner string) string {
	return fmt.Sprintf("repos-%s", owner)
}

func getMaxRepoCacheAge() time.Duration {
	return time.Duration(24*7) * time.Hour
}

// GetCache return the repository cache.
func (cache *ReposCache) GetCache(owner string) ([]model.Repo, error) {
	cacheKey := cache.getCacheKey(owner)

	repos := []model.Repo{}
	err := cache.getRawCache(cacheKey, getMaxRepoCacheAge(), &repos)
	if err != nil {
		return []model.Repo{}, err
	}
	if repos == nil {
		return []model.Repo{}, nil
	}

	return repos, nil
}

// Store stores the repos to the cache.
func (cache *ReposCache) Store(owner string, repos []model.Repo) ([]model.Repo, error) {
	cacheKey := cache.getCacheKey(owner)

	_, err := cache.storeRawData(cacheKey, repos)
	if err != nil {
		return repos, err
	}

	return repos, nil
}
