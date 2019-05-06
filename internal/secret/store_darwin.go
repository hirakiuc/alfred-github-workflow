// +build darwin

package secret

import (
	"os/user"

	"github.com/keybase/go-keychain"
)

func queryItem(key string) (keychain.Item, error) {
	query := keychain.NewItem()
	query.SetSecClass(keychain.SecClassGenericPassword)
	query.SetService(ServiceName)

	user, err := user.Current()
	if err != nil {
		return query, err
	}
	query.SetAccount(user.Username)
	query.SetLabel(ItemLabel)
	query.SetAccessGroup(ItemAccessGroup)
	query.SetMatchLimit(keychain.MatchLimitOne)
	query.SetReturnAttributes(true)
	query.SetReturnData(true)

	return query, nil
}

func findItem(query keychain.Item) (*keychain.QueryResult, error) {
	results, err := keychain.QueryItem(query)
	if err != nil {
		return nil, err
	}

	if len(results) == 0 {
		return nil, nil
	}

	return &results[0], nil
}

func newItem(key string) (keychain.Item, error) {
	item := keychain.NewItem()
	item.SetSecClass(keychain.SecClassGenericPassword)
	item.SetService(ServiceName)

	user, err := user.Current()
	if err != nil {
		return item, err
	}

	item.SetAccount(user.Username)
	item.SetLabel(ItemLabel)
	item.SetAccessGroup(ItemAccessGroup)

	item.SetSynchronizable(keychain.SynchronizableNo)
	item.SetAccessible(keychain.AccessibleWhenUnlocked)

	return item, nil
}

func newItemWithSecret(key string, secret string) (keychain.Item, error) {
	item, err := newItem(key)
	if err != nil {
		return item, err
	}

	item.SetData([]byte(secret))
	return item, nil
}

// Store stores the secret in the secret store.
func (store Store) Store(key string, secret string) error {
	query, err := queryItem(key)
	if err != nil {
		return err
	}

	found, err := findItem(query)
	if err != nil {
		return err
	}

	if found != nil {
		item, err := newItemWithSecret(key, secret)
		if err != nil {
			return err
		}

		return keychain.UpdateItem(query, item)
	}

	item, err := newItemWithSecret(key, secret)
	if err != nil {
		return err
	}

	return keychain.AddItem(item)
}

// Get fetch the value from the secret store.
func (store Store) Get(key string) (string, error) {
	query, err := queryItem(key)
	if err != nil {
		return "", err
	}

	found, err := findItem(query)
	if err != nil {
		return "", err
	}

	if found == nil {
		return "", nil
	}

	return string(found.Data), nil
}

// Delete removes the value from the secret store.
func (store Store) Delete(key string) error {
	item, err := newItem(key)
	if err != nil {
		return err
	}

	return keychain.DeleteItem(item)
}
