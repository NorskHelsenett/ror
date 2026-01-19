package clusterinterregator

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/azure/azureproviderinterregator/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/gke/gkeproviderinterregator/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/k3d/k3dproviderinterregator/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/kind/kindproviderinterregator/v2"
	talosproviderinterregator "github.com/NorskHelsenett/ror/pkg/kubernetes/providers/talos/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/tanzu/tanzuproviderinterregator/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/unknown/unknownproviderinterregator/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/vitistack/vitistackinterregator/v2"
	v1core "k8s.io/api/core/v1"
	v1meta "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// Ensure interfaces are implemented
var (
	_ interregatortypes.ClusterInterregator = (*vitistackinterregator.Vitistacktypes)(nil)
	_ interregatortypes.ClusterInterregator = (*talosproviderinterregator.Talostypes)(nil)
	_ interregatortypes.ClusterInterregator = (*tanzuproviderinterregator.TanzuProviderinterregator)(nil)
	_ interregatortypes.ClusterInterregator = (*kindproviderinterregator.Kindtypes)(nil)
	_ interregatortypes.ClusterInterregator = (*k3dproviderinterregator.K3dtypes)(nil)
	_ interregatortypes.ClusterInterregator = (*gkeproviderinterregator.Gketypes)(nil)
	_ interregatortypes.ClusterInterregator = (*unknownproviderinterregator.UnknownProviderinterregator)(nil)
	_ interregatortypes.ClusterInterregator = (*azureproviderinterregator.Azuretypes)(nil)
)

type ClusterProviderInterregator interface {
	NewInterregator([]v1core.Node) interregatortypes.ClusterInterregator
}

var interregators = []interregatortypes.ClusterProviderInterregator{
	vitistackinterregator.Interregator{},
	tanzuproviderinterregator.Interregator{},
	talosproviderinterregator.Interregator{},
	kindproviderinterregator.Interregator{},
	k3dproviderinterregator.Interregator{},
	gkeproviderinterregator.Interregator{},
	azureproviderinterregator.Interregator{},
}

func NewClusterInterregatorFromKubernetesClient(client *kubernetes.Clientset) interregatortypes.ClusterInterregator {
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), v1meta.ListOptions{})
	if err != nil {
		return unknownproviderinterregator.NewInterregator()
	}

	return NewClusterInterregator(nodes.Items)
}

func NewClusterInterregator(nodes []v1core.Node) interregatortypes.ClusterInterregator {
	for _, inter := range interregators {
		interregator := inter.NewInterregator(nodes)
		if interregator != nil {
			return interregator
		}
	}

	return unknownproviderinterregator.NewInterregator()

}
