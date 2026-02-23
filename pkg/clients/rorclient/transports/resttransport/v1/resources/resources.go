package resources

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
)

type V1Client struct {
	Client   *httpclient.HttpTransportClient
	basePath string
}

func NewV1Client(client *httpclient.HttpTransportClient) *V1Client {
	return &V1Client{
		Client:   client,
		basePath: "/v1/resources",
	}
}

func (c *V1Client) Create(resourceUpdate *apiresourcecontracts.ResourceUpdateModel) error {
	var ret any
	err := c.Client.PostJSON(c.basePath, resourceUpdate, &ret)
	return err
}

func (c *V1Client) Update(resourceUpdate *apiresourcecontracts.ResourceUpdateModel) error {
	var ret any
	err := c.Client.PutJSON(c.basePath+"/uid/"+resourceUpdate.Uid, resourceUpdate, &ret)
	return err
}

func (c *V1Client) Delete(uid string) error {
	var ret any
	err := c.Client.Delete(c.basePath+"/uid/"+uid, &ret)
	return err
}
