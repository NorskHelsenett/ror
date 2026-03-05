package apikeys

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/apikeystypes/v2"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpclient"
)

// implements ApiKeysInterface pkg/clients/rorclient/v2/interfaces/v2/apikeys
type V2Client struct {
	Client   *httpclient.HttpTransportClient
	basePath string
}

func NewV2Client(client *httpclient.HttpTransportClient) *V2Client {
	return &V2Client{
		Client:   client,
		basePath: "/v2/apikeys",
	}
}

func (c *V2Client) RegisterAgent(ctx context.Context, data apikeystypes.RegisterClusterRequest) (apikeystypes.RegisterClusterResponse, error) {
	var registerdata apikeystypes.RegisterClusterResponse
	err := c.Client.PostJSON(ctx, c.basePath+"/register/agent", data, &registerdata)
	if err != nil {
		return registerdata, err
	}

	return registerdata, nil
}
