package clusterinterregator

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/interregatortypes/v2"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/unknown/unknownproviderinterregator/v2"
	v1core "k8s.io/api/core/v1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/client-go/kubernetes"
)

// IsTypeOf(nodes []v1.Node) bool

func NewClusterInterregatorFromKubernetesClient(client *kubernetes.Clientset) interregatortypes.ClusterInterregator {
	nodes, err := client.CoreV1().Nodes().List(context.TODO(), v1.ListOptions{})
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
