package metrics

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
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

func (c *V1Client) CreatePVC(input apicontracts.PersistentVolumeClaimMetric) error {
	var dummy interface{}
	return c.Client.PostJSON(c.basePath+"/pvc", input, &dummy)
}

func (c *V1Client) PostReport(ctx context.Context, metricsReport apicontracts.MetricsReport) error {
	return c.Client.PostJSONWithContext(ctx, c.basePath, metricsReport, nil)
}
