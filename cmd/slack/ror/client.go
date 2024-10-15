package ror

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"github.com/NorskHelsenett/ror/pkg/config/rorversion"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpauthprovider"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
	"github.com/spf13/viper"
)

var (
	Client *rorclient.RorClient
)

func SetupRORClient() {
	authProvider := httpauthprovider.NewAuthProvider(httpauthprovider.AuthPoviderTypeAPIKey, viper.GetString(configconsts.API_KEY))
	clientConfig := httpclient.HttpTransportClientConfig{
		BaseURL:      viper.GetString(configconsts.API_ENDPOINT),
		AuthProvider: authProvider,
		Version: rorversion.RorVersion{
			Version: "v1",
			Commit:  "ror-ms-slack",
		},
		Role: viper.GetString(configconsts.ROLE),
	}
	transport := resttransport.NewRorHttpTransport(&clientConfig)
	Client = rorclient.NewRorClient(transport)
}
