package clusterinterregator

import (
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/azure/azureproviderinterregator"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/gke/gkeproviderinterregator"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/k3d/k3dproviderinterregator"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/kind/kindproviderinterregator"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/talos/talosproviderinterregator"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/tanzu/tanzuproviderinterregator"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/vitistack/vitistackinterregator"
)

// GetProviderInterregators returns a list of all the provider interregators implemented
// TODO: move to a config provider?
var interregators = []ClusterProviderInterregator{
	vitistackinterregator.NewInterregator(),
	tanzuproviderinterregator.NewInterregator(),
	azureproviderinterregator.NewInterregator(),
	k3dproviderinterregator.NewInterregator(),
	kindproviderinterregator.NewInterregator(),
	gkeproviderinterregator.NewInterregator(),
	talosproviderinterregator.NewInterregator(),
}
