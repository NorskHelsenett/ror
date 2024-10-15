package clients

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var client *resty.Client

func GetOrCreateRorClient() (*resty.Client, error) {
	if client != nil {
		return client, nil
	}

	client = resty.New()
	client.SetBaseURL(viper.GetString(configconsts.API_ENDPOINT))
	apikey := viper.GetString(configconsts.API_KEY)
	client.Header.Add("X-API-KEY", apikey)

	return client, nil
}
