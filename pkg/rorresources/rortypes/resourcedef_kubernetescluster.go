package rortypes

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	vitiv1alpha1 "github.com/vitistack/common/pkg/v1alpha1"
)

// ResourceKubernetesCluster uses the external types from vitistack/common
type ResourceKubernetesCluster struct {
	Spec   ResourceKubernetesClusterSpec   `json:"spec"`
	Status ResourceKubernetesClusterStatus `json:"status"`
}

type ResourceKubernetesClusterSpec struct {
	SlackChannels []string              `json:"slackChannels"`
	VitiSpec      KubernetesClusterSpec `json:"vitiSpec,omitzero" bson:"vitispec,omitempty"`
}

type ResourceKubernetesClusterStatus struct {
	ProviderStatus KubernetesClusterStatus      `json:"providerstatus,omitzero" bson:"providerstatus,omitempty"`
	AgentStatus    KubernetesClusterAgentStatus `json:"agentstatus,omitzero" bson:"agentstatus,omitempty"`
}

type KubernetesClusterAgentStatus struct {
	ClusterId          string                            `json:"clusterId,omitempty" bson:"clusterid,omitempty"`
	ClusterName        string                            `json:"clusterName,omitempty" bson:"clustername,omitempty"`
	KubernetesProvider providermodels.ProviderType       `json:"kubernetesProvider,omitempty" bson:"kubernetesprovider,omitempty"`
	Az                 string                            `json:"az,omitempty" bson:"az,omitempty"`
	Region             string                            `json:"region,omitempty" bson:"region,omitempty"`
	Country            string                            `json:"country,omitempty" bson:"country,omitempty"`
	Workspace          string                            `json:"workspaceId,omitempty" bson:"workspaceid,omitempty"`
	Environment        string                            `json:"environment,omitempty" bson:"environment,omitempty"`
	Datacenter         string                            `json:"datacenter,omitempty" bson:"datacenter,omitempty"`
	Nodes              KubernetesClusterAgentStatusNodes `json:"nodes,omitzero" bson:"nodes,omitempty"`
	Versions           map[string]string                 `json:"versions,omitempty" bson:"versions,omitempty"`
	Urls               map[string]string                 `json:"urls,omitempty" bson:"urls,omitempty"`
	CreatedAt          time.Time                         `json:"createdAt,omitempty" bson:"createdat,omitempty"`
	LastSeen           time.Time                         `json:"lastSeen,omitempty" bson:"lastseen,omitempty"`
}

type KubernetesClusterAgentStatusNodes struct {
	ControllPlane []KubernetesClusterAgentStatusNodesNodepoolsNodes `json:"controlPlane,omitzero" bson:"controlplane,omitempty"`
	Nodepools     []KubernetesClusterAgentStatusNodesNodepools      `json:"nodepools,omitzero" bson:"nodepools,omitempty"`
}
type KubernetesClusterAgentStatusNodesNodepools struct {
	Name  string                                            `json:"name,omitempty" bson:"name,omitempty"`
	Nodes []KubernetesClusterAgentStatusNodesNodepoolsNodes `json:"nodes,omitzero" bson:"nodes,omitempty"`
}
type KubernetesClusterAgentStatusNodesNodepoolsNodes struct {
	Name              string                                                  `json:"name,omitempty" bson:"name,omitempty"`
	Cpu               KubernetesClusterAgentStatusNodesNodepoolsNodesResource `json:"cpu,omitzero" bson:"cpu,omitempty"`
	Memory            KubernetesClusterAgentStatusNodesNodepoolsNodesResource `json:"memory,omitzero" bson:"memory,omitempty"`
	Architecture      string                                                  `json:"architecture,omitempty" bson:"architecture,omitempty"`
	KubernetesVersion string                                                  `json:"kubernetesVersion,omitempty" bson:"kubernetesversion,omitempty"`
}

type KubernetesClusterAgentStatusNodesNodepoolsNodesResource struct {
	Capacity Quantity `json:"capacity,omitzero" bson:"capacity,omitempty"`
	Used     Quantity `json:"allocated,omitzero" bson:"allocated,omitempty"`
}

