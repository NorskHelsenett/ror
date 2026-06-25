package acl

import (
	"context"
	"net/url"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

type V2Client struct {
	Client   *httpclient.HttpTransportClient
	BasePath string
}

func NewV2Client(client *httpclient.HttpTransportClient) *V2Client {
	return &V2Client{
		Client:   client,
		BasePath: "/v2/acl",
	}
}

func (c V2Client) Lookup(ctx context.Context, access string, scopes []string, subjects []string) (*aclmodels.AclV3LookupResponse, error) {
	u, err := url.Parse(c.BasePath)
	if err != nil {
		return nil, err
	}

	u = u.JoinPath("lookup")

	query := make(map[string]string)
	if access != "" {
		query["access"] = access
	}
	if len(scopes) > 0 {
		query["scope"] = strings.Join(scopes, ",")
	}
	if len(subjects) > 0 {
		query["subject"] = strings.Join(subjects, ",")
	}

	var res aclmodels.AclV3LookupResponse
	err = c.Client.GetJSON(ctx, u.String(), &res, httpclient.HttpTransportClientParams{
		Key:   httpclient.HttpTransportClientOptsQuery,
		Value: query,
	})
	if err != nil {
		return nil, err
	}

	return &res, nil
}
