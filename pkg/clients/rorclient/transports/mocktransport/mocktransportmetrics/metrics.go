package mocktransportmetrics

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type V1Client struct{}

func NewV1Client() *V1Client {
	return &V1Client{}
}

func (c *V1Client) PostReport(ctx context.Context, metricsReport apicontracts.MetricsReport) error {
	// Mock implementation - just return nil to simulate success
	return nil
}

func (c *V1Client) CreatePVC(input apicontracts.PersistentVolumeClaimMetric) error {
	// Mock implementation - just return nil to simulate success
	return nil
}
