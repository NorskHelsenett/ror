package rortypes

// K8s certificate struct
type ResourceCertificate struct {
	Spec   ResourceCertificateSpec   `json:"spec"`
	Status ResourceCertificateStatus `json:"status"`
}

type ResourceCertificateSpec struct {
	DnsNames   []string                         `json:"dnsNames"`
	SecretName string                           `json:"secretName"`
	IssuerRef  ResourceCertificateSpecIssuerref `json:"issuerRef"`
	Usages     []string                         `json:"usages,omitempty"`
}
type ResourceCertificateSpecIssuerref struct {
	Group string `json:"group"`
	Kind  string `json:"kind"`
	Name  string `json:"name"`
}
type ResourceCertificateStatus struct {
	NotBefore   string                               `json:"notBefore"`
	NotAfter    string                               `json:"notAfter"`
	RenewalTime string                               `json:"renewalTime"`
	Conditions  []ResourceCertificateStatusCondition `json:"conditions"`
}

type ResourceCertificateStatusCondition struct {
	LastTransitionTime string `json:"lastTransitionTime"`
	ObservedGeneration int    `json:"observedGeneration"`
	Message            string `json:"message"`
	Reason             string `json:"reason"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}
