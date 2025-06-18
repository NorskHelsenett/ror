package apiresourcecontracts

// K8s deployment struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceDaemonSet struct {
	ApiVersion string                  `json:"apiVersion"`
	Kind       string                  `json:"kind"`
	Metadata   ResourceMetadata        `json:"metadata"`
	Status     ResourceDaemonSetStatus `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceDaemonSetStatus struct {
	NumberReady            int `json:"numberReady"`
	NumberUnavailable      int `json:"numberUnavailable"`
	NumberMisscheduled     int `json:"currentReplicas"`
	NumberAvailable        int `json:"numberAvailable"`
	UpdatedNumberScheduled int `json:"updatedNumberScheduled"`
	DesiredNumberScheduled int `json:"desiredNumberScheduled"`
	CurrentNumberScheduled int `json:"currentNumberScheduled"`
}
