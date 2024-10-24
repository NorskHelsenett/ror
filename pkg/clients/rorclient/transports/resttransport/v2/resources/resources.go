package resources

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2resources"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"

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
	//url, err := url.JoinPath(c.basePath, "uid", uid)
	//if err != nil {
	//	return nil, fmt.Errorf("could not create url: %w", err)
	//}

	// we need ta add a head request to client
	//err = c.Client.(url, &out)
	//if err != nil {
	//	return nil, err
	//}

	return false, fmt.Errorf("not implemented")
}

func (c *V2Client) GetOwnHashes() (apicontractsv2resources.HashList, error) {
	var hashList apicontractsv2resources.HashList
	err := c.Client.GetJSON(c.basePath+"/self/hashes", &hashList)
	if err != nil {
		return hashList, err
	}

	return hashList, nil
}
