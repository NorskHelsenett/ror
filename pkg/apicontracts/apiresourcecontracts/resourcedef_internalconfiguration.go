package apiresourcecontracts

type ResourceConfiguration struct {
	ApiVersion string                    `json:"apiVersion"`
	Kind       string                    `json:"kind"`
	Metadata   ResourceMetadata          `json:"metadata"`
	Spec       ResourceConfigurationSpec `json:"spec"`
}

type ResourceConfigurationSpec struct {
	Type   string `json:"type"`
	B64enc bool   `json:"b64enc"`
	Data   string `json:"data"`
}
