package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

type ReviewRequestsCache struct {
	wf *aw.Workflow
}

func NewReviewRequestsCache(wf *aw.Workflow) *ReviewRequestsCache {
	return &ReviewRequestsCache{
		wf: wf,
	}
}

func (cache *ReviewRequestsCache) getCacheKey(user string) string {
	return fmt.Sprintf("review-requests-%s", user)
}

func (cache *ReviewRequestsCache) GetCache(user string) ([]model.Issue, error) {
	cacheKey := cache.getCacheKey(user)

	store := cache.wf.Cache
	issues := []model.Issue{}
	if !store.Exists(cacheKey) {
		return issues, nil
	}

	if store.Expired(cacheKey, getMaxCacheAge()) {
		return issues, nil
	}

	if err := store.LoadJSON(cacheKey, &issues); err != nil {
		cache.wf.FatalError(err)
		return []model.Issue{}, nil
	}

	return issues, nil
}

func (cache *ReviewRequestsCache) Store(user string, issues []model.Issue) ([]model.Issue, error) {
	cacheKey := cache.getCacheKey(user)
	store := cache.wf.Cache

	if err := store.StoreJSON(cacheKey, issues); err != nil {
		cache.wf.FatalError(err)
		return []model.Issue{}, err
	}

	return issues, nil
}
