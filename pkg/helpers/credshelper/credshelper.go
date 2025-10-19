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
