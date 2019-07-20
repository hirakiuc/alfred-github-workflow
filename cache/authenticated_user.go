package cache

import (
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// AuthenticatedUserCache describe an instance of User cache.
type AuthenticatedUserCache struct {
	wf *aw.Workflow
}

// NewAuthenticatedUserCache return an instance of Authenticated User cache store.
func NewAuthenticatedUserCache(wf *aw.Workflow) *AuthenticatedUserCache {
	return &AuthenticatedUserCache{
		wf: wf,
	}
}

func (cache *AuthenticatedUserCache) getCacheKey() string {
	return "authenticated-user"
}

// GetCache return the authenticated user cache.
func (cache *AuthenticatedUserCache) GetCache() (*model.User, error) {
	cacheKey := cache.getCacheKey()

	store := cache.wf.Cache
	if !store.Exists(cacheKey) {
		return nil, nil
	}

	if store.Expired(cacheKey, getMaxCacheAge()) {
		return nil, nil
	}

	var user model.User
	if err := store.LoadJSON(cacheKey, &user); err != nil {
		cache.wf.FatalError(err)
		return nil, err
	}

	return &user, nil
}

// Store stores the authenticated user to the cache.
func (cache *AuthenticatedUserCache) Store(user *model.User) (*model.User, error) {
	cacheKey := cache.getCacheKey()
	store := cache.wf.Cache

	if err := store.StoreJSON(cacheKey, *user); err != nil {
		cache.wf.FatalError(err)
		return user, err
	}

	return user, nil
}
