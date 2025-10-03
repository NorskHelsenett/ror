package rortypes

type ResourceApplicationInstance struct {
	Spec   ResourceApplicationInstanceSpec   `json:"spec"`
	Status ResourceApplicationInstanceStatus `json:"status"`
}

type ResourceApplicationInstanceSpec struct {
	AppProject  string            `json:"appProject"`
	Application string            `json:"application"`
	RepoUrl     string            `json:"repo"`
	Config      map[string]string `json:"config,omitempty"`
}

type ResourceApplicationInstanceStatus struct {
}
