package rortypes

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

type ResourceMachine struct {
	Spec   MachineSpec   `json:"spec"`
	Status MachineStatus `json:"status"`
}

type MachineProviderType string

const (
	MachineProviderTypeKubevirt MachineProviderType = "kubevirt"
	MachineProviderTypeProxmox  MachineProviderType = "proxmox"
	MachineProviderTypeVMware   MachineProviderType = "vmware"
)

type MachineSpec struct {
	// The name of the machine
	Name string `json:"name,omitempty"`

	// The machine class of the machine (e.g., t3.medium, Standard_B2s, n1-standard-2)
	MachineClass string `json:"machineClass,omitempty"`

	// The provider-specific machine type override
	MachineType string `json:"machineType,omitempty"`

	// CPU configuration
	CPU MachineCPU `json:"cpu"`

	// Memory configuration in bytes
	Memory int64 `json:"memory,omitempty"`

	// Disk configuration
	Disks []MachineSpecDisk `json:"disks,omitempty"`

	// Network configuration
	Network MachineNetwork `json:"network"`

	// Operating system configuration
	OS MachineOS `json:"os"`

	// Machine Provider
	Provider MachineProviderType `json:"provider"`

	// Cloud provider configuration
	ProviderConfig CloudProviderConfig `json:"providerConfig"`

	// SSH key configuration
	SSHKeys []string `json:"sshKeys,omitempty"`

	// User data script to run on first boot
	UserData string `json:"userData,omitempty"`

	// Tags/labels to apply to the machine
	Tags map[string]string `json:"tags,omitempty"`

	// Security groups or firewall rules
	SecurityGroups []string `json:"securityGroups,omitempty"`

	// Whether to enable monitoring
	Monitoring bool `json:"monitoring,omitempty"`

	// Backup configuration
	Backup MachineBackup `json:"backup"`

	// Cloud-init configuration for VM initialization
	CloudInit *CloudInitConfig `json:"cloudInit,omitempty"`
}

type MachineCPU struct {
	// Number of CPU cores
	Cores int `json:"cores,omitempty"`
	// Number of threads per core
	ThreadsPerCore int `json:"threadsPerCore,omitempty"`
	// Number of CPU sockets
	Sockets int `json:"sockets,omitempty"`
}

type MachineSpecDisk struct {
	// Name of the disk
	Name string `json:"name,omitempty"`
	// Size of the disk in GB
	SizeGB int64 `json:"sizeGB,omitempty"`
	// Type of the disk (e.g., gp2, gp3, pd-ssd, Premium_LRS)
	Type string `json:"type,omitempty"`
	// Whether this is the boot disk
	Boot bool `json:"boot,omitempty"`
	// Device name (e.g., /dev/sda, /dev/nvme0n1)
	Device string `json:"device,omitempty"`
	// Encryption settings
	Encrypted bool `json:"encrypted,omitempty"`
	// IOPS for the disk (if supported by provider)
	IOPS int `json:"iops,omitempty"`
	// Throughput in MB/s (if supported by provider)
	Throughput int `json:"throughput,omitempty"`
}

type MachineNetwork struct {
	// NetworkNamespaceName is the name of the NetworkNamespace object in the same namespace
	// that provides network configuration (VLAN, IP prefix, etc.) for this machine.
	// When set, operators will look up this specific NetworkNamespace by name instead of
	// listing all NetworkNamespaces in the namespace and picking the first one.
	// This is typically propagated from the KubernetesCluster spec.
	NetworkNamespaceName string `json:"networkNamespaceName,omitempty"`
	// VPC/Virtual Network ID
	VPC string `json:"vpc,omitempty"`
	// Subnet ID
	Subnet string `json:"subnet,omitempty"`
	// Whether to assign a public IP
	AssignPublicIP bool `json:"assignPublicIP,omitempty"`
	// Static private IP address
	PrivateIP string `json:"privateIP,omitempty"`
	// Static public IP address or Elastic IP
	PublicIP string `json:"publicIP,omitempty"`
	// Network interfaces
	Interfaces []NetworkInterface `json:"interfaces,omitempty"`
}

