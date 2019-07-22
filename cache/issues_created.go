package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

type IssuesCreatedCache struct {
	*BaseCache
}

func NewIssuesCreatedCache(wf *aw.Workflow) *IssuesCreatedCache {
	return &IssuesCreatedCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *IssuesCreatedCache) getCacheKey(user string) string {
	return fmt.Sprintf("issues-created-%s", user)
}

func (cache *IssuesCreatedCache) GetCache(user string) ([]model.Issue, error) {
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

func (cache *IssuesCreatedCache) Store(user string, issues []model.Issue) ([]model.Issue, error) {
	cacheKey := cache.getCacheKey(user)

	_, err := cache.storeRawData(cacheKey, issues)
	if err != nil {
		return issues, err
	}

	return issues, nil
}
