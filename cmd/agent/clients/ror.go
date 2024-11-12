// The package implements clients for the ror-agent
package clients

import (
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var client *resty.Client
var clientNonAuth *resty.Client

func GetOrCreateRorClient() (*resty.Client, error) {
	if client != nil {
		return client, nil
	}

	client = resty.New()
	client.SetBaseURL(viper.GetString(configconsts.API_ENDPOINT))
	client.Header.Add("X-API-KEY", viper.GetString(configconsts.API_KEY))
	client.Header.Set("User-Agent", fmt.Sprintf("ROR-Agent/%s", viper.GetString(configconsts.VERSION)))

	return client, nil
}

func GetOrCreateRorClientNonAuth() (*resty.Client, error) {
	if clientNonAuth != nil {
		return clientNonAuth, nil
	}

	clientNonAuth = resty.New()
	clientNonAuth.SetBaseURL(viper.GetString(configconsts.API_ENDPOINT))
	clientNonAuth.Header.Set("User-Agent", fmt.Sprintf("ROR-Agent/%s", viper.GetString(configconsts.VERSION)))

	return clientNonAuth, nil
}
