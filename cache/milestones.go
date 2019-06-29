package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// MilestonesCache describe an instance of milestone cache.
type MilestonesCache struct {
	wf *aw.Workflow
}

// NewMilestonesCache return an instance of milestone cache store.
func NewMilestonesCache(wf *aw.Workflow) *MilestonesCache {
	return &MilestonesCache{
		wf: wf,
	}
}

func (cache MilestonesCache) getCacheKey(owner string, repo string) string {
	return fmt.Sprintf("milestones-%s-%s", owner, repo)
}

// GetCache return the issue cache.
func (cache *MilestonesCache) GetCache(owner string, repo string) ([]model.Milestone, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	store := cache.wf.Cache

	milestones := []model.Milestone{}
	if !store.Exists(cacheKey) {
		return []model.Milestone{}, nil
	}

	if store.Expired(cacheKey, getMaxCacheAge()) {
		return []model.Milestone{}, nil
	}

	if err := store.LoadJSON(cacheKey, &milestones); err != nil {
		cache.wf.FatalError(err)
		return []model.Milestone{}, err
	}

	return milestones, nil
}

// Store stores the branches to the cache.
func (cache *MilestonesCache) Store(owner string, repo string, branches []model.Milestone) ([]model.Milestone, error) {
	cacheKey := cache.getCacheKey(owner, repo)
	store := cache.wf.Cache

	if err := store.StoreJSON(cacheKey, branches); err != nil {
		cache.wf.FatalError(err)
		return []model.Milestone{}, err
	}

	return branches, nil
}
