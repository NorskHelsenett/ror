package rortypes

type ResourceConfiguration struct {
	Spec ResourceConfigurationSpec `json:"spec"`
}

type ResourceConfigurationSpec struct {
	Type   string `json:"type"`
	B64enc bool   `json:"b64enc"`
	Data   string `json:"data"`
}
