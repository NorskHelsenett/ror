package nodeservice

import (
	"context"
	"errors"
	"fmt"
	"github.com/NorskHelsenett/ror/internal/kubernetes/k8smodels"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/providerinterregationreport"
	"github.com/NorskHelsenett/ror/pkg/models/providers"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
	metrics "k8s.io/metrics/pkg/client/clientset/versioned"
)

func GetNodes(k8sClient *kubernetes.Clientset, metricsClient *metrics.Clientset) ([]k8smodels.Node, error) {
	var nodes []k8smodels.Node
	list, err := k8sClient.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
	if err != nil {
		rlog.Error("could not get node list", err)
		return nil, errors.New("could not extract node list ")
	}
	report, err := providerinterregationreport.NewInterregationReport(list.Items)
	if err != nil {
		rlog.Error("could not get provider type", err)
	}
	rlog.Debug(fmt.Sprintf("Provider detected: %s", report.GetProvider()))

	for _, node := range list.Items {

		n := k8smodels.Node{}

		n.Workspace = report.Workspace
		n.ClusterName = report.ClusterName
		n.Datacenter = report.Datacenter
		n.Provider = report.Provider

		n.OsImage = node.Status.NodeInfo.OSImage
		n.Created = node.CreationTimestamp.Time
		n.Annotations = node.Annotations
		n.Name = node.Name
		n.Labels = node.Labels

		switch report.GetProvider() {
		case providers.ProviderTypeTanzu:
			fillNodeTanzu(&n)
		case providers.ProviderTypeAks:
			fillNodeAzure(&n)
		case providers.ProviderTypeK3d:
			fillNodeK3d(&n)
		case providers.ProviderTypeKind:
			fillNodeKind(&n)
		case providers.ProviderTypeGke:
			fillNodeGke(&n)
		case providers.ProviderTypeTalos:
			fillNodeTalos(&n)
		default:
			fillNodeDefault(&n)
		}

		n.Architecture = node.Status.NodeInfo.Architecture
		n.ContainerRuntimeVersion = node.Status.NodeInfo.ContainerRuntimeVersion
		n.KernelVersion = node.Status.NodeInfo.KernelVersion
		n.KubeProxyVersion = node.Status.NodeInfo.KubeProxyVersion
		n.KubeletVersion = node.Status.NodeInfo.KubeletVersion
		n.OperatingSystem = node.Status.NodeInfo.OperatingSystem

		nodeMetrics, err := metricsClient.MetricsV1beta1().NodeMetricses().Get(context.TODO(), node.Name, v1.GetOptions{})
		if err == nil {
			cpuUsage := nodeMetrics.Usage.Cpu()
			cpuAllocated, _ := node.Status.Allocatable.Cpu().AsInt64()

			memoryUsageInt, _ := nodeMetrics.Usage.Memory().AsInt64()
			memoryAllocated := node.Status.Allocatable.Memory().Value()

			n.Resources = apicontracts.NodeResources{
				Allocated: apicontracts.ResourceAllocated{
					Cpu:         cpuAllocated,
					MemoryBytes: memoryAllocated,
				},
				Consumed: apicontracts.ResourceConsumed{
					CpuMilliValue: cpuUsage.MilliValue(),
					MemoryBytes:   memoryUsageInt,
				},
			}
		} else {
			rlog.Debug("could not fetch node metrics", rlog.String("name", node.Name))
		}

		nodes = append(nodes, n)
	}

	return nodes, nil
}

func fillNodeTanzu(n *k8smodels.Node) {
	n.MachineName = n.Labels["kubernetes.io/hostname"]
	workspaceArray := strings.Split(n.Workspace, "-")
	if len(workspaceArray) > 0 {
		n.Datacenter = workspaceArray[0]
	}
}

func fillNodeAzure(n *k8smodels.Node) {
	n.MachineName = n.Labels["kubernetes.io/hostname"]
}

func fillNodeK3d(n *k8smodels.Node) {
	hostname := n.Labels["kubernetes.io/hostname"]

	n.MachineName = fmt.Sprintf("%s-%s", hostname, "localhost")
}
func fillNodeKind(n *k8smodels.Node) {
	hostname := n.Labels["kubernetes.io/hostname"]
	n.MachineName = fmt.Sprintf("%s-%s", hostname, "localhost")
}
func fillNodeGke(n *k8smodels.Node) {
	n.MachineName = n.Labels["kubernetes.io/hostname"]
}

func fillNodeTalos(n *k8smodels.Node) {
	n.MachineName = n.Labels["kubernetes.io/hostname"]
}

func fillNodeDefault(n *k8smodels.Node) {
	n.MachineName = "Unknown"
}
