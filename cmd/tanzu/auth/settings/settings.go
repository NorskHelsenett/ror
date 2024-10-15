package settings

import (
	"encoding/json"
	"fmt"
	"io"
	"os"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
	"k8s.io/client-go/rest"
)

var (
	TanzuAuthVersionNumber = "1.1.0"
	TanzuAuthCommit        = "local-dev"
	Development            = false

	TanzuUsername = ""
	TanzuPwd      = ""

	KubeVspherePath = "kubectl"
	KubectlPath     = "kubectl"

	K8sConfig *rest.Config

	ErrorCount int64 = 0
)

func init() {
	viper.AutomaticEnv()

	viper.SetDefault(configconsts.DEVELOPMENT, false)
	viper.SetDefault(configconsts.HTTP_PORT, "18082")
	viper.SetDefault(configconsts.API_ENDPOINT, "https://api.ror.sky.test.nhn.no")
	viper.SetDefault(configconsts.ROR_OPERATOR_NAMESPACE, "nhn-ror")
	viper.SetDefault(configconsts.TANZU_AUTH_KUBE_VSPHERE_PATH, "kubectl-vsphere")
	viper.SetDefault(configconsts.TANZU_AUTH_KUBECTL_PATH, "kubectl")
	viper.SetDefault(configconsts.TANZU_AUTH_CONFIG_FOLDER_PATH, "/configs")
}

func Load() {
	environment := viper.GetString(configconsts.ENVIRONMENT)
	rlog.Info("loaded environment", rlog.String("Environment", environment))
	Development = viper.GetBool(configconsts.DEVELOPMENT)

	kubeVspherePath := viper.GetString(configconsts.TANZU_AUTH_KUBE_VSPHERE_PATH)
	if len(kubeVspherePath) > 0 {
		KubeVspherePath = kubeVspherePath
	}
	kubectlPath := viper.GetString(configconsts.TANZU_AUTH_KUBECTL_PATH)
	if len(kubectlPath) > 0 {
		KubectlPath = kubectlPath
	}

	getVersion()
}

func getVersion() {
	var version Version
	if !Development {
		jsonFile, err := os.Open("/version.json")
		if err != nil {
			rlog.Error("error opening json file", err)
		}
		defer func(jsonFile *os.File) {
			_ = jsonFile.Close()
		}(jsonFile)

		byteValue, _ := io.ReadAll(jsonFile)

		_ = json.Unmarshal(byteValue, &version)
		TanzuAuthVersionNumber = fmt.Sprintf("%d.%d.%d", version.Major, version.Minor, version.Patch)
		TanzuAuthCommit = version.CommitSha
	}
	viper.Set(configconsts.VERSION, TanzuAuthVersionNumber)
	viper.Set(configconsts.COMMIT, TanzuAuthCommit)
}

type Version struct {
	Major     int    `json:"major"`
	Minor     int    `json:"minor"`
	Patch     int    `json:"patch"`
	CommitSha string `json:"commitSha"`
}
