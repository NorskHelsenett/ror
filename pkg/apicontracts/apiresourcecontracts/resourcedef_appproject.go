package apiresourcecontracts

// K8s applicationProject struct used with ArgoCD
type ResourceAppProject struct {
	ApiVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Metadata   ResourceMetadata       `json:"metadata"`
	Spec       ResourceAppProjectSpec `json:"spec"`
}
type ResourceAppProjectSpec struct {
	Description  string                               `json:"description"`
	SourceRepos  []string                             `json:"sourceRepos"`
	Destinations []ResourceApplicationSpecDestination `json:"destinations"`
}
