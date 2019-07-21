package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

type ReviewRequestsCache struct {
	*BaseCache
}

func NewReviewRequestsCache(wf *aw.Workflow) *ReviewRequestsCache {
	return &ReviewRequestsCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *ReviewRequestsCache) getCacheKey(user string) string {
	return fmt.Sprintf("review-requests-%s", user)
}

func (cache *ReviewRequestsCache) GetCache(user string) ([]model.Issue, error) {
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

func (cache *ReviewRequestsCache) Store(user string, issues []model.Issue) ([]model.Issue, error) {
	cacheKey := cache.getCacheKey(user)

	_, err := cache.storeRawData(cacheKey, issues)
	if err != nil {
		return issues, err
	}

	return issues, nil
}
