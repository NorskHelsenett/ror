package rorclient

import (
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpauthprovider"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"

	"github.com/NorskHelsenett/ror/pkg/config/rorversion"

	"github.com/spf13/viper"
)

var (
	// RorClient is the client used to communicate with the ROR API
	RorClient *rorclient.RorClient
)

func SetupRORClient() {
	authProvider := httpauthprovider.NewAuthProvider(httpauthprovider.AuthPoviderTypeAPIKey, viper.GetString(configconsts.API_KEY))
	clientConfig := httpclient.HttpTransportClientConfig{
		BaseURL:      viper.GetString(configconsts.API_ENDPOINT),
		AuthProvider: authProvider,
		Version: rorversion.RorVersion{
			Version: "v1",
			Commit:  "kind",
		},
		Role: viper.GetString(configconsts.ROLE),
	}
	transport := resttransport.NewRorHttpTransport(&clientConfig)
	RorClient = rorclient.NewRorClient(transport)
}
