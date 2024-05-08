package datacenter

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type V1Client struct {
	Client   *httpclient.HttpTransportClient
	basePath string
}

func NewV1Client(client *httpclient.HttpTransportClient) *V1Client {
	return &V1Client{
		Client:   client,
		basePath: "/v1/datacenters",
	}
}

func (c *V1Client) Get() (*[]apicontracts.Datacenter, error) {
	var datacenters []apicontracts.Datacenter

	err := c.Client.GetJSON(c.basePath, &datacenters)
	if err != nil {
		return nil, err
	}

	return &datacenters, nil
}

func (c *V1Client) GetById(id string) (*apicontracts.Datacenter, error) {
	var datacenter apicontracts.Datacenter
	err := c.Client.GetJSON(c.basePath+"/id/"+id, &datacenter)
	if err != nil {
		return nil, err
	}

	return &datacenter, nil
}

func (c *V1Client) GetByName(name string) (*apicontracts.Datacenter, error) {
	var datacenter apicontracts.Datacenter
	err := c.Client.GetJSON(c.basePath+"/"+name, &datacenter)
	if err != nil {
		return nil, err
	}

	return &datacenter, nil
}

func (c *V1Client) Post(data apicontracts.DatacenterModel) (*apicontracts.Datacenter, error) {
	// Implement the logic to update an existing datacenter
	return nil, fmt.Errorf("UpdateDatacenter: not implemented")
}

func (c *V1Client) Put(id string, data apicontracts.DatacenterModel) (*apicontracts.Datacenter, error) {
	// Implement the logic to delete a datacenter by ID
	return nil, fmt.Errorf("DeleteDatacenter: not implemented")
}
