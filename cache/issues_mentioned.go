package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

type IssuesMentionedCache struct {
	*BaseCache
}

func NewIssuesMentionedCache(wf *aw.Workflow) *IssuesMentionedCache {
	return &IssuesMentionedCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *IssuesMentionedCache) getCacheKey(user string) string {
	return fmt.Sprintf("issues-mentioned-%s", user)
}

func (cache *IssuesMentionedCache) GetCache(user string) ([]model.Issue, error) {
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

func (cache *IssuesMentionedCache) Store(user string, issues []model.Issue) ([]model.Issue, error) {
	cacheKey := cache.getCacheKey(user)

	_, err := cache.storeRawData(cacheKey, issues)
	if err != nil {
		return issues, err
	}

	return issues, nil
}
