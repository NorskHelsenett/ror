package info

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/config/rorconfig"
	"github.com/NorskHelsenett/ror/pkg/config/rorversion"
	"go.opentelemetry.io/otel"
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
	ctx, span := otel.GetTracerProvider().Tracer(rorconfig.GetString(rorconfig.TRACER_ID)).Start(ctx, "info.GetVersion")
	defer span.End()
	var versiondata rorversion.RorVersion

	err := c.Client.GetJSON(ctx, c.basePath+"/version", &versiondata, httpclient.HttpTransportClientParams{Key: httpclient.HttpTransportClientOptsNoAuth})
	if err != nil {
		return "", err
	}

	return versiondata.GetVersionWithCommit(), nil
}
