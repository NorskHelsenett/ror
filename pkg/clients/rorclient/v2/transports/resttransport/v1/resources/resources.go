package resources

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpclient"
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

func (c *V1Client) Create(ctx context.Context, resourceUpdate *apiresourcecontracts.ResourceUpdateModel) error {
	err := c.Client.PostJSON(ctx, c.basePath, resourceUpdate, nil)
	return err
}

func (c *V1Client) Update(ctx context.Context, resourceUpdate *apiresourcecontracts.ResourceUpdateModel) error {
	err := c.Client.PutJSON(ctx, c.basePath+"/uid/"+resourceUpdate.Uid, resourceUpdate, nil)
	return err
}

func (c *V1Client) Delete(ctx context.Context, uid string) error {
	err := c.Client.Delete(ctx, c.basePath+"/uid/"+uid, nil)
	return err
}
