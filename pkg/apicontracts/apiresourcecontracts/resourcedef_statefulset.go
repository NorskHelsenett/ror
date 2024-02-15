package apiresourcecontracts

// K8s deployment struct
type ResourceStatefulSet struct {
	ApiVersion string                    `json:"apiVersion"`
	Kind       string                    `json:"kind"`
	Metadata   ResourceMetadata          `json:"metadata"`
	Status     ResourceStatefulSetStatus `json:"status"`
}
type ResourceStatefulSetStatus struct {
	Replicas          int `json:"replicas"`
	AvailableReplicas int `json:"availableReplicas"`
	CurrentReplicas   int `json:"currentReplicas"`
	ReadyReplicas     int `json:"readyReplicas"`
	UpdatedReplicas   int `json:"updatedReplicas"`
}
