package apiresourcecontracts

// K8s namepace struct
type ResourceNamespace struct {
	ApiVersion string           `json:"apiVersion"`
	Kind       string           `json:"kind"`
	Metadata   ResourceMetadata `json:"metadata"`
}
