package resources

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strconv"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2resources"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"

	"github.com/NorskHelsenett/ror/pkg/rorresources"
)

type V2Client struct {
	Client   *httpclient.HttpTransportClient
	basePath string
}

func NewV2Client(client *httpclient.HttpTransportClient) *V2Client {
	return &V2Client{
		Client:   client,
		basePath: "/v2/resources",
	}
}

func (c *V2Client) Get(ctx context.Context, query rorresources.ResourceQuery) (rorresources.ResourceSet, error) {
	var res rorresources.ResourceSet
	jsonQuery, err := json.Marshal(&query)
	if err != nil {
		return res, err
	}
	queryString := base64.StdEncoding.EncodeToString(jsonQuery)
	err = c.Client.GetJSONWithContext(ctx, c.basePath+"?query="+queryString, &res)
	return res, err
}

func (c *V2Client) Update(ctx context.Context, res *rorresources.ResourceSet) (*rorresources.ResourceUpdateResults, error) {
	var ret *rorresources.ResourceUpdateResults
	err := c.Client.PostJSONWithContext(ctx, c.basePath, res, &ret, httpclient.HttpTransportClientParams{Key: httpclient.HttpTransportClientOptsHeaders, Value: map[string]string{"X-Resources-Count": strconv.Itoa(res.Len())}})
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *V2Client) Delete(ctx context.Context, uid string) (*rorresources.ResourceUpdateResults, error) {
	var out rorresources.ResourceUpdateResults

	uri, err := url.JoinPath(c.basePath, "uid", uid)
	if err != nil {
		return nil, fmt.Errorf("could not create url: %w", err)
	}

	err = c.Client.DeleteWithContext(ctx, uri, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *V2Client) Exists(ctx context.Context, uid string) (bool, error) {
	uri, err := url.JoinPath(c.basePath, "uid", uid)
	if err != nil {
		return false, fmt.Errorf("could not create url: %w", err)
	}

	_, status, err := c.Client.HeadWithContext(ctx, uri)
	if err != nil {
		return false, err
	}

	if status == http.StatusNoContent {
		return true, nil
	}

	return false, nil
}

func (c *V2Client) GetOwnHashes(ctx context.Context, clusterId string) (apicontractsv2resources.HashList, error) {
	var hashList apicontractsv2resources.HashList
	params := httpclient.HttpTransportClientParams{
		Key: httpclient.HttpTransportClientOptsQuery,
		Value: map[string]string{
			"ownerScope":   string(aclmodels.Acl2ScopeCluster),
			"ownerSubject": clusterId,
		},
	}
	err := c.Client.GetJSONWithContext(ctx, c.basePath+"/hashes", &hashList, params)
	if err != nil {
		return hashList, err
	}

	return hashList, nil
}
