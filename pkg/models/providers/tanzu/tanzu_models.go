package tanzu

type TanzuKubernetesClusterInput struct {
	Name         string       `json:"name"`
	Namespace    string       `json:"namespace"`
	DataCenter   string       `json:"dataCenter"`
	ControlPlane ControlPlane `json:"controlPlane"`
	NodePools    []NodePool   `json:"nodePools"`
}

type ControlPlane struct {
	HighAvailability  bool   `json:"highAvailability"`
	KubernetesVersion string `json:"kubernetesVersion"`
	VmClass           string `json:"vmClass"`
}

type NodePool struct {
	Name              string `json:"name"`
	KubernetesVersion string `json:"kubernetesVersion"`
	Replicas          int64  `json:"replicas"`
	VmClass           string `json:"vmClass"`
}

type DataCenter string

const (
	DataCenterUnknown  DataCenter = ""
	DataCenterTrd1Cl01 DataCenter = "trd1-cl01"
	DataCenterTrd1Cl02 DataCenter = "trd1-cl02"
	DataCenterOsl1Cl01 DataCenter = "osl1-cl01"
)
