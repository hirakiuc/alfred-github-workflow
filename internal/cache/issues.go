package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// IssuesCache describe an instance of issue cache.
type IssuesCache struct {
	wf *aw.Workflow
}

// NewIssuesCache return an instance of issue cache store.
func NewIssuesCache(wf *aw.Workflow) *IssuesCache {
	return &IssuesCache{
		wf: wf,
	}
}

func (cache *IssuesCache) getCacheKey(owner string, repo string) string {
	return fmt.Sprintf("issues-%s-%s", owner, repo)
}

// GetCache return the issue cache.
func (cache *IssuesCache) GetCache(owner string, repo string) ([]model.Issue, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	store := cache.wf.Cache

	issues := []model.Issue{}
	if !store.Exists(cacheKey) {
		return []model.Issue{}, nil
	}

	if store.Expired(cacheKey, maxCacheAge) {
		return []model.Issue{}, nil
	}

	if err := store.LoadJSON(cacheKey, &issues); err != nil {
		cache.wf.FatalError(err)
		return []model.Issue{}, err
	}

	return issues, nil
}

// Store stores the issues to the cache.
func (cache *IssuesCache) Store(owner string, repo string, issues []model.Issue) ([]model.Issue, error) {
	cacheKey := cache.getCacheKey(owner, repo)
	store := cache.wf.Cache

	if err := store.StoreJSON(cacheKey, issues); err != nil {
		cache.wf.FatalError(err)
		return []model.Issue{}, err
	}

	return issues, nil
}
