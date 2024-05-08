package rortypes

// K8s applicationProject struct used with ArgoCD
type ResourceAppProject struct {
	CommonResource `json:",inline"`
	Spec           ResourceAppProjectSpec `json:"spec"`
}
type ResourceAppProjectSpec struct {
	Description  string                               `json:"description"`
	SourceRepos  []string                             `json:"sourceRepos"`
	Destinations []ResourceApplicationSpecDestination `json:"destinations"`
}
