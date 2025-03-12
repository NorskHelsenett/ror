package rabbitmq

import "os"

type Authenticator interface {
	GetCredentials() (string, string)
}

type authenticator struct {
}

func (a *authenticator) GetCredentials() (string, string) {
	username, password := "guest", "guest"
	u, ok := os.LookupEnv("RABBITMQ_USERNAME")
	if ok {
		username = u
	}
	p, ok := os.LookupEnv("RABBITMQ_PASSWORD")
	if ok {
		password = p
	}
	return username, password
}
