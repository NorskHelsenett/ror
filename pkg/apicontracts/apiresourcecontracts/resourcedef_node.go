package apiresourcecontracts

// K8s node struct
type ResourceNode struct {
	ApiVersion string             `json:"apiVersion"`
	Kind       string             `json:"kind"`
	Metadata   ResourceMetadata   `json:"metadata"`
	Spec       ResourceNodeSpec   `json:"spec"`
	Status     ResourceNodeStatus `json:"status"`
}

type ResourceNodeSpec struct {
	PodCIDR    string                   `json:"podCIDR,omitempty"`
	PodCIDRs   []string                 `json:"podCIDRs,omitempty"`
	ProviderID string                   `json:"providerID,omitempty"`
	Taints     []ResourceNodeSpecTaints `json:"taints,omitempty"`
}

type ResourceNodeSpecTaints struct {
	Effect string `json:"effect"`
	Key    string `json:"key"`
}

type ResourceNodeStatus struct {
	Addresses  []ResourceNodeStatusAddresses  `json:"addresses"`
	Capacity   ResourceNodeStatusCapacity     `json:"capacity"`
	Conditions []ResourceNodeStatusConditions `json:"conditions"`
	NodeInfo   ResourceNodeStatusNodeinfo     `json:"nodeInfo"`
}

type ResourceNodeStatusAddresses struct {
	Address string `json:"address"`
	Type    string `json:"type"`
}
type ResourceNodeStatusCapacity struct {
	Cpu              string `json:"cpu"`
	EphemeralStorage string `json:"ephemeral-storage"`
	Memory           string `json:"memory"`
	Pods             string `json:"pods"`
}
type ResourceNodeStatusConditions struct {
	LastHeartbeatTime  string `json:"lastHeartbeatTime"`
	LastTransitionTime string `json:"lastTransitionTime"`
	Message            string `json:"message"`
	Reason             string `json:"reason"`
	Status             string `json:"status"`
	Type               string `json:"type"`
}

type ResourceNodeStatusNodeinfo struct {
	Architecture            string `json:"architecture"`
	BootID                  string `json:"bootID"`
	ContainerRuntimeVersion string `json:"containerRuntimeVersion"`
	KernelVersion           string `json:"kernelVersion"`
	KubeProxyVersion        string `json:"kubeProxyVersion"`
	KubeletVersion          string `json:"kubeletVersion"`
	MachineID               string `json:"machineID"`
	OperatingSystem         string `json:"operatingSystem"`
	OsImage                 string `json:"osImage"`
	SystemUUID              string `json:"systemUUID"`
}
