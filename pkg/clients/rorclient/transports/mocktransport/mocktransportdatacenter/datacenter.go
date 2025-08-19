package mocktransportdatacenter

import (
	"errors"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type V1Client struct{}

func NewV1Client() *V1Client {
	return &V1Client{}
}

func (c *V1Client) Get() (*[]apicontracts.Datacenter, error) {
	datacenters := []apicontracts.Datacenter{
		{
			ID:   "mock-dc-001",
			Name: "Mock Datacenter 1",
		},
		{
			ID:   "mock-dc-002",
			Name: "Mock Datacenter 2",
		},
	}
	return &datacenters, nil
}

func (c *V1Client) GetById(id string) (*apicontracts.Datacenter, error) {
	if id == "" {
		return nil, errors.New("datacenter ID cannot be empty")
	}

	datacenter := &apicontracts.Datacenter{
		ID:   id,
		Name: "Mock Datacenter " + id,
	}
	return datacenter, nil
}

func (c *V1Client) GetByName(name string) (*apicontracts.Datacenter, error) {
	if name == "" {
		return nil, errors.New("datacenter name cannot be empty")
	}

	datacenter := &apicontracts.Datacenter{
		ID:   "mock-dc-by-name",
		Name: name,
	}
	return datacenter, nil
}

func (c *V1Client) Post(data apicontracts.DatacenterModel) (*apicontracts.Datacenter, error) {
	if data.Name == "" {
		return nil, errors.New("datacenter name cannot be empty")
	}

	datacenter := &apicontracts.Datacenter{
		ID:   "mock-dc-new-001",
		Name: data.Name,
	}
	return datacenter, nil
}

func (c *V1Client) Put(id string, data apicontracts.DatacenterModel) (*apicontracts.Datacenter, error) {
	if id == "" {
		return nil, errors.New("datacenter ID cannot be empty")
	}
	if data.Name == "" {
		return nil, errors.New("datacenter name cannot be empty")
	}

	datacenter := &apicontracts.Datacenter{
		ID:   id,
		Name: data.Name,
	}
	return datacenter, nil
}
