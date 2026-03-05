package metrics

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpclient"
)

type V1Client struct {
	Client   *httpclient.HttpTransportClient
	basePath string
}

func NewV1Client(client *httpclient.HttpTransportClient) *V1Client {
	return &V1Client{
		Client:   client,
		basePath: "/v1/metrics",
	}
}

func (c *V1Client) CreatePVC(ctx context.Context, input apicontracts.PersistentVolumeClaimMetric) error {
	var dummy any
	return c.Client.PostJSON(ctx, c.basePath+"/pvc", input, &dummy)
}

func (c *V1Client) PostReport(ctx context.Context, metricsReport apicontracts.MetricsReport) error {
	var dummy any
	return c.Client.PostJSON(ctx, c.basePath, metricsReport, &dummy)
}
