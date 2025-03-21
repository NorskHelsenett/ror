package rortypes

import "github.com/NorskHelsenett/ror/pkg/models/providers"

// K8s deployment struct
type ResourceDatacenter struct {
	Spec   ResourceDatacenterSpec   `json:"spec"`
	Status ResourceDatacenterStatus `json:"status"`
	Legacy Datacenter               `json:"legacy"`
}

type ResourceDatacenterSpec struct {
	Workspaces []ResourceWorkspace `json:"workspaces"`
}

type ResourceDatacenterStatus struct {
	Workspaces  []ResourceWorkspace `json:"workspaces"`
	Location    DatacenterLocation  `json:"location"`
	APIEndpoint string              `json:"apiEndpoint"`
}

type Datacenter struct {
	ID          string                 `json:"id" bson:"_id,omitempty"`
	Name        string                 `json:"name"`
	Provider    providers.ProviderType `json:"provider"`
	Location    DatacenterLocation     `json:"location"`
	APIEndpoint string                 `json:"apiEndpoint"`
}

type DatacenterLocation struct {
	ID      string `json:"id" bson:"id,omitempty"`
	Region  string `json:"region"`
	Country string `json:"country"`
}

type Workspace struct {
	ID           string     `json:"id" bson:"_id,omitempty"`
	Name         string     `json:"name"`
	DatacenterID string     `json:"datacenterId" bson:"datacenterid,omitempty"`
	Datacenter   Datacenter `json:"datacenter"`
}
