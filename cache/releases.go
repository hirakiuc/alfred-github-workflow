package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// ReleasesCache describe an instance of release cache.
type ReleasesCache struct {
	*BaseCache
}

// NewReleasesCache return an instance of release cache store.
func NewReleasesCache(wf *aw.Workflow) *ReleasesCache {
	return &ReleasesCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *ReleasesCache) getCacheKey(owner string, repo string) string {
	return fmt.Sprintf("releases-%s-%s", owner, repo)
}

// GetCache return the releases cache.
func (cache *ReleasesCache) GetCache(owner string, repo string) ([]model.Release, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	rels := []model.Release{}
	err := cache.getRawCache(cacheKey, getMaxCacheAge(), &rels)
	if err != nil {
		return []model.Release{}, err
	}
	if rels == nil {
		return []model.Release{}, nil
	}

	return rels, nil
}

// Store stores the rels to the cache.
func (cache *ReleasesCache) Store(owner string, repo string, rels []model.Release) ([]model.Release, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	_, err := cache.storeRawData(cacheKey, rels)
	if err != nil {
		return rels, err
	}

	return rels, nil
}
