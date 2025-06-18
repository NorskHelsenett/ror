package apiresourcecontracts

// K8s certificate struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceCertificate struct {
	ApiVersion string                    `json:"apiVersion"`
	Kind       string                    `json:"kind"`
	Metadata   ResourceMetadata          `json:"metadata"`
	Spec       ResourceCertificateSpec   `json:"spec"`
	Status     ResourceCertificateStatus `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceCertificateSpec struct {
	DnsNames   []string                         `json:"dnsNames"`
	SecretName string                           `json:"secretName"`
	IssuerRef  ResourceCertificateSpecIssuerref `json:"issuerRef"`
	Usages     []string                         `json:"usages,omitempty"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceCertificateSpecIssuerref struct {
	Group string `json:"group"`
	Kind  string `json:"kind"`
	Name  string `json:"name"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceCertificateStatus struct {
	NotBefore   string                               `json:"notBefore"`
	NotAfter    string                               `json:"notAfter"`
	RenewalTime string                               `json:"renewalTime"`
	Conditions  []ResourceCertificateStatusCondition `json:"conditions"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceCertificateStatusCondition struct {
	LastTransitionTime string `json:"lastTransitionTime"`
	ObservedGeneration int    `json:"observedGeneration"`
	Message            string `json:"message"`
	Reason             string `json:"reason"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}
