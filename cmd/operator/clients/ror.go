// The package implements clients for the ror-agent
package clients

import (
	kubernetesclient "github.com/NorskHelsenett/ror/pkg/clients/kubernetes"
	"github.com/NorskHelsenett/ror/pkg/config/configconsts"
	"github.com/NorskHelsenett/ror/pkg/config/rorclientconfig"

	"github.com/spf13/viper"
)

var Kubernetes *kubernetesclient.K8sClientsets
var RorConfig *rorclientconfig.RorClientConfig

func InitClients(clientConfig rorclientconfig.ClientConfig) {
	rorclientconfig.InitRorClientConfig(clientConfig)
	RorConfig = rorclientconfig.RorConfig
	Kubernetes = RorConfig.GetKubernetesClientSet()
	viper.Set(configconsts.CLUSTER_ID, RorConfig.GetClusterId())
	viper.Set(configconsts.API_KEY, RorConfig.GetApiKey())
}
