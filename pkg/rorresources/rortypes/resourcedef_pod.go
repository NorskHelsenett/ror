package rortypes

// ResourcePod
// K8s namepace struct
type ResourcePod struct {
	Spec   ResourcePodSpec   `json:"spec,omitempty" bson:"spec,omitempty"`
	Status ResourcePodStatus `json:"status,omitempty" bson:"status,omitempty"`
}

type ResourcePodSpec struct {
	Containers         []ResourcePodSpecContainers `json:"containers,omitempty"`
	ServiceAccountName string                      `json:"serviceAccountName,omitempty"`
	NodeName           string                      `json:"nodeName,omitempty"`
}
type ResourcePodSpecContainers struct {
	Name  string                           `json:"name,omitempty"`
	Image string                           `json:"image,omitempty"`
	Ports []ResourcePodSpecContainersPorts `json:"ports,omitempty"`
}
type ResourcePodSpecContainersPorts struct {
	Name          string `json:"name,omitempty"`
	ContainerPort int    `json:"containerPort,omitempty"`
	Protocol      string `json:"protocol,omitempty"`
}

type ResourcePodStatus struct {
	Message   string `json:"message,omitempty"`
	Phase     string `json:"phase,omitempty"`
	Reason    string `json:"reason,omitempty"`
	StartTime string `json:"startTime,omitempty"`
}
