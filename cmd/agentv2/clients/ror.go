// The package implements clients for the ror-agent
package clients

import (
	"fmt"

	kubernetesclient "github.com/NorskHelsenett/ror/pkg/clients/kubernetes"
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"github.com/NorskHelsenett/ror/pkg/config/rorclientconfig"

	"github.com/go-resty/resty/v2"
	"github.com/spf13/viper"
)

var Kubernetes *kubernetesclient.K8sClientsets
var RorConfig *rorclientconfig.RorClientConfig
var client *resty.Client

// Deprecated: GetOrCreateRorClient is deprecated use rorconfig.GetRorClient() instead
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

func InitClients(clientConfig rorclientconfig.ClientConfig) {
	rorclientconfig.InitRorClientConfig(clientConfig)
	RorConfig = rorclientconfig.RorConfig
	Kubernetes = RorConfig.GetKubernetesClientSet()
	viper.Set(configconsts.CLUSTER_ID, RorConfig.GetClusterId())
}
