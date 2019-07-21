package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// BranchesCache describe an instance of branch cache
type BranchesCache struct {
	*BaseCache
}

// NewBranchesCache return an instance of branch cache store.
func NewBranchesCache(wf *aw.Workflow) *BranchesCache {
	return &BranchesCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *BranchesCache) getCacheKey(owner string, repo string) string {
	return fmt.Sprintf("branches-%s-%s", owner, repo)
}

// GetCache return the branch cache.
func (cache *BranchesCache) GetCache(owner string, repo string) ([]model.Branch, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	branches := []model.Branch{}
	err := cache.getRawCache(cacheKey, getMaxCacheAge(), &branches)
	if err != nil {
		return []model.Branch{}, err
	}
	if branches == nil {
		return []model.Branch{}, nil
	}

	return branches, nil
}

// Store stores the branches to the cache.
func (cache *BranchesCache) Store(owner string, repo string, branches []model.Branch) ([]model.Branch, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	_, err := cache.storeRawData(cacheKey, branches)
	if err != nil {
		return branches, err
	}

	return branches, nil
}
