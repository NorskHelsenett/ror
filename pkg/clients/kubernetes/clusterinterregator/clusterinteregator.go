package clusterinterregator

import (
	"fmt"

	kubernetesclient "github.com/NorskHelsenett/ror/pkg/clients/kubernetes"

	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/clusterinterregator"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/interregators/providerinterregationreport"
)

func InterregateCluster(kc kubernetesclient.K8SClientInterface) (*providerinterregationreport.InterregationReport, error) {
	// Get all the nodes in the cluster
	nodes, err := kc.GetNodes()
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 || nodes == nil {
		return nil, fmt.Errorf("no nodes found")
	}
	return providerinterregationreport.NewInterregationReport(nodes)
}

func GetInterregator(kc kubernetesclient.K8SClientInterface) (clusterinterregator.ClusterInterregator, error) {
	// Get all the nodes in the cluster
	nodes, err := kc.GetNodes()
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 || nodes == nil {
		return nil, fmt.Errorf("no nodes found")
	}
	return clusterinterregator.NewClusterInterregator(nodes), nil
}
