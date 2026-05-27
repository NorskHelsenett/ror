package rortypes

type ResourceConfiguration struct {
	Spec ResourceConfigurationSpec `json:"spec"`
}

type ResourceConfigurationSpec struct {
	Type   string `json:"type"`
	B64enc bool   `json:"b64enc"`
	Data   string `json:"data"`
}

// (r ResourceConfiguration) Get returns a pointer to the resource of type ResourceConfiguration
func (r *ResourceConfiguration) Get() *ResourceConfiguration {
	return r
}
