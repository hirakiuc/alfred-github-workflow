package pulls

import (
	"context"

	aw "github.com/deanishe/awgo"
	"github.com/hirakiuc/alfred-github-workflow/api"
	"github.com/hirakiuc/alfred-github-workflow/cache"
	"github.com/hirakiuc/alfred-github-workflow/model"
)

func fetchUser(ctx context.Context, wf *aw.Workflow, client *api.Client) (*model.User, error) {
	store := cache.NewAuthenticatedUserCache(wf)

	user, err := store.GetCache()
	if err != nil {
		return nil, err
	}
	if user != nil {
		return user, nil
	}

	user, err = client.FetchAuthenticatedUser(ctx)
	if err != nil {
		return nil, err
	}

	return store.Store(user)
}
