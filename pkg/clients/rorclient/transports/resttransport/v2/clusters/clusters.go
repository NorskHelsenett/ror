package clusters

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/clustersapi/v2"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
)

type V2Client struct {
	Client   *httpclient.HttpTransportClient
	basePath string
}

func NewV2Client(client *httpclient.HttpTransportClient) *V2Client {
	return &V2Client{
		Client:   client,
		basePath: "/v2/clusters",
	}
}

func (c *V2Client) Register(data clustersapi.RegisterClusterRequest) (clustersapi.RegisterClusterResponse, error) {
	var selfdata clustersapi.RegisterClusterResponse
	err := c.Client.GetJSON(c.basePath+"/self", &selfdata)
	if err != nil {
		return selfdata, err
	}

	return selfdata, nil
}
