package acl

import (
	"context"
	"net/url"
	"strings"

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

	u = u.JoinPath("check", scope.String(), subject.String(), access.String())

	err = c.Client.GetJSON(ctx, u.String(), nil)
	return err == nil
}