func (r *KubernetesClusterAgentStatusNodesNodepoolsNodesResource) UsedPercent() float64 {
	if r.Capacity.Value() == 0 {
		return 0
	}
	return getRoundedValue((float64(r.Used.Value())/float64(r.Capacity.Value()))*100, 2)
}

func (r *KubernetesClusterAgentStatus) GetNodepoolCount() int {
	return len(r.Nodes.Nodepools)
}

func (r *KubernetesClusterAgentStatus) GetNodeCount() int {
	nodeCount := 0
	for _, nodepool := range r.Nodes.Nodepools {
		nodeCount += len(nodepool.Nodes)
	}
	return nodeCount
}

func (r *KubernetesClusterAgentStatus) GetKubernetesVersion() string {
	if len(r.Nodes.ControllPlane) > 0 && r.Nodes.ControllPlane[0].KubernetesVersion != "" {
		return r.Nodes.ControllPlane[0].KubernetesVersion
	}
	return "Unknown"
}
func (r *KubernetesClusterAgentStatus) GetVersionByKey(key string) string {
	version, _ := r.Versions[key]
	return nomalizeVersion(version)
}

// nomalizeVersion returns "Unknown" if the version string is empty, otherwise it returns the version string ensuring it starts with a v

func nomalizeVersion(version string) string {
	if version == "" {
		return "Unknown"
	}
	if version[0] != 'v' && version[0] != 'V' {
		return "v" + version
	}
	return version
}

func (r *KubernetesClusterAgentStatus) GetUrlByKey(key string) string {
	if url, ok := r.Urls[key]; ok {
		return url
	}
	return "Unknown"
}

func (r *KubernetesClusterAgentStatus) GetStatus() string {
	diff := time.Since(r.LastSeen)
	if diff < 5*time.Minute {
		return "ok"
	} else if diff < 15*time.Minute {
		return "warning"
	} else {
		return "error"
	}
}

// GetTotalCpu returns the total CPU resources of the cluster by summing the CPU of all nodes in the control plane and node pools.
func (r *KubernetesClusterAgentStatus) GetTotalCpu() *Quantity {
	cpu := &Quantity{}
	cpu.Add(r.GetControlPlaneCpu().Quantity)
	cpu.Add(r.GetNodePoolCpu().Quantity)
	return cpu
}

// GetTotalMemory returns the total memory resources of the cluster by summing the memory of all nodes in the control plane and node pools.
func (r *KubernetesClusterAgentStatus) GetTotalMemory() *Quantity {
	memory := &Quantity{}
	memory.Add(r.GetControlPlaneMemory().Quantity)
	memory.Add(r.GetNodePoolMemory().Quantity)
	return memory
}

func (r *KubernetesClusterAgentStatus) GetControlPlaneCpu() *Quantity {
	cpu := &Quantity{}
	for _, node := range r.Nodes.ControllPlane {
		cpu.Add(node.Cpu.Capacity.Quantity)
	}
	return cpu
}

func (r *KubernetesClusterAgentStatus) GetControlPlaneMemory() *Quantity {
	memory := &Quantity{}
	for _, node := range r.Nodes.ControllPlane {
		memory.Add(node.Memory.Capacity.Quantity)
	}
	return memory
}

func (r *KubernetesClusterAgentStatus) GetNodePoolCpu() *Quantity {
	cpu := &Quantity{}
	for _, nodepool := range r.Nodes.Nodepools {
		for _, node := range nodepool.Nodes {
			cpu.Add(node.Cpu.Capacity.Quantity)
		}
	}
	return cpu
}

func (r *KubernetesClusterAgentStatus) GetNodePoolMemory() *Quantity {
	memory := &Quantity{}
	for _, nodepool := range r.Nodes.Nodepools {
		for _, node := range nodepool.Nodes {
			memory.Add(node.Memory.Capacity.Quantity)
		}
	}
	return memory
}

// GetControlPlaneUsedCpu returns the total used CPU across the control plane.
func (r *KubernetesClusterAgentStatus) GetControlPlaneUsedCpu() *Quantity {
	cpu := &Quantity{}
	for _, node := range r.Nodes.ControllPlane {
		cpu.Add(node.Cpu.Used.Quantity)
	}
	return cpu
}

