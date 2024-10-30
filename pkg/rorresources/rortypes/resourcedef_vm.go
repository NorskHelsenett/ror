package rortypes

type ResourceVirtualMachine struct {
	Id     string                       `json:"id"`
	Name   string                       `json:"name"`
	Spec   ResourceVirtualMachineSpec   `json:"spec"`
	Status ResourceVirtualMachineStatus `json:"status"`
}

// Desired state
type ResourceVirtualMachineSpec struct {
	OperativeSystem ResourceVirtualMachineOperativeSystem `json:"guest"`
	Config          ResourceVirtualMachineConfig          `json:"config"`
	Runtime         ResourceVirtualMachineRuntime         `json:"runtime"`
	Tags            []ResourceVirtualMachineTag           `json:"tags"`
}

// Observed state
type ResourceVirtualMachineStatus struct {
	OperativeSystem ResourceVirtualMachineOperativeSystem `json:"guest"`
	Config          ResourceVirtualMachineConfig          `json:"config"`
	Runtime         ResourceVirtualMachineRuntime         `json:"runtime"`
	Tags            []ResourceVirtualMachineTag           `json:"tags"`
}

// The guest operating system running on the vm
type ResourceVirtualMachineOperativeSystem struct {
	Id          string `json:"id"`
	Family      string `json:"family"`
	FullName    string `json:"fullName"`
	HostName    string `json:"hostName"`
	IpV4Address string `json:"ipV4Address"`
	IpV6Address string `json:"ipV6Address"`
	State       string `json:"state"`
}

type ResourceVirtualMachineConfig struct {
	Name             string `json:"name"`
	MemorySize       int    `json:"memorySize"`
	CpuCount         int    `json:"cpuCount"`
	VirtualDiskCount int    `json:"virtualDiskCount"`
}

type ResourceVirtualMachineRuntime struct {
	ConnectionState string `json:"connectionState"`
	PowerState      string `json:"powerState"`
	MaxCpu          int    `json:"maxCpu"`
	MaxMemory       int    `json:"maxMemory"`
}

type ResourceVirtualMachineTag struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
