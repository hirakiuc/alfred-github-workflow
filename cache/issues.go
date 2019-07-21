package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// IssuesCache describe an instance of issue cache.
type IssuesCache struct {
	*BaseCache
}

// NewIssuesCache return an instance of issue cache store.
func NewIssuesCache(wf *aw.Workflow) *IssuesCache {
	return &IssuesCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *IssuesCache) getCacheKey(owner string, repo string) string {
	return fmt.Sprintf("issues-%s-%s", owner, repo)
}

// GetCache return the issue cache.
func (cache *IssuesCache) GetCache(owner string, repo string) ([]model.Issue, error) {
	cacheKey := cache.getCacheKey(owner, repo)

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

// Store stores the issues to the cache.
func (cache *IssuesCache) Store(owner string, repo string, issues []model.Issue) ([]model.Issue, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	_, err := cache.storeRawData(cacheKey, issues)
	if err != nil {
		return issues, err
	}

	return issues, nil
}