// GetNodePoolUsedCpu returns the total used CPU across all node pools.
func (r *KubernetesClusterAgentStatus) GetNodePoolUsedCpu() *Quantity {
	cpu := &Quantity{}
	for _, nodepool := range r.Nodes.Nodepools {
		for _, node := range nodepool.Nodes {
			cpu.Add(node.Cpu.Used.Quantity)
		}
	}
	return cpu
}

// GetTotalUsedCpu returns the total used CPU across the control plane and all node pools.
func (r *KubernetesClusterAgentStatus) GetTotalUsedCpu() *Quantity {
	cpu := &Quantity{}
	cpu.Add(r.GetControlPlaneUsedCpu().Quantity)
	cpu.Add(r.GetNodePoolUsedCpu().Quantity)
	return cpu
}

// GetControlPlaneUsedMemory returns the total used memory across the control plane.
func (r *KubernetesClusterAgentStatus) GetControlPlaneUsedMemory() *Quantity {
	memory := &Quantity{}
	for _, node := range r.Nodes.ControllPlane {
		memory.Add(node.Memory.Used.Quantity)
	}
	return memory
}

// GetNodePoolUsedMemory returns the total used memory across all node pools.
func (r *KubernetesClusterAgentStatus) GetNodePoolUsedMemory() *Quantity {
	memory := &Quantity{}
	for _, nodepool := range r.Nodes.Nodepools {
		for _, node := range nodepool.Nodes {
			memory.Add(node.Memory.Used.Quantity)
		}
	}
	return memory
}

// GetTotalUsedMemory returns the total used memory across the control plane and all node pools.
func (r *KubernetesClusterAgentStatus) GetTotalUsedMemory() *Quantity {
	memory := &Quantity{}
	memory.Add(r.GetControlPlaneUsedMemory().Quantity)
	memory.Add(r.GetNodePoolUsedMemory().Quantity)
	return memory
}

// GetControlPlaneCpuResource returns the aggregate CPU resource (capacity + used) for the control plane.
func (r *KubernetesClusterAgentStatus) GetControlPlaneCpuResource() *KubernetesClusterAgentStatusNodesNodepoolsNodesResource {
	res := &KubernetesClusterAgentStatusNodesNodepoolsNodesResource{}
	for _, node := range r.Nodes.ControllPlane {
		res.Capacity.Add(node.Cpu.Capacity.Quantity)
		res.Used.Add(node.Cpu.Used.Quantity)
	}
	return res
}

// GetNodePoolCpuResource returns the aggregate CPU resource (capacity + used) for all node pools.
func (r *KubernetesClusterAgentStatus) GetNodePoolCpuResource() *KubernetesClusterAgentStatusNodesNodepoolsNodesResource {
	res := &KubernetesClusterAgentStatusNodesNodepoolsNodesResource{}
	for _, nodepool := range r.Nodes.Nodepools {
		for _, node := range nodepool.Nodes {
			res.Capacity.Add(node.Cpu.Capacity.Quantity)
			res.Used.Add(node.Cpu.Used.Quantity)
		}
	}
	return res
}

// GetCpuResource returns an aggregate CPU resource (capacity + used) across the control plane and all node pools.
func (r *KubernetesClusterAgentStatus) GetCpuResource() *KubernetesClusterAgentStatusNodesNodepoolsNodesResource {
	res := &KubernetesClusterAgentStatusNodesNodepoolsNodesResource{}
	cp := r.GetControlPlaneCpuResource()
	np := r.GetNodePoolCpuResource()
	res.Capacity.Add(cp.Capacity.Quantity)
	res.Used.Add(cp.Used.Quantity)
	res.Capacity.Add(np.Capacity.Quantity)
	res.Used.Add(np.Used.Quantity)
	return res
}

