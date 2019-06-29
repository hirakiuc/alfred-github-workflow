package cache

import (
	"fmt"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/internal/model"
)

// PullsCache describe an instance of pull request cache.
type PullsCache struct {
	wf *aw.Workflow
}

// NewPullsCache return an instance of pull request cache store.
func NewPullsCache(wf *aw.Workflow) *PullsCache {
	return &PullsCache{
		wf: wf,
	}
}

func (cache *PullsCache) getCacheKey(owner string, repo string) string {
	return fmt.Sprintf("pulls-%s-%s", owner, repo)
}

// GetCache return the pull request cache.
func (cache *PullsCache) GetCache(owner string, repo string) ([]model.PullRequest, error) {
	cacheKey := cache.getCacheKey(owner, repo)

	store := cache.wf.Cache

	pulls := []model.PullRequest{}
	if !store.Exists(cacheKey) {
		return []model.PullRequest{}, nil
	}

	if store.Expired(cacheKey, getMaxCacheAge()) {
		return []model.PullRequest{}, nil
	}

	if err := store.LoadJSON(cacheKey, &pulls); err != nil {
		cache.wf.FatalError(err)
		return []model.PullRequest{}, err
	}

	return pulls, nil
}

// Store stores the pull requests to the cache.
func (cache *PullsCache) Store(owner string, repo string, pulls []model.PullRequest) ([]model.PullRequest, error) {
	cacheKey := cache.getCacheKey(owner, repo)
	store := cache.wf.Cache

	if err := store.StoreJSON(cacheKey, pulls); err != nil {
		cache.wf.FatalError(err)
		return []model.PullRequest{}, err
	}

	return pulls, nil
}
