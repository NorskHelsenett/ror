package acl

import (
	"context"
	"fmt"
	"mime"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclsnapshot"
)

func (c V2Client) GetSnapshot(ctx context.Context, etag string, limits aclsnapshot.Limits) (*aclsnapshot.FetchResponse, error) {
	if err := limits.Validate(); err != nil {
		return nil, err
	}
	headers := map[string]string{"Accept": aclsnapshot.ContentType}
	if etag != "" {
		headers["If-None-Match"] = etag
	}
	res, err := c.Client.GetRaw(ctx, c.BasePath+"/snapshot",
		httpclient.HttpTransportClientParams{Key: httpclient.HttpTransportClientOptsHeaders, Value: headers},
		httpclient.HttpTransportClientParams{Key: httpclient.HttpTransportClientOptsMaxBytes, Value: int64(limits.MaxEnvelopeBytes)})
	if err != nil {
		return nil, err
	}
	responseETag := res.Header.Get("ETag")
	if responseETag == "" {
		return nil, fmt.Errorf("snapshot response is missing ETag")
	}
	switch res.StatusCode {
	case http.StatusNotModified:
		return &aclsnapshot.FetchResponse{ETag: responseETag, NotModified: true}, nil
	case http.StatusOK:
		mediaType, _, err := mime.ParseMediaType(res.Header.Get("Content-Type"))
		if err != nil || mediaType != aclsnapshot.ContentType {
			return nil, fmt.Errorf("unexpected snapshot content type %q", res.Header.Get("Content-Type"))
		}
		return &aclsnapshot.FetchResponse{Envelope: res.Body, ETag: responseETag}, nil
	default:
		return nil, fmt.Errorf("unexpected snapshot status %d", res.StatusCode)
	}
}
