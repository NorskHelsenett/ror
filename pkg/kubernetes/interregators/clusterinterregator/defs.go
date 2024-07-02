package clusterinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/azure/azureproviderinterregator"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/gke/gkeproviderinterregator"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/k3d/k3dproviderinterregator"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/kind/kindproviderinterregator"
	talosproviderinterregator "github.com/NorskHelsenett/ror/pkg/kubernetes/providers/talos"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/tanzu/tanzuproviderinterregator"
)

// GetProviderInterregators returns a list of all the provider interregators implemented
// TODO: move to a config provider?
var interregators = []ClusterProviderInterregator{
	tanzuproviderinterregator.NewInterregator(),
	azureproviderinterregator.NewInterregator(),
	k3dproviderinterregator.NewInterregator(),
	kindproviderinterregator.NewInterregator(),
	gkeproviderinterregator.NewInterregator(),
	talosproviderinterregator.NewInterregator(),
}
