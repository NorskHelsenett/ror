package rortypes

// K8s deployment struct
type ResourceStatefulSet struct {
	Status ResourceStatefulSetStatus `json:"status"`
}
type ResourceStatefulSetStatus struct {
	Replicas          int `json:"replicas"`
	AvailableReplicas int `json:"availableReplicas"`
	CurrentReplicas   int `json:"currentReplicas"`
	ReadyReplicas     int `json:"readyReplicas"`
	UpdatedReplicas   int `json:"updatedReplicas"`
}
