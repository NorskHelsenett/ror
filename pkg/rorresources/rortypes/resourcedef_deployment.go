package rortypes

// K8s deployment struct
type ResourceDeployment struct {
	Status ResourceDeploymentStatus `json:"status"`
}
type ResourceDeploymentStatus struct {
	Replicas          int `json:"replicas"`
	AvailableReplicas int `json:"availableReplicas"`
	ReadyReplicas     int `json:"readyReplicas"`
	UpdatedReplicas   int `json:"updatedReplicas"`
}
