package clusterhelper

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

func SetStatus(cluster *apicontracts.Cluster) {
	cluster.HealthStatus.Health = apicontracts.HealthUnknown
	if cluster.LastObserved.IsZero() {
		return
	}

	timeDiff := time.Since(cluster.LastObserved)
	switch {
	case timeDiff.Minutes() <= 6:
		cluster.HealthStatus.Health = apicontracts.HealthHealthy
	case timeDiff.Minutes() > 6 && timeDiff.Minutes() <= 18:
		cluster.HealthStatus.Health = apicontracts.HealthUnhealthy
	case timeDiff.Minutes() > 18:
		cluster.HealthStatus.Health = apicontracts.HealthBad
	default:
		cluster.HealthStatus.Health = apicontracts.HealthUnknown
	}
}