type NetworkInterface struct {
	// Name of the network interface
	Name string `json:"name,omitempty"`
	// Subnet for this interface
	Subnet string `json:"subnet,omitempty"`
	// Security groups for this interface
	SecurityGroups []string `json:"securityGroups,omitempty"`
	// Whether this is the primary interface
	Primary bool `json:"primary,omitempty"`
}

type MachineOS struct {
	// Operating system family (linux, windows)
	Family string `json:"family,omitempty"`
	// Distribution (ubuntu, centos, rhel, windows-server, debian, alpine)
	Distribution string `json:"distribution,omitempty"`
	// Version of the OS
	Version string `json:"version,omitempty"`
	// Architecture (amd64, arm64)
	Architecture string `json:"architecture,omitempty"`
	// Image ID/AMI/Template ID
	ImageID string `json:"imageID,omitempty"`
	// ISO URI for custom installations
	ISOUri string `json:"isoUri,omitempty"`
	// Image family or marketplace image
	ImageFamily string `json:"imageFamily,omitempty"`
}

type CloudProviderConfig struct {
	// Provider name (aws, azure, gcp, vsphere, openstack)
	Name string `json:"name,omitempty"`
	// Region where the machine should be created
	Region string `json:"region,omitempty"`
	// Availability zone
	Zone string `json:"zone,omitempty"`
	// Provider-specific configuration
	Config map[string]string `json:"config,omitempty"`
	// Credentials reference
	CredentialsRef *CredentialsReference `json:"credentialsRef,omitempty"`
}

type CredentialsReference struct {
	// Name of the secret containing credentials
	SecretName string `json:"secretName,omitempty"`
	// Namespace of the secret (defaults to machine namespace)
	Namespace string `json:"namespace,omitempty"`
}

type MachineBackup struct {
	// Whether to enable automated backups
	Enabled bool `json:"enabled,omitempty"`
	// Backup schedule (cron format)
	Schedule string `json:"schedule,omitempty"`
	// Retention period in days
	RetentionDays int `json:"retentionDays,omitempty"`
}

// CloudInitConfig defines the cloud-init configuration for a VM
// Cloud-init data can be provided inline, from a ConfigMap, or from a Secret
type CloudInitConfig struct {
	// Type specifies the cloud-init type to use
	// Valid values are: noCloud, configDrive
	Type CloudInitType `json:"type,omitempty"`

	// UserData contains cloud-init user data content directly
	// This is typically a cloud-config YAML starting with #cloud-config
	UserData string `json:"userData,omitempty"`

	// UserDataBase64 contains base64-encoded cloud-init user data
	// Use this for binary or pre-encoded data
	UserDataBase64 string `json:"userDataBase64,omitempty"`

	// NetworkData contains cloud-init network configuration directly
	// This follows the cloud-init network config v1 or v2 format
	NetworkData string `json:"networkData,omitempty"`

	// NetworkDataBase64 contains base64-encoded network data
	NetworkDataBase64 string `json:"networkDataBase64,omitempty"`

	// UserDataSecretRef references a Secret containing cloud-init user data
	// The Secret should have a key named 'userdata' or a custom key specified in UserDataSecretKey
	UserDataSecretRef *CloudInitSecretRef `json:"userDataSecretRef,omitempty"`

	// NetworkDataSecretRef references a Secret containing cloud-init network data
	// The Secret should have a key named 'networkdata' or a custom key specified in NetworkDataSecretKey
	NetworkDataSecretRef *CloudInitSecretRef `json:"networkDataSecretRef,omitempty"`

	// UserDataConfigMapRef references a ConfigMap containing cloud-init user data
	// The ConfigMap should have a key named 'userdata' or a custom key specified
	UserDataConfigMapRef *CloudInitConfigMapRef `json:"userDataConfigMapRef,omitempty"`

	// NetworkDataConfigMapRef references a ConfigMap containing cloud-init network data
	// The ConfigMap should have a key named 'networkdata' or a custom key specified
	NetworkDataConfigMapRef *CloudInitConfigMapRef `json:"networkDataConfigMapRef,omitempty"`
}

