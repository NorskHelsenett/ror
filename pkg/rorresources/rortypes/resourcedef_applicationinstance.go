package rortypes

type resourceApplicationInstance struct {
	Spec   resourceApplicationInstanceSpec   `json:"spec"`
	Status resourceApplicationInstanceStatus `json:"status"`
}

type resourceApplicationInstanceSpec struct {
	hostUid     string            `json:"hostUid"`
	clusterName string            `json:"clusterName"`
	application string            `json:"application"`
	config      map[string]string `json:"config,omitempty"`
}

type resourceApplicationInstanceStatus struct {
}
