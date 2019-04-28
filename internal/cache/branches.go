package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// BranchesCache describe an instance of branch cache
type BranchesCache struct {
	wf *aw.Workflow
}

// NewBranchesCache return an instance of branch cache store.
func NewBranchesCache(wf *aw.Workflow) *BranchesCache {
	return &BranchesCache{
		wf: wf,
	}
}

func (cache *BranchesCache) getCacheKey(owner string, repo string) string {
	return fmt.Sprintf("branches-%s-%s", owner, repo)
}

// GetCache return the branch cache.
func (cache *BranchesCache) GetCache(owner string, repo string) ([]model.Branch, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	store := cache.wf.Cache

	branches := []model.Branch{}
	if !store.Exists(cacheKey) {
		return []model.Branch{}, nil
	}

	if store.Expired(cacheKey, maxCacheAge) {
		return []model.Branch{}, nil
	}

	if err := store.LoadJSON(cacheKey, &branches); err != nil {
		cache.wf.FatalError(err)
		return []model.Branch{}, nil
	}

	return branches, nil
}

// Store stores the branches to the cache.
func (cache *BranchesCache) Store(owner string, repo string, branches []model.Branch) ([]model.Branch, error) {
	cacheKey := cache.getCacheKey(owner, repo)
	store := cache.wf.Cache

	if err := store.StoreJSON(cacheKey, branches); err != nil {
		cache.wf.FatalError(err)
		return []model.Branch{}, err
	}

	return branches, nil
}