// CloudInitType specifies the type of cloud-init volume to use
type CloudInitType string

const (
	// CloudInitTypeNoCloud uses cloud-init NoCloud datasource
	// This is the most common type for KubeVirt VMs
	CloudInitTypeNoCloud CloudInitType = "noCloud"

	// CloudInitTypeConfigDrive uses cloud-init ConfigDrive datasource
	// This is useful for OpenStack compatibility
	CloudInitTypeConfigDrive CloudInitType = "configDrive"
)

// CloudInitSecretRef references a Secret containing cloud-init data
type CloudInitSecretRef struct {
	// Name is the name of the Secret in the same namespace as the Machine
	Name string `json:"name"`

	// Key is the key within the Secret containing the cloud-init data
	// Defaults to 'userdata' for user data secrets and 'networkdata' for network data secrets
	Key string `json:"key,omitempty"`
}

// CloudInitConfigMapRef references a ConfigMap containing cloud-init data
type CloudInitConfigMapRef struct {
	// Name is the name of the ConfigMap in the same namespace as the Machine
	Name string `json:"name"`

	// Key is the key within the ConfigMap containing the cloud-init data
	// Defaults to 'userdata' for user data configs and 'networkdata' for network data configs
	Key string `json:"key,omitempty"`
}

type MachineStatus struct {
	// Current phase of the machine (Pending, Creating, Running, Stopping, Stopped, Terminating, Terminated, Failed)
	Phase string `json:"phase,omitempty"`

	// Detailed status message
	Message string `json:"message,omitempty"`

	// The unique identifier assigned by the provider
	ProviderID string `json:"providerID,omitempty"`

	// Internal machine identifier
	MachineID string `json:"machineID,omitempty"`

	// The current state of the machine
	State string `json:"state,omitempty"`

	// The last time the machine status was updated
	LastUpdated metav1.Time `json:"lastUpdated"`

	// The provider that created this machine
	Provider MachineProviderType `json:"provider,omitempty"`

	// The region where the machine is located
	Region string `json:"region,omitempty"`

	// The zone where the machine is located
	Zone string `json:"zone,omitempty"`

	// The IP addresses of the machine
	IPAddresses []string `json:"ipAddresses,omitempty"`

	// IPv6 addresses of the machine
	IPv6Addresses []string `json:"ipv6Addresses,omitempty"`

	// Public IP addresses
	PublicIPAddresses []string `json:"publicIPAddresses,omitempty"`

	// Private IP addresses
	PrivateIPAddresses []string `json:"privateIPAddresses,omitempty"`

	// The machine's hostname
	Hostname string `json:"hostname,omitempty"`

	// The machine's CPU architecture
	Architecture string `json:"architecture,omitempty"`

	// The machine's operating system
	OperatingSystem string `json:"operatingSystem,omitempty"`

	// The machine's operating system version
	OperatingSystemVersion string `json:"operatingSystemVersion,omitempty"`

	// The machine's kernel version
	KernelVersion string `json:"kernelVersion,omitempty"`

	// Actual CPU count
	CPUs int `json:"cpus,omitempty"`

	// Actual memory in bytes
	Memory int64 `json:"memory,omitempty"`

	// Actual disk information
	Disks []MachineStatusDisk `json:"disks,omitempty"`

	// Network interface information
	NetworkInterfaces []NetworkInterfaceStatus `json:"networkInterfaces,omitempty"`

	// Conditions represent the latest available observations of the machine's state
	Conditions []MachineCondition `json:"conditions,omitempty"`

	// Boot time of the machine
	BootTime *metav1.Time `json:"bootTime,omitempty"`

	// Creation time of the machine
	CreationTime *metav1.Time `json:"creationTime,omitempty"`

	// Failure reason if the machine failed to be created
	FailureReason *string `json:"failureReason,omitempty"`

	// Failure message if the machine failed to be created
	FailureMessage *string `json:"failureMessage,omitempty"`
}

