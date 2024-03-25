package metrics

import "github.com/NorskHelsenett/ror/pkg/apicontracts"

type MetricsInterface interface {
	CreatePVC(input apicontracts.PersistentVolumeClaimMetric) error
}
