package rabbitmq

import "os"

type Authenticator interface {
	GetCredentials() (string, string)
}

// authenticator is a struct used for defining a default Authenticator.
type authenticator struct {
}

// GetCredentials sets a default username and password, and overwrites them with
// values from the environment if they exist.
func (a *authenticator) GetCredentials() (string, string) {
	username, password := "guest", "guest"
	u, ok := os.LookupEnv("ROR_RABBITMQ_USERNAME")
	if ok {
		username = u
	}
	p, ok := os.LookupEnv("ROR_RABBITMQ_PASSWORD")
	if ok {
		password = p
	}
	return username, password
}
