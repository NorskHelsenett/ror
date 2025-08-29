package apiresourcecontracts

// K8s namepace struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressClass struct {
	ApiVersion string                   `json:"apiVersion"`
	Kind       string                   `json:"kind"`
	Metadata   ResourceMetadata         `json:"metadata"`
	Spec       ResourceIngressClassSpec `json:"spec"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressClassSpec struct {
	Controller string                             `json:"controller"`
	Parameters ResourceIngressClassSpecParameters `json:"parameters"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceIngressClassSpecParameters struct {
	ApiGroup  string `json:"apiGroup"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Scope     string `json:"scope"`
}
