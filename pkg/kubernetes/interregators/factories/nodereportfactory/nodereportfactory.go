package nodereportfactory

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	v1 "k8s.io/api/core/v1"
)

func NewNodeReportFactory(nodes []v1.Node) interregatortypes.ClusterNodeReport {
	return ClusterNodeReportFactory{
		nodereport: NewClusterNodeReport(nodes),
	}
}

type ClusterNodeReportFactory struct {
	nodereport               interregatortypes.ClusterNodeReport
	GetFunc                  func() []v1.Node
	GetByNameFunc            func(name string) *v1.Node
	GetByUidFunc             func(uid string) *v1.Node
	GetByHostnameFunc        func(hostname string) *v1.Node
	GetByMachineProviderFunc func(machineProvider providermodels.ProviderType) []v1.Node
}

func (f ClusterNodeReportFactory) Get() []v1.Node {
	if f.GetFunc != nil {
		return f.GetFunc()
	}
	return f.nodereport.Get()
}

func (f ClusterNodeReportFactory) GetByName(name string) *v1.Node {
	if f.GetByNameFunc != nil {
		return f.GetByNameFunc(name)
	}
	return f.nodereport.GetByName(name)
}

func (f ClusterNodeReportFactory) GetByUid(uid string) *v1.Node {
	if f.GetByUidFunc != nil {
		return f.GetByUidFunc(uid)
	}
	return f.nodereport.GetByUid(uid)
}

func (f ClusterNodeReportFactory) GetByHostname(hostname string) *v1.Node {
	if f.GetByHostnameFunc != nil {
		return f.GetByHostnameFunc(hostname)
	}
	return f.nodereport.GetByHostname(hostname)
}

func (f ClusterNodeReportFactory) GetByMachineProvider(machineProvider providermodels.ProviderType) []v1.Node {
	if f.GetByMachineProviderFunc != nil {
		return f.GetByMachineProviderFunc(machineProvider)
	}
	return f.nodereport.GetByMachineProvider(machineProvider)
}
