package resources

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"

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

func (c *V2Client) Get(query rorresources.ResourceQuery) (rorresources.ResourceSet, error) {
	var res rorresources.ResourceSet
	jsonQuery, err := json.Marshal(&query)
	if err != nil {
		return res, err
	}
	queryString := base64.StdEncoding.EncodeToString(jsonQuery)
	err = c.Client.GetJSON(c.basePath+"?query="+queryString, &res)
	return res, err
}

func (c *V2Client) Update(res *rorresources.ResourceSet) (*rorresources.ResourceUpdateResults, error) {
	var ret *rorresources.ResourceUpdateResults
	err := c.Client.PostJSON(c.basePath, res, &ret)
	if err != nil {
		return nil, err
	}

	return ret, nil
}

func (c *V2Client) Delete(uid string) (*rorresources.ResourceUpdateResults, error) {
	var out rorresources.ResourceUpdateResults

	url, err := url.JoinPath(c.basePath, "uid", uid)
	if err != nil {
		return nil, fmt.Errorf("could not create url: %w", err)
	}

	err = c.Client.Delete(url, &out)
	if err != nil {
		return nil, err
	}

	return &out, nil
}

func (c *V2Client) Exists(uid string) (bool, error) {
	url, err := url.JoinPath(c.basePath, "uid", uid)
	if err != nil {
		return false, fmt.Errorf("could not create url: %w", err)
	}

	_, status, err := c.Client.Head(url)
	if err != nil {
		return false, err
	}

	if status == http.StatusNoContent {
		return true, nil
	}

	return false, nil
}

func (c *V2Client) GetOwnHashes(clusterId string) (apicontractsv2resources.HashList, error) {
	var hashList apicontractsv2resources.HashList
	params := httpclient.HttpTransportClientParams{
		Key: httpclient.HttpTransportClientOptsQuery,
		Value: map[string]string{
			"ownerScope":   string(aclmodels.Acl2ScopeCluster),
			"ownerSubject": clusterId,
		},
	}
	err := c.Client.GetJSON(c.basePath+"/hashes", &hashList, params)
	if err != nil {
		return hashList, err
	}

	return hashList, nil
}
