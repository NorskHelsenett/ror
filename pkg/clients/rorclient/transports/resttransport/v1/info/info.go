package info

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/config/rorversion"
)

type V1Client struct {
	Client   *httpclient.HttpTransportClient
	basePath string
}

func NewV1Client(client *httpclient.HttpTransportClient) *V1Client {
	return &V1Client{
		Client:   client,
		basePath: "/v1/info",
	}
}

func (c *V1Client) GetVersion(ctx context.Context) (string, error) {
	var versiondata rorversion.RorVersion

	err := c.Client.GetJSONWithContext(ctx, c.basePath+"/version", &versiondata, httpclient.HttpTransportClientParams{Key: httpclient.HttpTransportClientOptsNoAuth})
	if err != nil {
		return "", err
	}

	return versiondata.GetVersionWithCommit(), nil
}
