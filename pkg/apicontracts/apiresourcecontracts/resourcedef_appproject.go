package apiresourcecontracts

// K8s applicationProject struct used with ArgoCD// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceAppProject struct {
	ApiVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Metadata   ResourceMetadata       `json:"metadata"`
	Spec       ResourceAppProjectSpec `json:"spec"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceAppProjectSpec struct {
	Description  string                               `json:"description"`
	SourceRepos  []string                             `json:"sourceRepos"`
	Destinations []ResourceApplicationSpecDestination `json:"destinations"`
}
