package metrics

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type MetricsInterface interface {
	PostReport(ctx context.Context, metricsReport apicontracts.MetricsReport) error
	CreatePVC(ctx context.Context, input apicontracts.PersistentVolumeClaimMetric) error
}
