package acl

import (
	"context"
	"net/http"
	"net/url"
	"strings"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclcaps"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
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

func (c V2Client) Lookup(ctx context.Context, access aclcaps.AccessTypeV3, scopes []aclscope.Scope, subjects []aclscope.Subject) (*aclmodels.AclV3LookupResponse, error) {
	u, err := url.Parse(c.BasePath)
	if err != nil {
		return nil, err
	}

	u = u.JoinPath("lookup")

	query := make(map[string]string)
	if access != "" {
		query["access"] = access.String()
	}
	if len(scopes) > 0 {
		stringscopes := make([]string, len(scopes))
		for i, s := range scopes {
			stringscopes[i] = s.String()
		}
		query["scope"] = strings.Join(stringscopes, ",")
	}
	if len(subjects) > 0 {
		stringsubjects := make([]string, len(subjects))
		for i, s := range subjects {
			stringsubjects[i] = s.String()
		}
		query["subject"] = strings.Join(stringsubjects, ",")
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

func (c V2Client) LookupByScopeSubject(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject) (*aclmodels.Acl3LookupByScopeSubjectResponse, error) {
	u, err := url.Parse(c.BasePath)
	if err != nil {
		return nil, err
	}

	u = u.JoinPath("lookup", scope.String(), subject.String())

	var res aclmodels.Acl3LookupByScopeSubjectResponse
	err = c.Client.GetJSON(ctx, u.String(), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
func (c V2Client) CheckAccess(ctx context.Context, scope aclscope.Scope, subject aclscope.Subject, access aclcaps.AccessTypeV3) bool {
	u, err := url.Parse(c.BasePath)
	if err != nil {
		return false
	}

	u = u.JoinPath("lookup", scope.String(), subject.String(), access.String())

	_, statusCode, err := c.Client.Head(ctx, u.String())
	return err == nil && statusCode == http.StatusOK
}

func (c V2Client) Create(ctx context.Context, item aclmodels.AclV3ListItem) (*aclmodels.AclV3ListItem, error) {
	u, err := url.Parse(c.BasePath)
	if err != nil {
		return nil, err
	}

	var created aclmodels.AclV3ListItem
	if err := c.Client.PostJSON(ctx, u.String(), item, &created); err != nil {
		return nil, err
	}
	return &created, nil
}

func (c V2Client) Update(ctx context.Context, id string, item aclmodels.AclV3ListItem) (*aclmodels.AclV3ListItem, error) {
	u, err := url.Parse(c.BasePath)
	if err != nil {
		return nil, err
	}

	u = u.JoinPath(id)

	var updated aclmodels.AclV3ListItem
	if err := c.Client.PutJSON(ctx, u.String(), item, &updated); err != nil {
		return nil, err
	}
	return &updated, nil
}

func (c V2Client) Delete(ctx context.Context, id string) error {
	u, err := url.Parse(c.BasePath)
	if err != nil {
		return err
	}

	u = u.JoinPath(id)

	var res bool
	return c.Client.Delete(ctx, u.String(), &res)
}

func (c V2Client) GetById(ctx context.Context, id string) (*aclmodels.AclV3ListItem, error) {
	u, err := url.Parse(c.BasePath)
	if err != nil {
		return nil, err
	}

	u = u.JoinPath(id)

	var item aclmodels.AclV3ListItem
	if err := c.Client.GetJSON(ctx, u.String(), &item); err != nil {
		return nil, err
	}
	return &item, nil
}

func (c V2Client) GetByFilter(ctx context.Context, filter apicontracts.Filter) (*apicontracts.PaginatedResult[aclmodels.AclV3ListItem], error) {
	u, err := url.Parse(c.BasePath)
	if err != nil {
		return nil, err
	}

	u = u.JoinPath("filter")

	var res apicontracts.PaginatedResult[aclmodels.AclV3ListItem]
	if err := c.Client.PostJSON(ctx, u.String(), filter, &res); err != nil {
		return nil, err
	}
	return &res, nil
}

// GetAll pages through GetByFilter (500 per page) and returns every entry.
func (c V2Client) GetAll(ctx context.Context) (*[]aclmodels.AclV3ListItem, error) {
	const paginationLimit = 500
	nextBatch := 0
	var acls []aclmodels.AclV3ListItem
	for {
		batch, err := c.GetByFilter(ctx, apicontracts.Filter{Limit: paginationLimit, Skip: nextBatch})
		if err != nil {
			return nil, err
		}
		if batch == nil || len(batch.Data) == 0 {
			return &acls, nil
		}
		acls = append(acls, batch.Data...)
		nextBatch += paginationLimit
	}
}
