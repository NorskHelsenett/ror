package rortypes

import (
	"github.com/NorskHelsenett/ror/pkg/config/globalconfig"
)

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

// (r *ResourceReplicaSet) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceReplicaSet) ApplyInputFilter(cr *CommonResource) error {
	if globalconfig.InternalNamespaces[cr.Metadata.Namespace] {
		cr.RorMeta.Internal = true
	}
	return nil
}

// (r ResourceReplicaSet) Get returns a pointer to the resource of type ResourceReplicaSet
func (r *ResourceReplicaSet) Get() *ResourceReplicaSet {
	return r
}
