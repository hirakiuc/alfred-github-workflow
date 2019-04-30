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

// NewStore return an instance of Store.
func NewStore() Store {
	return Store{}
}
