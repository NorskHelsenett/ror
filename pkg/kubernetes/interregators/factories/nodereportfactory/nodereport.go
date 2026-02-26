package nodereportfactory

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

type NodeReportNotImplemented struct{}

func (n NodeReportNotImplemented) Get() []v1.Node {
	return nil
}

func (n NodeReportNotImplemented) GetByName(name string) *v1.Node {
	return nil
}

func (n NodeReportNotImplemented) GetByUid(uid string) *v1.Node {
	return nil
}

func (n NodeReportNotImplemented) GetByHostname(hostname string) *v1.Node {
	return nil
}

func (n NodeReportNotImplemented) GetByMachineProvider(machineProvider providermodels.ProviderType) []v1.Node {
	return nil
}

type NodeReport struct {
	nodes []v1.Node
}

func (n NodeReport) Get() []v1.Node {
	return n.nodes
}

func (n NodeReport) GetByName(name string) *v1.Node {
	for _, node := range n.nodes {
		if node.Name == name {
			return &node
		}
	}
	return nil
}

func (n NodeReport) GetByUid(uid string) *v1.Node {
	for _, node := range n.nodes {
		if string(node.UID) == uid {
			return &node
		}
	}
	return nil
}

func (n NodeReport) GetByHostname(hostname string) *v1.Node {
	for _, node := range n.nodes {
		if node.Labels["kubernetes.io/hostname"] == hostname {
			return &node
		}
	}
	return nil
}

func (n NodeReport) GetByMachineProvider(machineProvider providermodels.ProviderType) []v1.Node {
	var nodes []v1.Node
	for _, node := range n.nodes {
		if node.Labels["machineProvider"] == string(machineProvider) {
			nodes = append(nodes, node)
		}
	}
	return nodes
}

func NewClusterNodeReport(nodes []v1.Node) interregatortypes.ClusterNodeReport {
	if len(nodes) == 0 {
		return &NodeReportNotImplemented{}
	}
	return &NodeReport{nodes: nodes}
}
