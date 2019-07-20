package api

import (
	"context"
	"errors"

	"github.com/hirakiuc/alfred-github-workflow/model"
)

// GetAuthenticatedUser fetch the authenticated user.
func (client *Client) FetchAuthenticatedUser(ctx context.Context) (*model.User, error) {
	user, _, err := client.github.Users.Get(ctx, "")
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, errors.New("user not found")
	}

	m := model.ConvertUser(user)
	return &m, nil
}
