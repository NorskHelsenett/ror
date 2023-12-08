package clients

import "github.com/dotse/go-health"

type CommonClient interface {
	Ping() bool
	CheckHealth() []health.Check
}
