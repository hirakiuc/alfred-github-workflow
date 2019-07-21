package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// PullsCache describe an instance of pull request cache.
type PullsCache struct {
	*BaseCache
}

// NewPullsCache return an instance of pull request cache store.
func NewPullsCache(wf *aw.Workflow) *PullsCache {
	return &PullsCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *PullsCache) getCacheKey(owner string, repo string) string {
	return fmt.Sprintf("pulls-%s-%s", owner, repo)
}

// GetCache return the pull request cache.
func (cache *PullsCache) GetCache(owner string, repo string) ([]model.PullRequest, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	pulls := []model.PullRequest{}
	err := cache.getRawCache(cacheKey, getMaxCacheAge(), &pulls)
	if err != nil {
		return []model.PullRequest{}, err
	}
	if pulls == nil {
		return []model.PullRequest{}, nil
	}

	return pulls, nil
}

// Store stores the pull requests to the cache.
func (cache *PullsCache) Store(owner string, repo string, pulls []model.PullRequest) ([]model.PullRequest, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	_, err := cache.storeRawData(cacheKey, pulls)
	if err != nil {
		return pulls, err
	}

	return pulls, nil
}
