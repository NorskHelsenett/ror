package apiresourcecontracts

// ResourcePod
// K8s namepace struct// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourcePod struct {
	ApiVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   ResourceMetadata  `json:"metadata"`
	Spec       ResourcePodSpec   `json:"spec"`
	Status     ResourcePodStatus `json:"status"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourcePodSpec struct {
	Containers         []ResourcePodSpecContainers `json:"containers"`
	ServiceAccountName string                      `json:"serviceAccountName"`
	NodeName           string                      `json:"nodeName"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourcePodSpecContainers struct {
	Name  string                           `json:"name"`
	Image string                           `json:"image"`
	Ports []ResourcePodSpecContainersPorts `json:"ports"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourcePodSpecContainersPorts struct {
	Name          string `json:"name"`
	ContainerPort int    `json:"containerPort"`
	Protocol      string `json:"protocol"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourcePodStatus struct {
	Message   string `json:"message,omitempty"`
	Phase     string `json:"phase"`
	Reason    string `json:"reason,omitempty"`
	StartTime string `json:"startTime"`
}
