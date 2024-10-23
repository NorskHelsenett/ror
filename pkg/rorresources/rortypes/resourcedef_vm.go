package rortypes

type ResourceVm struct {
	Id      string            `json:"id"`
	Name    string            `json:"name"`
	Guest   ResourceVMGuest   `json:"guest"`
	Config  ResourceVMConfig  `json:"config"`
	Runtime ResourceVMRuntime `json:"runtime"`
	Tags    []ResourceVMTag   `json:"tags"`
}

type ResourceVMGuest struct {
	Id        string `json:"id"`
	Family    string `json:"family"`
	FullName  string `json:"full_name"`
	HostName  string `json:"host_name"`
	IpAddress string `json:"ip_address"`
	State     string `json:"state"`
}
type ResourceVMConfig struct {
	Name             string `json:"name"`
	MemorySize       int32  `json:"memory_size"`
	CpuCount         int32  `json:"cpu_count"`
	VirtualDiskCount int32  `json:"virtual_disk_count"`
	Annotation       string `json:"annotation"` //should this be a slice?
}
type ResourceVMRuntime struct {
	ConnectionState string `json:"connectionState"`
	PowerState      string `json:"powerState"`
	MaxCpu          int32  `json:"maxCpuUsage"`
	MaxMemory       int32  `json:"maxMemoryUsage"`
}

type ResourceVMTag struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
