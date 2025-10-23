package clients

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
)

type CommonClient interface {
	Ping() bool
	PingWithContext(ctx context.Context) bool
	CommonHealthChecker
}

type CommonHealthChecker interface {
	CheckHealthWithoutContext() []rorhealth.Check
	CheckHealth(ctx context.Context) []rorhealth.Check
}
