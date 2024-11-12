package rortypes

type ResourceVirtualMachine struct {
	Id     string                       `json:"id"`
	Name   string                       `json:"name"`
	Spec   ResourceVirtualMachineSpec   `json:"spec"`
	Status ResourceVirtualMachineStatus `json:"status"`
}

// Desired state
type ResourceVirtualMachineSpec struct {
	Cpu             ResourceVirtualMachineCpuSpec             `json:"cpu"`
	Tags            []ResourceVirtualMachineTagSpec           `json:"tags"`
	Disks           []ResourceVirtualMachineDiskSpec          `json:"disks"`
	Memory          ResourceVirtualMachineMemorySpec          `json:"memory"`
	Networks        []ResourceVirtualMachineNetworkSpec       `json:"networks"`
	OperatingSystem ResourceVirtualMachineOperatingSystemSpec `json:"operatingSystem"`
}

// Observed state
type ResourceVirtualMachineStatus struct {
	Cpu             ResourceVirtualMachineCpuStatus             `json:"cpu"`
	Disks           []ResourceVirtualMachineDiskStatus          `json:"disks"`
	Memory          ResourceVirtualMachineMemoryStatus          `json:"memory"`
	Networks        []ResourceVirtualMachineNetworkStatus       `json:"networks"`
	OperatingSystem ResourceVirtualMachineOperatingSystemStatus `json:"operatingSystem"`
}

type ResourceVirtualMachineDiskSpec struct {
	Id   string `json:"id"`
	Size int    `json:"size"`
	Name string `json:"name"`
	Type string `json:"type"`
}

type ResourceVirtualMachineDiskStatus struct {
	Id    string `json:"id"`
	Usage string `json:"usage"`
}

type ResourceVirtualMachineNetworkSpec struct {
	Id      string `json:"id"`
	Dns     string `json:"dns"`
	Ipv4    string `json:"ipv4"`
	Ipv6    string `json:"ipv6"`
	Mask    string `json:"mask"`
	Gateway string `json:"gateway"`
}

type ResourceVirtualMachineNetworkStatus struct {
	Id string `json:"id"`
}

type ResourceVirtualMachineOperatingSystemSpec struct {
	Id string `json:"id"`
}

type ResourceVirtualMachineOperatingSystemStatus struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Version      string `json:"version"`
	HostName     string `json:"hostName"`
	PowerState   string `json:"powerState"`
	ToolVersion  string `json:"toolVersion"` //many hypervisors install tools into the guest operating system
	Architecture string `json:"architecture"`
}

type ResourceVirtualMachineCpuSpec struct {
	Id    string `json:"id"`
	Count int    `json:"count"`
}
type ResourceVirtualMachineCpuStatus struct {
	Id    string `json:"id"`
	Usage string `json:"usage"`
}

type ResourceVirtualMachineMemorySpec struct {
	Id   string `json:"id"`
	Size int    `json:"size"`
}

type ResourceVirtualMachineMemoryStatus struct {
	Id    string `json:"id"`
	Usage string `json:"usage"`
}

type ResourceVirtualMachineTagSpec struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
