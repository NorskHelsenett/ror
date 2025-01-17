package rortypes

type ResourceVirtualMachine struct {
	Id       string                       `json:"id"`
	Spec     ResourceVirtualMachineSpec   `json:"spec"`
	Status   ResourceVirtualMachineStatus `json:"status"`
	Provider string                       `json:"provider"`
}

// things we can change
type ResourceVirtualMachineSpec struct {
	Cpu    ResourceVirtualMachineCpuSpec    `json:"cpu"`
	Name   string                           `json:"name"`
	Disks  []ResourceVirtualMachineDiskSpec `json:"disks"`
	Memory ResourceVirtualMachineMemorySpec `json:"memory"`
}

// things we can't change
type ResourceVirtualMachineStatus struct {
	Cpu             ResourceVirtualMachineCpuStatus             `json:"cpu"`
	Tags            []ResourceVirtualMachineTag                 `json:"tags"`
	Disks           []ResourceVirtualMachineDiskStatus          `json:"disks"`
	Memory          ResourceVirtualMachineMemoryStatus          `json:"memory"`
	Networks        []ResourceVirtualMachineNetworkStatus       `json:"networks"`
	OperatingSystem ResourceVirtualMachineOperatingSystemStatus `json:"operatingSystem"`
}

type ResourceVirtualMachineDiskSpec struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	SizeBytes int    `json:"sizeBytes"`
}

type ResourceVirtualMachineDiskStatus struct {
	UsageBytes string `json:"usageBytes"`

	// is this disk mounted by the os. A disk might be attached to the vm but
	// mountd by the OS. Might not be set based on the avalible
	IsMounted string `json:"isMounted"`

	ResourceVirtualMachineDiskSpec `json:"spec"`
}

type ResourceVirtualMachineNetworkStatus struct {
	Id      string `json:"id"`
	Dns     string `json:"dns"`
	Ipv4    string `json:"ipv4"`
	Ipv6    string `json:"ipv6"`
	Mask    string `json:"mask"`
	Gateway string `json:"gateway"`
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
	CoresPerSocket int    `json:"coresPerSocket"` //cores per socket
}

type ResourceVirtualMachineCpuStatus struct {
	Unit                          string `json:"unit"` //describes what unit the usage is given in
	Usage                         int    `json:"usage"`
	ResourceVirtualMachineCpuSpec `json:"spec"`
}

type ResourceVirtualMachineMemorySpec struct {
	Id        string `json:"id"`
	SizeBytes int    `json:"sizeBytes"`
}

type ResourceVirtualMachineMemoryStatus struct {
	Unit                             string `json:"unit"` //describes what unit the usage is given in
	Usage                            int    `json:"usage"`
	ResourceVirtualMachineMemorySpec `json:"spec"`
}

type ResourceVirtualMachineTag struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
