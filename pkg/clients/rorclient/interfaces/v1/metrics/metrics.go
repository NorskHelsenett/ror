package metrics

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type MetricsInterface interface {
	PostReport(ctx context.Context, metricsReport apicontracts.MetricsReport) error
	CreatePVC(input apicontracts.PersistentVolumeClaimMetric) error
}
