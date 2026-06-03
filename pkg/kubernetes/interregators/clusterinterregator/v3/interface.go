package clusterinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/factories/interregatorfactory/v3"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v3"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/azure/azureproviderinterregator/v3"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/gke/gkeproviderinterregator/v3"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/k3d/k3dproviderinterregator/v3"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/kind/kindproviderinterregator/v3"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/talos/talosproviderinterregator/v3"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/tanzu/tanzuproviderinterregator/v3"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/unknown/unknownproviderinterregator/v3"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/vitistack/vitistackinterregator/v3"
	"k8s.io/client-go/kubernetes"
)

var interregators = []interregatortypes.ProviderDetectFunc{
	vitistackinterregator.Detect,
	tanzuproviderinterregator.Detect,
	talosproviderinterregator.Detect,
	kindproviderinterregator.Detect,
	k3dproviderinterregator.Detect,
	gkeproviderinterregator.Detect,
	azureproviderinterregator.Detect,
}

func NewClusterInterregatorFromKubernetesClient(client *kubernetes.Clientset) interregatortypes.ClusterInterregator {
	return NewClusterInterregator(client)
}

func NewClusterInterregator(client *kubernetes.Clientset) interregatortypes.ClusterInterregator {
	if client == nil {
		return unknownproviderinterregator.NewInterregator()
	}
	for _, detect := range interregators {
		provider := detect(client)
		if provider != nil {
			return interregatorfactory.NewClusterInterregatorFactory(client, provider)
		}
	}

	return unknownproviderinterregator.NewInterregator()
}
