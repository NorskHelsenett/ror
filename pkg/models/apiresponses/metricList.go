package apiresponses

import (
	models "github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type MetricList struct {
	Items []MetricItem `json:"items"`
}

type MetricItem struct {
	Id      string         `json:"id"`
	Metrics models.Metrics `json:"metrics"`
}

type Metric struct {
	Id               string `json:"id"`
	PriceMonth       int64  `json:"priceMonth"`
	PriceYear        int64  `json:"priceYear"`
	Cpu              int64  `json:"cpu"`
	Memory           int64  `json:"memory"`
	CpuConsumed      int64  `json:"cpuConsumed"`
	MemoryConsumed   int64  `json:"memoryConsumed"`
	CpuPercentage    int64  `json:"cpuPercentage"`
	MemoryPercentage int64  `json:"memoryPercentage"`
	NodePoolCount    int64  `json:"nodePoolCount"`
	NodeCount        int64  `json:"nodeCount"`
	ClusterCount     int64  `json:"clusterCount"`
}
