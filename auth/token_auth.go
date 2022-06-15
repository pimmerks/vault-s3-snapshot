package auth

import "github.com/pimmerks/vault-s3-snapshot/config"

type TokenAuthentication struct {
	Token string
}

// GetType implements VaultAuthenticator
func (*TokenAuthentication) GetType() string {
	return "TokenAuthentication"
}

// GetToken implements VaultAuthenticator
func (auth *TokenAuthentication) GetToken() string {
	return auth.Token
}

// CreateTokenAuth creates a new VaultAuthenticator that authenticates using the token.
func CreateTokenAuth(config *config.Configuration) VaultAuthenticator {
	auth := &TokenAuthentication{}
	auth.Token = config.Token
	return auth
}
