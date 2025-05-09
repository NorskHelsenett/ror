package rortypes

type ResourceWorkspace struct {
	Spec   ResourceWorkspaceSpec   `json:"spec"`
	Status ResourceWorkspaceStatus `json:"status"`
}

type ResourceWorkspaceSpec struct {
	KubernetesClusters      []ResourceKubernetesCluster     `json:"kubernetesClusters"`
	AvailableMachineClasses []ResourceWorkspaceMachineClass `json:"availableMachineClasses"`
	DefaultMachineClass     ResourceWorkspaceMachineClass   `json:"defaultMachineClass"`
	AvailableStorageClasses []ResourceWorkspaceStorageClass `json:"availableStorageClasses"`
	DefaultStorageClass     ResourceWorkspaceStorageClass   `json:"defaultStorageClass"`
}

type ResourceWorkspaceStatus struct {
	DatacenterId            string                          `json:"datacenterId,omitempty"`
	KubernetesClusters      []ResourceKubernetesCluster     `json:"kubernetesClusters"`
	AvailableMachineClasses []ResourceWorkspaceMachineClass `json:"availableMachineClasses"`
	DefaultMachineClass     ResourceWorkspaceMachineClass   `json:"defaultMachineClass"`
	AvailableStorageClasses []ResourceWorkspaceStorageClass `json:"availableStorageClasses"`
	DefaultStorageClass     ResourceWorkspaceStorageClass   `json:"defaultStorageClass"`
}

type ResourceWorkspaceMachineClass struct {
	Name string `json:"name"`
}

type ResourceWorkspaceStorageClass struct {
	Name string `json:"name"`
}
