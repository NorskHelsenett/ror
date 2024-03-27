package apicontracts

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
)

type MetricMetadata struct {
	Type      string `json:"type" bson:"type"`
	ClusterId string `json:"clusterId"  bson:"clusterId"`
	Name      string `json:"name" bson:"name"`
}

type PersistentVolumeClaimMetric struct {
	Metadata            MetricMetadata `json:"metadata" bson:"metadata"`
	Timestamp           time.Time      `json:"timestamp" bson:"timestamp"`
	RequestedAllocation string         `json:"requestedAllocation" bson:"requestedAllocation"`
}

type PodMetricsList struct {
	Kind       string               `json:"kind"`
	APIVersion string               `json:"apiVersion"`
	Items      []PodMetricsListItem `json:"items"`
}

type PodMetricsListItem struct {
	Metadata struct {
		Name      string `json:"name"`
		Namespace string `json:"namespace"`
	} `json:"metadata"`
	Timestamp  time.Time `json:"timestamp"`
	Window     string    `json:"window"`
	Containers []struct {
		Name  string `json:"name"`
		Usage struct {
			CPU    string `json:"cpu"`
			Memory string `json:"memory"`
		} `json:"usage"`
	} `json:"containers"`
}

type NodeMetricsList struct {
	Kind       string                `json:"kind"`
	APIVersion string                `json:"apiVersion"`
	Items      []NodeMetricsListItem `json:"items"`
}

type NodeMetricsListItem struct {
	Metadata struct {
		Name string `json:"name"`
	} `json:"metadata"`
	Timestamp time.Time `json:"timestamp"`
	Window    string    `json:"window"`
	Usage     struct {
		CPU    string `json:"cpu"`
		Memory string `json:"memory"`
	} `json:"usage"`
}

type PodMetric struct {
	Name        string    `json:"name"`
	Namespace   string    `json:"namespace"`
	TimeStamp   time.Time `json:"time"`
	CpuUsage    int64     `json:"cpu"`
	MemoryUsage int64     `json:"memory"`
}

type NodeMetric struct {
	Name             string    `json:"name"`
	TimeStamp        time.Time `json:"time"`
	CpuUsage         int64     `json:"cpu"`
	CpuAllocated     int64     `json:"cpuallocated,omitempty"`
	CpuPercentage    float64   `json:"cpupercentage,omitempty"`
	MemoryUsage      int64     `json:"memory"`
	MemoryAllocated  int64     `json:"memoryallocated,omitempty"`
	MemoryPercentage float64   `json:"memorypercentage,omitempty"`
}

type MetricsReport struct {
	Owner apiresourcecontracts.ResourceOwnerReference `json:"owner"`
	Nodes []NodeMetric                                `json:"nodes"`
}

type MetricsResult struct {
	Key                 MetricsResultKey `json:"key" bson:"_id"`
	AvgCpu              float64          `json:"avgCpu"`
	AvgCpuAllocated     float64          `json:"avgAllocatedCpu,omitempty" bson:"avgAllocatedCpu,truncate"`
	AvgCpuPercentage    float64          `json:"avgPercentageCpu,omitempty" bson:"avgPercentageCpu,truncate"`
	AvgMemory           float64          `json:"avgMemory"`
	AvgMemoryAllocated  float64          `json:"avgAllocatedMemory,omitempty" bson:"avgAllocatedMemory,truncate"`
	AvgMemoryPercentage float64          `json:"avgPercentageMemory,omitempty" bson:"avgPercentageMemory,truncate"`
}

type MetricsResultKey struct {
	Date      MetricsResultKeyTime `json:"date,omitempty"`
	Name      string               `json:"name,omitempty"`
	Namespace string               `json:"namespace,omitempty"`
	ClusterId string               `json:"clusterId,omitempty"`
}
type MetricsResultKeyTime struct {
	Year   int `json:"year,omitempty"`
	Month  int `json:"month,omitempty"`
	Day    int `json:"day,omitempty"`
	Hour   int `json:"hour"`
	Minute int `json:"minute"`
}

type MetricsFilter struct {
	// Type must be set to either node or pod
	Type     string `json:"type"`
	Metadata struct {
		ClusterId string `json:"clusterId,omitempty"`
		Name      string `json:"name,omitempty"`
		Namespace string `json:"namespace,omitempty"`
	} `json:"metadata"`
	Time struct {
		Resolution TimeResolution `json:"resolution,omitempty"`
		Start      time.Time      `json:"start,omitempty"`
		Stop       time.Time      `json:"stop,omitempty"`
	} `json:"time"`
	GroupBy struct {
		Name      bool `json:"name"`
		Namespace bool `json:"namespace"`
		Cluster   bool `json:"cluster"`
	} `json:"groupby"`
}

type TimeResolution int

const (
	ResolutionUnknown TimeResolution = iota
	ResolutionYear
	ResolutionMonth
	ResolutionDay
	ResolutionHour
	ResolutionMinute
)
