package types

import (
	v1 "k8s.io/api/core/v1"
)

type InterregationReport struct {
	Nodes       []v1.Node
	Provider    string
	ClusterName string
	Workspace   string
	Datacenter  string
}

type ClusterProviderinterregator interface {
	IsOfType(*InterregationReport) bool
	GetProvider(*InterregationReport) string
	GetClusterName(*InterregationReport) string
	GetWorkspace(*InterregationReport) string
	GetDatacenter(*InterregationReport) string
}

func (c *InterregationReport) GetInterregator(interregators []ClusterProviderinterregator) ClusterProviderinterregator {
	for _, interregator := range interregators {
		if interregator.IsOfType(c) {
			return interregator
		}
	}
	return nil
}
