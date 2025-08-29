package apiresourcecontracts

import "k8s.io/apimachinery/pkg/util/intstr"

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceNetworkPolicy struct {
	ApiVersion string                      `json:"apiVersion"`
	Kind       string                      `json:"kind"`
	Metadata   ResourceMetadata            `json:"metadata"`
	Spec       ResourceNetworkPolicySpec   `json:"spec"`
	Status     ResourceNetworkPolicyStatus `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceNetworkPolicySpec struct {
	Egress      []ResourceNetworkPolicyEgressRule  `json:"egress"`
	Ingress     []ResourceNetworkPolicyIngressRule `json:"ingress"`
	PodSelector ResourceNetworkPolicyPodSelector   `json:"podSelector"`
	PolicyTypes []string                           `json:"policyTypes"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceNetworkPolicyPodSelector struct {
	MatchLabels map[string]string `json:"matchLabels"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceNetworkPolicyEgressRule struct {
	Ports []ResourceNetworkPolicyPort `json:"ports"`
	To    []ResourceNetworkPolicyPeer `json:"to"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceNetworkPolicyPeer struct {
	IpBlock           *ResourceNetworkPolicyIpBlock  `json:"ipBlock"`
	NamespaceSelector *ResourceNetworkPolicySelector `json:"namespaceSelector"`
	PodSelector       *ResourceNetworkPolicySelector `json:"podSelector"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceNetworkPolicyIpBlock struct {
	Cidr   string   `json:"cidr"`
	Except []string `json:"except"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceNetworkPolicySelector struct {
	MatchExpressions []ResourceNetworkPolicySelectorExpression `json:"matchExpressions"`
	MatchLabels      map[string]string                         `json:"matchLabels"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceNetworkPolicySelectorExpression struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceNetworkPolicyPort struct {
	Endport  int                `json:"endPort"`
	Port     intstr.IntOrString `json:"port"`
	Protocol string             `json:"protocol"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceNetworkPolicyIngressRule struct {
	From  []ResourceNetworkPolicyPeer `json:"from"`
	Ports []ResourceNetworkPolicyPort `json:"ports"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceNetworkPolicyStatus struct {
	Conditions []ResourceNetworkPolicyCondition `json:"conditions"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceNetworkPolicyCondition struct {
	LastTransitionTime string `json:"lastTransitionTime"`
	Message            string `json:"message"`
	ObservedGeneration int    `json:"observedGeneration"`
	Reason             string `json:"reason"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}
