package resources

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2resources"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
)

type ResourcesInterface interface {
	Get(query rorresources.ResourceQuery) (rorresources.ResourceSet, error)
	Update(res *rorresources.ResourceSet) (*rorresources.ResourceUpdateResults, error)
	Delete(uid string) (*rorresources.ResourceUpdateResults, error)
	Exists(uid string) (bool, error)
	GetOwnHashes(clusterId string) (apicontractsv2resources.HashList, error)
}
