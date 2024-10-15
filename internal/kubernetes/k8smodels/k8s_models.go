package k8smodels

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/models/providers"
)

type Node struct {
	Name                    string                     `json:"name"`
	Created                 time.Time                  `json:"created"`
	OsImage                 string                     `json:"osImage"`
	ClusterName             string                     `json:"clusterName"`
	Workspace               string                     `json:"workspace"`
	Datacenter              string                     `json:"datacenter"`
	MachineName             string                     `json:"machineName"`
	Labels                  map[string]string          `json:"labels"`
	Annotations             map[string]string          `json:"annotations"`
	Resources               apicontracts.NodeResources `json:"resources"`
	Architecture            string                     `json:"architecture"`
	ContainerRuntimeVersion string                     `json:"containerRuntimeVersion"`
	KernelVersion           string                     `json:"kernelVersion"`
	KubeProxyVersion        string                     `json:"kubeProxyVersion"`
	KubeletVersion          string                     `json:"kubeletVersion"`
	OperatingSystem         string                     `json:"operatingSystem"`
	Provider                providers.ProviderType     `json:"provider"`
}

type NhnTooling struct {
	Version      string   `json:"version"`
	Branch       string   `json:"branch"`
	Environment  string   `json:"environment"`
	AccessGroups []string `json:"accessGroups"`
}
