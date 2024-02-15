package apiresourcecontracts

// K8s namepace struct
type ResourceIngressClass struct {
	ApiVersion string                   `json:"apiVersion"`
	Kind       string                   `json:"kind"`
	Metadata   ResourceMetadata         `json:"metadata"`
	Spec       ResourceIngressClassSpec `json:"spec"`
}

type ResourceIngressClassSpec struct {
	Controller string                             `json:"controller"`
	Parameters ResourceIngressClassSpecParameters `json:"parameters"`
}

type ResourceIngressClassSpecParameters struct {
	ApiGroup  string `json:"apiGroup"`
	Kind      string `json:"kind"`
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
	Scope     string `json:"scope"`
}
