package clusterinterregator

import (
	"fmt"

	kubernetesclient "github.com/NorskHelsenett/ror/pkg/clients/kubernetes"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providerinterregator"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providerinterregator/types"
)

func InterregateCluster(kc kubernetesclient.K8SClientInterface) (*types.InterregationReport, error) {
	// Get all the nodes in the cluster
	nodes, err := kc.GetNodes()
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 || nodes == nil {
		return nil, fmt.Errorf("no nodes found")
	}
	return providerinterregator.NewInterregationReport(nodes)
}
