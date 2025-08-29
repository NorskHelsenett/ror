package apiresourcecontracts

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceConfiguration struct {
	ApiVersion string                    `json:"apiVersion"`
	Kind       string                    `json:"kind"`
	Metadata   ResourceMetadata          `json:"metadata"`
	Spec       ResourceConfigurationSpec `json:"spec"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceConfigurationSpec struct {
	Type   string `json:"type"`
	B64enc bool   `json:"b64enc"`
	Data   string `json:"data"`
}
