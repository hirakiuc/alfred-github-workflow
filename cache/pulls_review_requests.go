package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

type PullsReviewRequestsCache struct {
	*BaseCache
}

func NewPullsReviewRequestsCache(wf *aw.Workflow) *PullsReviewRequestsCache {
	return &PullsReviewRequestsCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *PullsReviewRequestsCache) getCacheKey(user string) string {
	return fmt.Sprintf("review-requests-%s", user)
}

func (cache *PullsReviewRequestsCache) GetCache(user string) ([]model.Issue, error) {
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

func (cache *PullsReviewRequestsCache) Store(user string, issues []model.Issue) ([]model.Issue, error) {
	cacheKey := cache.getCacheKey(user)

	_, err := cache.storeRawData(cacheKey, issues)
	if err != nil {
		return issues, err
	}

	return issues, nil
}
