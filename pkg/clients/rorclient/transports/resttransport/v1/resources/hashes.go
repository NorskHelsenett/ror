package resources

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts/apiresourcecontracts"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
)

func (c *V1Client) GetHashList(ownerref rortypes.RorResourceOwnerReference) (apiresourcecontracts.HashList, error) {
	var hashList apiresourcecontracts.HashList
	params := []httpclient.HttpTransportClientParams{
		{Key: httpclient.HttpTransportClientOptsQuery, Value: ownerref.GetQueryParams()},
	}

	err := c.Client.GetJSON(c.basePath+"/hashes", &hashList, params...)
	if err != nil {
		return hashList, err
	}

	return hashList, nil
}
