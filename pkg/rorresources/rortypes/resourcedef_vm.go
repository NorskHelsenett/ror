package rortypes

import metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

type VmId string

type ResourceVirtualMachine struct {
	// This is the ID of the vm in the hypervisor layer
	ExternalId VmId                         `json:"externalId"`
	Spec       ResourceVirtualMachineSpec   `json:"spec"`
	Status     ResourceVirtualMachineStatus `json:"status"`
	Provider   string                       `json:"provider"`
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
	LastUpdated     metav1.Time                                 `json:"lastUpdated"`
	Location        string                                      `json:"location"`
	Cpu             ResourceVirtualMachineCpuStatus             `json:"cpu"`
	Tags            map[string]ResourceVirtualMachineTag        `json:"tags"`
	State           ResourceVirtualMachineState                 `json:"state"`
	Disks           []ResourceVirtualMachineDiskStatus          `json:"disks"`
	Memory          ResourceVirtualMachineMemoryStatus          `json:"memory"`
	Networks        []ResourceVirtualMachineNetworkStatus       `json:"networks"`
	OperatingSystem ResourceVirtualMachineOperatingSystemStatus `json:"operatingSystem"`
}

type ResourceVirtualMachineState struct {
	State  string `json:"state"`
	Reason string `json:"reason"`
	Time   string `json:"time"`
}

type ResourceVirtualMachineDiskSpec struct {
	Id        string `json:"id"`
	Name      string `json:"name"`
	Type      string `json:"type"`
	SizeBytes int    `json:"sizeBytes"`
}

type ResourceVirtualMachineDiskStatus struct {
	UsageBytes int `json:"usageBytes"`

	// is this disk mounted by the os? A disk might be attached to the vm but
	// not mounted by the OS, it can also be unknown because the vm might not report
	// this or may not have tools installed, false can mean we dont know or that
	// it is actually not mounted.
	IsMounted bool `json:"isMounted"`

	ResourceVirtualMachineDiskSpec
}

type ResourceVirtualMachineNetworkStatus struct {
	Id      string `json:"id"`
	Dns     string `json:"dns"`
	Ipv4    string `json:"ipv4"`
	Ipv6    string `json:"ipv6"`
	Mask    string `json:"mask"`
	Gateway string `json:"gateway"`
	Mac     string `json:"mac"`
}

type ResourceVirtualMachineOperatingSystemStatus struct {
	Id           string `json:"id"`
	Name         string `json:"name"`
	Family       string `json:"family"`
	Version      string `json:"version"`
	HostName     string `json:"hostName"`
	PowerState   string `json:"powerState"`
	ToolVersion  string `json:"toolVersion"`
	Architecture string `json:"architecture"`
}

type ResourceVirtualMachineCpuSpec struct {
	Sockets        int `json:"sockets"`
	CoresPerSocket int `json:"coresPerSocket"` //cores per socket
}

type ResourceVirtualMachineCpuStatus struct {
	Unit  string `json:"unit"` //describes what unit the usage is given in
	Usage int    `json:"usage"`
	ResourceVirtualMachineCpuSpec
}

type ResourceVirtualMachineMemorySpec struct {
	SizeBytes int `json:"sizeBytes"`
}

type ResourceVirtualMachineMemoryStatus struct {
	Unit  string `json:"unit"` //describes what unit the usage is given in
	Usage int    `json:"usage"`
	ResourceVirtualMachineMemorySpec
}

type ResourceVirtualMachineTag struct {
	Key         string `json:"key"`
	Value       string `json:"value"`
	Description string `json:"description"`
}
