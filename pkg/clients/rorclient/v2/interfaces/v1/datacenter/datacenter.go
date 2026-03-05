package datacenter

import "github.com/NorskHelsenett/ror/pkg/apicontracts"

type DatacenterInterface interface {
	Get() (*[]apicontracts.Datacenter, error)
	GetById(id string) (*apicontracts.Datacenter, error)
	GetByName(name string) (*apicontracts.Datacenter, error)
	Post(data apicontracts.DatacenterModel) (*apicontracts.Datacenter, error)
	Put(id string, data apicontracts.DatacenterModel) (*apicontracts.Datacenter, error)
}
