package rortypes

import (
	"k8s.io/apimachinery/pkg/util/intstr"
)

type ResourceNetworkPolicy struct {
	Spec   ResourceNetworkPolicySpec   `json:"spec"`
	Status ResourceNetworkPolicyStatus `json:"status"`
}

type ResourceNetworkPolicySpec struct {
	Egress      []ResourceNetworkPolicyEgressRule  `json:"egress"`
	Ingress     []ResourceNetworkPolicyIngressRule `json:"ingress"`
	PodSelector ResourceNetworkPolicyPodSelector   `json:"podSelector"`
	PolicyTypes []string                           `json:"policyTypes"`
}

type ResourceNetworkPolicyPodSelector struct {
	MatchLabels map[string]string `json:"matchLabels"`
}

type ResourceNetworkPolicyEgressRule struct {
	Ports []ResourceNetworkPolicyPort `json:"ports"`
	To    []ResourceNetworkPolicyPeer `json:"to"`
}

type ResourceNetworkPolicyPeer struct {
	IpBlock           *ResourceNetworkPolicyIpBlock  `json:"ipBlock"`
	NamespaceSelector *ResourceNetworkPolicySelector `json:"namespaceSelector"`
	PodSelector       *ResourceNetworkPolicySelector `json:"podSelector"`
}

type ResourceNetworkPolicyIpBlock struct {
	Cidr   string   `json:"cidr"`
	Except []string `json:"except"`
}

type ResourceNetworkPolicySelector struct {
	MatchExpressions []ResourceNetworkPolicySelectorExpression `json:"matchExpressions"`
	MatchLabels      map[string]string                         `json:"matchLabels"`
}

type ResourceNetworkPolicySelectorExpression struct {
	Key      string   `json:"key"`
	Operator string   `json:"operator"`
	Values   []string `json:"values"`
}

type ResourceNetworkPolicyPort struct {
	Endport  int                `json:"endPort"`
	Port     intstr.IntOrString `json:"port"`
	Protocol string             `json:"protocol"`
}

type ResourceNetworkPolicyIngressRule struct {
	From  []ResourceNetworkPolicyPeer `json:"from"`
	Ports []ResourceNetworkPolicyPort `json:"ports"`
}

type ResourceNetworkPolicyStatus struct {
	Conditions []ResourceNetworkPolicyCondition `json:"conditions"`
}

type ResourceNetworkPolicyCondition struct {
	LastTransitionTime string `json:"lastTransitionTime"`
	Message            string `json:"message"`
	ObservedGeneration int    `json:"observedGeneration"`
	Reason             string `json:"reason"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}
