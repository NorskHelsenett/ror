package apikeys

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apikeystypes/v2"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
)

// implements ApiKeysInterface pkg/clients/rorclient/v2/apikeys
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

func (c *V2Client) RegisterAgent(data apikeystypes.RegisterClusterRequest) (apikeystypes.RegisterClusterResponse, error) {
	var registerdata apikeystypes.RegisterClusterResponse
	err := c.Client.PostJSON(c.basePath+"/register/agent", data, &registerdata)
	if err != nil {
		return registerdata, err
	}

	return registerdata, nil
}
