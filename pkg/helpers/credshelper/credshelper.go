package credshelper

type CredHelperWithRenew interface {
	CredHelper
	CheckAndRenew() bool // Returns true if credentials were renewed
}

type CredHelper interface {
	GetUsername() string              // Returns the current username
	GetPassword() string              // Returns the current password
	GetCredentials() (string, string) // Returns username and password
}

// ForceRenewer is an optional interface a credential helper may implement to
// force an immediate credential renewal regardless of any expiry schedule. It
// lets clients recover when credentials are revoked server-side before their
// local lease is considered expired.
type ForceRenewer interface {
	ForceRenew() error
}
