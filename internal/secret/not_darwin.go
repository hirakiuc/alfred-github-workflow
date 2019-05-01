// +build linux

package secret

// Store stores the secret in the secret store.
func (store Store) Store(key string, secret string) error {
	return nil
}

// Get fetch the value from the secret store.
func (store Store) Get(key string) (string, error) {
	return "", nil
}

// Delete removes the value from the secret store
func (store Store) Delete(key string) error {
	return nil
}
