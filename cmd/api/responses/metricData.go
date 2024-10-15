package responses

import (
	models "github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type MetricData struct {
	Total    models.Metrics `json:"total"`
	Filtered models.Metrics `json:"filtered"`
}
