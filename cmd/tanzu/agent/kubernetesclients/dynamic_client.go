package kubernetesclients

import (
	"github.com/NorskHelsenett/ror/cmd/tanzu/agent/auth"

	"github.com/NorskHelsenett/ror/pkg/config/configconsts"

	"github.com/NorskHelsenett/ror/pkg/rlog"

	"github.com/spf13/viper"
	"k8s.io/client-go/dynamic"
)

func GetDynamicClientOrDie() dynamic.Interface {
	tanzuAccess := viper.GetBool(configconsts.TANZU_AGENT_TANZU_ACCESS)
	k8sConfig, err := auth.GetK8sConfig(tanzuAccess)
	if err != nil {
		rlog.Error("failed to get k8s config", err)
		panic(err)
	}

	dynamicClient, err := dynamic.NewForConfig(k8sConfig)
	if err != nil {
		rlog.Error("failed to get dynamic client", err)
		panic(err)
	}

	return dynamicClient
}
