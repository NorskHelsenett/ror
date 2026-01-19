package clusterinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/vitistack/vitistackinterregator/v2"
	v1 "k8s.io/api/core/v1"
)

var (
	_ interregatortypes.ClusterInterregator = (*vitistackinterregator.Vitistacktypes)(nil)
	// _ ClusterInterregator = (*tanzuproviderinterregator.TanzuTypes)(nil)
	// _ ClusterInterregator = (*azureproviderinterregator.AzureTypes)(nil)
	// _ ClusterInterregator = (*k3dproviderinterregator.K3dTypes)(nil)
	// _ ClusterInterregator = (*kindproviderinterregator.KindTypes)(nil)
	// _ ClusterInterregator = (*gkeproviderinterregator.GkeTypes)(nil)
	// _ ClusterInterregator = (*talosproviderinterregator.TalosTypes)(nil)
)

type ClusterProviderInterregator interface {
	NewInterregator([]v1.Node) interregatortypes.ClusterInterregator
}

var interregators = []interregatortypes.ClusterProviderInterregator{
	vitistackinterregator.Interregator{},
}
