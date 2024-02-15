package apiresourcecontracts

// K8s deployment struct
type ResourceDeployment struct {
	ApiVersion string                   `json:"apiVersion"`
	Kind       string                   `json:"kind"`
	Metadata   ResourceMetadata         `json:"metadata"`
	Status     ResourceDeploymentStatus `json:"status"`
}
type ResourceDeploymentStatus struct {
	Replicas          int `json:"replicas"`
	AvailableReplicas int `json:"availableReplicas"`
	ReadyReplicas     int `json:"readyReplicas"`
	UpdatedReplicas   int `json:"updatedReplicas"`
}
