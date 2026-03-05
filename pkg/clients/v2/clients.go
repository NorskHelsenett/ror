package clients

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
)

type CommonClient interface {
	Ping(ctx context.Context) bool
	CommonHealthChecker
}

type CommonHealthChecker interface {
	CheckHealth(ctx context.Context) []rorhealth.Check
}
