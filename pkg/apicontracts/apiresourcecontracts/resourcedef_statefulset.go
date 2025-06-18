package apiresourcecontracts

// K8s deployment struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceStatefulSet struct {
	ApiVersion string                    `json:"apiVersion"`
	Kind       string                    `json:"kind"`
	Metadata   ResourceMetadata          `json:"metadata"`
	Status     ResourceStatefulSetStatus `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceStatefulSetStatus struct {
	Replicas          int `json:"replicas"`
	AvailableReplicas int `json:"availableReplicas"`
	CurrentReplicas   int `json:"currentReplicas"`
	ReadyReplicas     int `json:"readyReplicas"`
	UpdatedReplicas   int `json:"updatedReplicas"`
}
