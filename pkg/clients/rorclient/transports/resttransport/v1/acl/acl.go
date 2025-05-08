package acl

import (
	"context"
	"net/url"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
)

type V1Client struct {
	Client   *httpclient.HttpTransportClient
	BasePath string
}

func NewV1Client(client *httpclient.HttpTransportClient) *V1Client {
	return &V1Client{
		Client:   client,
		BasePath: "/v1/acl",
	}
}

func (c V1Client) GetById(ctx context.Context, id string) (*aclmodels.AclV2ListItem, error) {
	url, err := url.Parse(c.BasePath)
	if err != nil {
		return nil, err
	}

	url = url.JoinPath(id)

	var acl aclmodels.AclV2ListItem

	err = c.Client.GetJSONWithContext(ctx, url.String(), &acl)
	if err != nil {
		return nil, err
	}

	return &acl, nil
}

func (c V1Client) GetByFilter(ctx context.Context, filter apicontracts.Filter) (*apicontracts.PaginatedResult[aclmodels.AclV2ListItem], error) {
	url, err := url.Parse(c.BasePath)
	if err != nil {
		return nil, err
	}

	url = url.JoinPath("filter")

	var res apicontracts.PaginatedResult[aclmodels.AclV2ListItem]

	err = c.Client.PostJSONWithContext(ctx, url.String(), filter, &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c V1Client) Create(ctx context.Context, item aclmodels.AclV2ListItem) error {
	url, err := url.Parse(c.BasePath)
	if err != nil {
		return err
	}

	var res aclmodels.AclV2ListItem

	err = c.Client.PostJSONWithContext(ctx, url.String(), item, &res)
	if err != nil {
		return err
	}

	return nil
}

func (c V1Client) Update(ctx context.Context, id string, item aclmodels.AclV2ListItem) error {

	return nil
}

func (c V1Client) Delete(ctx context.Context, id string) error {
	url, err := url.Parse(c.BasePath)
	if err != nil {
		return err
	}

	url = url.JoinPath(id)

	var res bool
	err = c.Client.DeleteWithContext(ctx, url.String(), res)
	if err != nil {
		return err
	}

	return nil
}

func (c V1Client) CheckAccess(ctx context.Context, scope, subject, access string) bool {
	return false
}
