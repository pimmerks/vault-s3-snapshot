package auth

// VaultAuthenticator is the main interface for writing snapshots to varius places.
type VaultAuthenticator interface {
	// GetType returns the type of authenticator this is. Purely for logging.
	GetType() string

	// GetToken returns a token gathered by the specific authentication
	GetToken() string
}
