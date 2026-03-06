package restclientv2self

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2self"
)

func (c *V2Client) CreateOrUpdateApiKey(ctx context.Context, name string, ttl int64) (string, error) {
	resoponse := apicontractsv2self.CreateOrRenewApikeyResponse{}
	request := apicontractsv2self.CreateOrRenewApikeyRequest{
		Name: name,
		Ttl:  ttl,
	}
	err := c.Client.PostJSON(ctx, c.basePath+"/apikeys", request, &resoponse)
	if err != nil {
		return "", err
	}

	return resoponse.Token, nil
}
