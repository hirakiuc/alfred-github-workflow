package secret

const (
	// ServiceName is the service name.
	ServiceName = "alfred-github-workflow"
	// ItemAccessGroup is the item access group name.
	ItemAccessGroup = "jp.altab.app.alfred.workflow.github"
	// ItemLabel is the item label in the secrets.
	ItemLabel = "API token for alfred-github-workflow"

	// KeyGithubAPIToken is the key string to store the github api token.
	KeyGithubAPIToken = "github-api-token"
)

// Store describe the store to keep secrets.
type Store struct {
}

// Storer describe the interface of the store.
type Storer interface {
	// Store stores the secret in the secret store.
	Store(key string, secret string) error

	// Get fetch the value from the secret store.
	Get(key string) (string, error)

	// Delete removes the value from the secret store
	Delete(key string) error
}

// NewStore return an instance of Store.
func NewStore() Storer {
	return Store{}
}