// GetControlPlaneMemoryResource returns the aggregate memory resource (capacity + used) for the control plane.
func (r *KubernetesClusterAgentStatus) GetControlPlaneMemoryResource() *KubernetesClusterAgentStatusNodesNodepoolsNodesResource {
	res := &KubernetesClusterAgentStatusNodesNodepoolsNodesResource{}
	for _, node := range r.Nodes.ControllPlane {
		res.Capacity.Add(node.Memory.Capacity.Quantity)
		res.Used.Add(node.Memory.Used.Quantity)
	}
	return res
}

// GetNodePoolMemoryResource returns the aggregate memory resource (capacity + used) for all node pools.
func (r *KubernetesClusterAgentStatus) GetNodePoolMemoryResource() *KubernetesClusterAgentStatusNodesNodepoolsNodesResource {
	res := &KubernetesClusterAgentStatusNodesNodepoolsNodesResource{}
	for _, nodepool := range r.Nodes.Nodepools {
		for _, node := range nodepool.Nodes {
			res.Capacity.Add(node.Memory.Capacity.Quantity)
			res.Used.Add(node.Memory.Used.Quantity)
		}
	}
	return res
}

// GetMemoryResource returns an aggregate memory resource (capacity + used) across the control plane and all node pools.
func (r *KubernetesClusterAgentStatus) GetMemoryResource() *KubernetesClusterAgentStatusNodesNodepoolsNodesResource {
	res := &KubernetesClusterAgentStatusNodesNodepoolsNodesResource{}
	cp := r.GetControlPlaneMemoryResource()
	np := r.GetNodePoolMemoryResource()
	res.Capacity.Add(cp.Capacity.Quantity)
	res.Used.Add(cp.Used.Quantity)
	res.Capacity.Add(np.Capacity.Quantity)
	res.Used.Add(np.Used.Quantity)
	return res
}

// Type aliases for convenience and backward compatibility
type KubernetesClusterSpec = vitiv1alpha1.KubernetesClusterSpec
type KubernetesClusterSpecData = vitiv1alpha1.KubernetesClusterSpecData
type KubernetesClusterSpecTopology = vitiv1alpha1.KubernetesClusterSpecTopology
type KubernetesClusterSpecControlPlane = vitiv1alpha1.KubernetesClusterSpecControlPlane
type KubernetesClusterSpecMetadataDetails = vitiv1alpha1.KubernetesClusterSpecMetadataDetails
type KubernetesClusterStorage = vitiv1alpha1.KubernetesClusterStorage
type KubernetesClusterWorkers = vitiv1alpha1.KubernetesClusterWorkers
type KubernetesClusterNodePool = vitiv1alpha1.KubernetesClusterNodePool
type KubernetesClusterTaint = vitiv1alpha1.KubernetesClusterTaint
type KubernetesClusterAutoscalingConfig = vitiv1alpha1.KubernetesClusterAutoscalingConfig
type KubernetesClusterAutoscalingSpec = vitiv1alpha1.KubernetesClusterAutoscalingSpec
type KubernetesClusterStatus = vitiv1alpha1.KubernetesClusterStatus
type KubernetesClusterClusterState = vitiv1alpha1.KubernetesClusterClusterState
type KubernetesClusterEndpoint = vitiv1alpha1.KubernetesClusterEndpoint
type KubernetesClusterStatusCondition = vitiv1alpha1.KubernetesClusterStatusCondition
type KubernetesClusterStatusPrice = vitiv1alpha1.KubernetesClusterStatusPrice
type KubernetesClusterClusterDetails = vitiv1alpha1.KubernetesClusterClusterDetails
type KubernetesClusterStatusClusterStatusResources = vitiv1alpha1.KubernetesClusterStatusClusterStatusResources
type KubernetesClusterStatusClusterStatusResource = vitiv1alpha1.KubernetesClusterStatusClusterStatusResource
type KubernetesClusterControlPlaneStatus = vitiv1alpha1.KubernetesClusterControlPlaneStatus
type KubernetesClusterNodePoolStatus = vitiv1alpha1.KubernetesClusterNodePoolStatus
type KubernetesClusterVersion = vitiv1alpha1.KubernetesClusterVersion
type KubernetesClusterCondition = vitiv1alpha1.KubernetesClusterCondition
