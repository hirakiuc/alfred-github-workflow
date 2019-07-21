package cache

import (
	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

// AuthenticatedUserCache describe an instance of User cache.
type AuthenticatedUserCache struct {
	*BaseCache
}

// NewAuthenticatedUserCache return an instance of Authenticated User cache store.
func NewAuthenticatedUserCache(wf *aw.Workflow) *AuthenticatedUserCache {
	return &AuthenticatedUserCache{
		&BaseCache{
			wf: wf,
		},
	}
}

func (cache *AuthenticatedUserCache) getCacheKey() string {
	return "authenticated-user"
}

// GetCache return the authenticated user cache.
func (cache *AuthenticatedUserCache) GetCache() (*model.User, error) {
	cacheKey := cache.getCacheKey()

	var user model.User
	err := cache.getRawCache(cacheKey, getMaxCacheAge(), &user)
	if err != nil {
		return nil, err
	}
	if !user.IsValid() {
		return nil, err
	}

	return &user, nil
}

// Store stores the authenticated user to the cache.
func (cache *AuthenticatedUserCache) Store(user *model.User) (*model.User, error) {
	cacheKey := cache.getCacheKey()

	_, err := cache.storeRawData(cacheKey, *user)
	if err != nil {
		return user, err
	}

	return user, nil
}
