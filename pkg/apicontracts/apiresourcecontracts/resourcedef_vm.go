package apiresourcecontracts

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachine struct {
	ApiVersion string                       `json:"api_version"`
	Kind       string                       `json:"kind"`
	Metadata   ResourceMetadata             `json:"metadata"`
	Id         string                       `json:"id"`
	Name       string                       `json:"name"`
	Spec       ResourceVirtualMachineSpec   `json:"spec"`
	Status     ResourceVirtualMachineStatus `json:"status"`
}

// Desired state// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachineSpec struct {
	OperativeSystem ResourceVirtualMachineOperativeSystem `json:"guest"`
	Config          ResourceVirtualMachineConfig          `json:"config"`
	Runtime         ResourceVirtualMachineRuntime         `json:"runtime"`
	Tags            []ResourceVirtualMachineTag           `json:"tags"`
}

// Observed state// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachineStatus struct {
	OperativeSystem ResourceVirtualMachineOperativeSystem `json:"guest"`
	Config          ResourceVirtualMachineConfig          `json:"config"`
	Runtime         ResourceVirtualMachineRuntime         `json:"runtime"`
	Tags            []ResourceVirtualMachineTag           `json:"tags"`
}

// The guest operating system running on the vm// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachineOperativeSystem struct {
	Id          string `json:"id"`
	Family      string `json:"family"`
	FullName    string `json:"fullName"`
	HostName    string `json:"hostName"`
	IpV4Address string `json:"ipV4Address"`
	IpV6Address string `json:"ipV6Address"`
	State       string `json:"state"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachineConfig struct {
	Name             string `json:"name"`
	MemorySize       int    `json:"memorySize"`
	CpuCount         int    `json:"cpuCount"`
	VirtualDiskCount int    `json:"virtualDiskCount"`
	Annotation       string `json:"annotation"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachineRuntime struct {
	ConnectionState string `json:"connectionState"`
	PowerState      string `json:"powerState"`
	MaxCpu          int    `json:"maxCpu"`
	MaxMemory       int    `json:"maxMemory"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceVirtualMachineTag struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
