package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// MilestonesCache describe an instance of milestone cache.
type MilestonesCache struct {
	*BaseCache
}

// NewMilestonesCache return an instance of milestone cache store.
func NewMilestonesCache(wf *aw.Workflow) *MilestonesCache {
	return &MilestonesCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache MilestonesCache) getCacheKey(owner string, repo string) string {
	return fmt.Sprintf("milestones-%s-%s", owner, repo)
}

// GetCache return the issue cache.
func (cache *MilestonesCache) GetCache(owner string, repo string) ([]model.Milestone, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	milestones := []model.Milestone{}
	err := cache.getRawCache(cacheKey, getMaxCacheAge(), &milestones)
	if err != nil {
		return []model.Milestone{}, err
	}
	if milestones == nil {
		return []model.Milestone{}, nil
	}

	return milestones, nil
}

// Store stores the branches to the cache.
func (cache *MilestonesCache) Store(
	owner string, repo string, milestones []model.Milestone) ([]model.Milestone, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	_, err := cache.storeRawData(cacheKey, milestones)
	if err != nil {
		return milestones, err
	}

	return milestones, nil
}
