package token

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
)

type ExchangeTokenRequest struct {
	ClusterID string `json:"clusterId" validate:"required"`
	Admin     bool   `json:"admin,omitempty"`
	Token     string `json:"token" validate:"required"`
}

type V2Client struct {
	Client   *httpclient.HttpTransportClient
	basePath string
}

func NewV2Client(client *httpclient.HttpTransportClient) *V2Client {
	return &V2Client{
		Client:   client,
		basePath: "/v2/token",
	}
}
func (c *V2Client) Exchange(ctx context.Context, token string, clusterId string, admin bool) (string, error) {
	var tokendata string
	inn := ExchangeTokenRequest{
		ClusterID: clusterId,
		Admin:     admin,
		Token:     token,
	}

	tokendata = inn.Token

	err := c.Client.PostJSONWithContext(ctx, c.basePath+"/exchange", inn, &tokendata)
	if err != nil {
		return "", err
	}

	return tokendata, nil
}
