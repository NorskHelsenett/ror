package ssemodels

import (
	"time"
)

type SseType string

const (
	SseType_Unknown              SseType = "unknown"
	SseType_Time                 SseType = "time"
	SseType_Cluster_Created      SseType = "cluster.created"
	SseType_ClusterOrder_Updated SseType = "clusterOrder.updated"
)

type SSEBase struct {
	Event SseType `json:"event"`
}

type Time struct {
	SSEBase
	CurrentTime time.Time `json:"currentTime"`
}

type SseMessage struct {
	SSEBase
	Data interface{} `json:"data"`
}
