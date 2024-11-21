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
	Id        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	SizeBytes int    `json:"size"`
}

type ResourceVirtualMachineDiskStatus struct {
	Id         string `json:"id"`
	UsageBytes string `json:"usage"`
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
	ToolVersion  string `json:"toolVersion"`
	Architecture string `json:"architecture"`
}

type ResourceVirtualMachineCpuSpec struct {
	Id             string `json:"id"`
	Sockets        int    `json:"sockets"`
	CoresPerSocket int    `json:"cores"` //cores per socket
}
type ResourceVirtualMachineCpuStatus struct {
	Id    string `json:"id"`
	Unit  string `json:"unit"` //describes what unit the usage is given in
	Usage string `json:"usage"`
}

type ResourceVirtualMachineMemorySpec struct {
	Id        string `json:"id"`
	SizeBytes int    `json:"size"`
}

type ResourceVirtualMachineMemoryStatus struct {
	Id    string `json:"id"`
	Unit  string `json:"unit"` //describes what unit the usage is given in
	Usage string `json:"usage"`
}

type ResourceVirtualMachineTagSpec struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
