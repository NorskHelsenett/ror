package rortypes

// K8s namepace struct
type ResourceReplicaSet struct {
	Spec   ResourceReplicaSetSpec   `json:"spec"`
	Status ResourceReplicaSetStatus `json:"status"`
}

type ResourceReplicaSetSpec struct {
	Replicas int                            `json:"replicas"`
	Selector ResourceReplicaSetSpecSelector `json:"selector"`
}

type ResourceReplicaSetSpecSelector struct {
	MatchExpressions []ResourceReplicaSetSpecSelectorMatchExpressions `json:"matchExpressions"`
	MatchLabels      map[string]string                                `json:"matchLabels"`
}
type ResourceReplicaSetSpecSelectorMatchExpressions struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type ResourceReplicaSetStatus struct {
	AvailableReplicas int `json:"availableReplicas"`
	ReadyReplicas     int `json:"readyReplicas"`
	Replicas          int `json:"replicas"`
}
