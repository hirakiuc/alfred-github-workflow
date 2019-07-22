package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

type IssuesAssignedCache struct {
	*BaseCache
}

func NewIssuesAssignedCache(wf *aw.Workflow) *IssuesAssignedCache {
	return &IssuesAssignedCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *IssuesAssignedCache) getCacheKey(user string) string {
	return fmt.Sprintf("issues-assigned-%s", user)
}

func (cache *IssuesAssignedCache) GetCache(user string) ([]model.Issue, error) {
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

func (cache *IssuesAssignedCache) Store(user string, issues []model.Issue) ([]model.Issue, error) {
	cacheKey := cache.getCacheKey(user)

	_, err := cache.storeRawData(cacheKey, issues)
	if err != nil {
		return issues, err
	}

	return issues, nil
}