type MachineDisk struct {
	// The disk's name
	Name string `json:"name"`
	// The disk's size in bytes
	Size int64 `json:"size"`
	// The disk's type (e.g., SSD, HDD)
	Type string `json:"type"`
	// The disk's mount point
	MountPoint string `json:"mountPoint"`
	// The disk's filesystem type (e.g., ext4, xfs)
	FilesystemType string `json:"filesystemType"`
	// The disk's UUID
	UUID string `json:"uuid"`
	// The disk's label
	Label string `json:"label"`
	// The disk's serial number
	SerialNumber string `json:"serialNumber"`
}

type MachineStatusDisk struct {
	// The disk's name
	Name string `json:"name,omitempty"`
	// The disk's size in bytes
	Size int64 `json:"size,omitempty"`
	// The disk's type (e.g., SSD, HDD, gp2, gp3)
	Type string `json:"type,omitempty"`
	// The disk's mount point
	MountPoint string `json:"mountPoint,omitempty"`
	// PVC name
	PVCName string `json:"pvcName,omitempty"`
	// Volume mode, filesystem or block
	VolumeMode string `json:"volumeMode,omitempty"`
	// Access modes, readwriteonce, readwritemany, readonlymany
	AccessModes []string `json:"accessModes,omitempty"`
	// The disk's filesystem type (e.g., ext4, xfs)
	FilesystemType string `json:"filesystemType,omitempty"`
	// The disk's UUID
	UUID string `json:"uuid,omitempty"`
	// The disk's label
	Label string `json:"label,omitempty"`
	// The disk's serial number
	SerialNumber string `json:"serialNumber,omitempty"`
	// Device path (e.g., /dev/sda)
	Device string `json:"device,omitempty"`
	// Used space in bytes
	UsedBytes int64 `json:"usedBytes,omitempty"`
	// Available space in bytes
	AvailableBytes int64 `json:"availableBytes,omitempty"`
	// Usage percentage as string (e.g., "75.5%")
	UsagePercent string `json:"usagePercent,omitempty"`
}

type NetworkInterfaceStatus struct {
	// Name of the network interface
	Name string `json:"name,omitempty"`
	// MAC address
	MACAddress string `json:"macAddress,omitempty"`
	// IP addresses assigned to this interface
	IPAddresses []string `json:"ipAddresses,omitempty"`
	// IPv6 addresses assigned to this interface
	IPv6Addresses []string `json:"ipv6Addresses,omitempty"`
	// Interface state (up, down)
	State string `json:"state,omitempty"`
	// MTU size
	MTU int `json:"mtu,omitempty"`
	// Interface type (ethernet, wifi, etc.)
	Type string `json:"type,omitempty"`
}

type MachineCondition struct {
	// Type of condition
	Type string `json:"type,omitempty"`
	// Status of the condition (True, False, Unknown)
	Status string `json:"status,omitempty"`
	// Last time the condition transitioned from one status to another
	LastTransitionTime metav1.Time `json:"lastTransitionTime"`
	// The reason for the condition's last transition
	Reason string `json:"reason,omitempty"`
	// A human readable message indicating details about the transition
	Message string `json:"message,omitempty"`
}

// Common machine phases
const (
	MachinePhasePending     = "Pending"
	MachinePhaseCreating    = "Creating"
	MachinePhaseRunning     = "Running"
	MachinePhaseStopping    = "Stopping"
	MachinePhaseStopped     = "Stopped"
	MachinePhaseTerminating = "Terminating"
	MachinePhaseTerminated  = "Terminated"
	MachinePhaseFailed      = "Failed"
)

// Common machine condition types
const (
	MachineConditionReady               = "Ready"
	MachineConditionNetworkReady        = "NetworkReady"
	MachineConditionBootstrapReady      = "BootstrapReady"
	MachineConditionInfrastructureReady = "InfrastructureReady"
	MachineConditionDrainReady          = "DrainReady"
	MachineConditionBackupReady         = "BackupReady"
)
