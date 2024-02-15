package apiresourcecontracts

// K8s deployment struct
type ResourceDaemonSet struct {
	ApiVersion string                  `json:"apiVersion"`
	Kind       string                  `json:"kind"`
	Metadata   ResourceMetadata        `json:"metadata"`
	Status     ResourceDaemonSetStatus `json:"status"`
}
type ResourceDaemonSetStatus struct {
	NumberReady            int `json:"numberReady"`
	NumberUnavailable      int `json:"numberUnavailable"`
	NumberMisscheduled     int `json:"currentReplicas"`
	NumberAvailable        int `json:"numberAvailable"`
	UpdatedNumberScheduled int `json:"updatedNumberScheduled"`
	DesiredNumberScheduled int `json:"desiredNumberScheduled"`
	CurrentNumberScheduled int `json:"currentNumberScheduled"`
}
