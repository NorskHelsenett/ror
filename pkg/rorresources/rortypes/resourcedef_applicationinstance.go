package rortypes

type ResourceApplicationInstance struct {
	Spec   ResourceApplicationInstanceSpec   `json:"spec"`
	Status ResourceApplicationInstanceStatus `json:"status"`
}

type ResourceApplicationInstanceSpec struct {
	application string            `json:"application"`
	config      map[string]string `json:"config,omitempty"`
}

type ResourceApplicationInstanceStatus struct {
}
