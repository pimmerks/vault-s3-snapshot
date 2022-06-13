package enums

// VaultAuthMethod is an enum containing the authentication methods.
type VaultAuthMethod string

const (
	// Token is the token authentication.
	Token VaultAuthMethod = "token"
	// Kubernetes is the kubernetes authentication.
	Kubernetes VaultAuthMethod = "k8s"
	// AppRole is the approle authentication.
	AppRole VaultAuthMethod = "AppRole"
)
