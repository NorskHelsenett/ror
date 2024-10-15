package responses

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type ClusterPaginatedResult struct {
	Data       []apicontracts.Cluster `json:"data"`
	DataCount  int64                  `json:"dataCount"`
	TotalCount int64                  `json:"totalCount"`
	Offset     int64                  `json:"offset"`
}

type MetricsPaginatedResult struct {
	Data       []apicontracts.Metric `json:"data"`
	DataCount  int64                 `json:"dataCount"`
	TotalCount int64                 `json:"totalCount"`
	Offset     int64                 `json:"offset"`
}
