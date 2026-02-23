package resources

import (
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/helpers/resourcecache/resourcecachehashlist"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
)

func (c *V1Client) GetHashList(ownerref rorresourceowner.RorResourceOwnerReference) (resourcecachehashlist.HashList, error) {
	var hashList resourcecachehashlist.HashList
	params := []httpclient.HttpTransportClientParams{
		{Key: httpclient.HttpTransportClientOptsQuery, Value: ownerref.GetQueryParams()},
	}

	err := c.Client.GetJSON(c.basePath+"/hashes", &hashList, params...)
	if err != nil {
		return hashList, err
	}

	return hashList, nil
}
