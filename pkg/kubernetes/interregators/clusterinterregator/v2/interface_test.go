package clusterinterregator

import (
	"testing"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/azure/azureproviderinterregator/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/gke/gkeproviderinterregator/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/k3d/k3dproviderinterregator/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/kind/kindproviderinterregator/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/talos/talosproviderinterregator/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/tanzu/tanzuproviderinterregator/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/vitistack/vitistackinterregator/v2"

	"github.com/stretchr/testify/assert"
	v1core "k8s.io/api/core/v1"
)

func TestNewClusterInterregator_EmptyNodes_DoesNotPanic(t *testing.T) {
	assert.NotPanics(t, func() {
		_ = NewClusterInterregator([]v1core.Node{})
	})
}

func TestProviderInterregators_EmptyNodes_NewInterregator_DoesNotPanic(t *testing.T) {
	testCases := []struct {
		name string
		fn   func([]v1core.Node)
	}{
		{
			name: "vitistack",
			fn:   func(nodes []v1core.Node) { _ = vitistackinterregator.Interregator{}.NewInterregator(nodes) },
		},
		{
			name: "tanzu",
			fn:   func(nodes []v1core.Node) { _ = tanzuproviderinterregator.Interregator{}.NewInterregator(nodes) },
		},
		{
			name: "talos",
			fn:   func(nodes []v1core.Node) { _ = talosproviderinterregator.Interregator{}.NewInterregator(nodes) },
		},
		{
			name: "kind",
			fn:   func(nodes []v1core.Node) { _ = kindproviderinterregator.Interregator{}.NewInterregator(nodes) },
		},
		{
			name: "k3d",
			fn:   func(nodes []v1core.Node) { _ = k3dproviderinterregator.Interregator{}.NewInterregator(nodes) },
		},
		{
			name: "gke",
			fn:   func(nodes []v1core.Node) { _ = gkeproviderinterregator.Interregator{}.NewInterregator(nodes) },
		},
		{
			name: "azure",
			fn:   func(nodes []v1core.Node) { _ = azureproviderinterregator.Interregator{}.NewInterregator(nodes) },
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			assert.NotPanics(t, func() {
				tc.fn([]v1core.Node{})
			})
		})
	}
}
