package restclientv2self

import "github.com/NorskHelsenett/ror/pkg/apicontracts/v2/apicontractsv2self"

func (c *V2Client) CreateOrUpdateApiKey(name string, ttl int64) (string, error) {
	resoponse := apicontractsv2self.CreateOrRenewApikeyResponse{}
	request := apicontractsv2self.CreateOrRenewApikeyRequest{
		Name: name,
		Ttl:  ttl,
	}
	err := c.Client.PostJSON(c.basePath+"/apikeys", request, &resoponse)
	if err != nil {
		return "", err
	}

	return resoponse.Token, nil
}
