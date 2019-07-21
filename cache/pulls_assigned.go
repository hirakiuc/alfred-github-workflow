package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

type PullsAssignedCache struct {
	*BaseCache
}

func NewPullsAssignedCache(wf *aw.Workflow) *PullsAssignedCache {
	return &PullsAssignedCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *PullsAssignedCache) getCacheKey(user string) string {
	return fmt.Sprintf("pulls-assigned-%s", user)
}

func (cache *PullsAssignedCache) GetCache(user string) ([]model.Issue, error) {
	cacheKey := cache.getCacheKey(user)

	issues := []model.Issue{}
	err := cache.getRawCache(cacheKey, getMaxCacheAge(), &issues)
	if err != nil {
		return []model.Issue{}, err
	}
	if issues == nil {
		return []model.Issue{}, nil
	}

	return issues, nil
}

func (cache *PullsAssignedCache) Store(user string, issues []model.Issue) ([]model.Issue, error) {
	cacheKey := cache.getCacheKey(user)

	_, err := cache.storeRawData(cacheKey, issues)
	if err != nil {
		return issues, err
	}

	return issues, nil
}
