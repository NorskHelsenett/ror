package rortypes

// ResourcePod
// K8s namepace struct
type ResourcePod struct {
	Spec   ResourcePodSpec   `json:"spec"`
	Status ResourcePodStatus `json:"status"`
}

type ResourcePodSpec struct {
	Containers         []ResourcePodSpecContainers `json:"containers"`
	ServiceAccountName string                      `json:"serviceAccountName"`
	NodeName           string                      `json:"nodeName"`
}
type ResourcePodSpecContainers struct {
	Name  string                           `json:"name"`
	Image string                           `json:"image"`
	Ports []ResourcePodSpecContainersPorts `json:"ports"`
}
type ResourcePodSpecContainersPorts struct {
	Name          string `json:"name"`
	ContainerPort int    `json:"containerPort"`
	Protocol      string `json:"protocol"`
}

type ResourcePodStatus struct {
	Message   string `json:"message,omitempty"`
	Phase     string `json:"phase"`
	Reason    string `json:"reason,omitempty"`
	StartTime string `json:"startTime"`
}
