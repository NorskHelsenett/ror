package apiresourcecontracts

// K8s namepace struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceReplicaSet struct {
	ApiVersion string                   `json:"apiVersion"`
	Kind       string                   `json:"kind"`
	Metadata   ResourceMetadata         `json:"metadata"`
	Spec       ResourceReplicaSetSpec   `json:"spec"`
	Status     ResourceReplicaSetStatus `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceReplicaSetSpec struct {
	Replicas int                            `json:"replicas"`
	Selector ResourceReplicaSetSpecSelector `json:"selector"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceReplicaSetSpecSelector struct {
	MatchExpressions []ResourceReplicaSetSpecSelectorMatchExpressions `json:"matchExpressions"`
	MatchLabels      map[string]string                                `json:"matchLabels"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceReplicaSetSpecSelectorMatchExpressions struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceReplicaSetStatus struct {
	AvailableReplicas int `json:"availableReplicas"`
	ReadyReplicas     int `json:"readyReplicas"`
	Replicas          int `json:"replicas"`
}
