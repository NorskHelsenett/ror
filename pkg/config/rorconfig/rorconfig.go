package rorconfig

import (
	"os"

	kubernetesclient "github.com/NorskHelsenett/ror/pkg/clients/kubernetes"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient"
)

type RorConfig struct {
	clients  rorClients
	autoload bool
	config   ConfigMap
}

type rorClients struct {
	rorclient *rorclient.RorClient
	k8sclient *kubernetesclient.K8sClientsets
}

type ConfigMap map[ConfigConst]ConfigData

type ConfigData string

func (rc *RorConfig) LoadEnv(key ConfigConst) {
	rc.config[key] = ConfigData(os.Getenv(string(key)))
}

func (rc *RorConfig) Default(key ConfigConst, defaultValue string) {
	if rc.autoload {
		rc.LoadEnv(key)
	}
	if _, exists := rc.config[key]; !exists || rc.config[key] == "" {
		rc.config[key] = ConfigData(defaultValue)
	}
}

func (rc *RorConfig) AutoLoadEnv() {
	rc.autoload = true
	for key := range rc.config {
		rc.LoadEnv(key)
	}
}

func (rc *RorConfig) GetString(key ConfigConst) string {
	return rc.config[key].String()
}

// GetK8sClient returns the Kubernetes client from the RorConfig.
func (rc *RorConfig) GetK8sClient() *kubernetesclient.K8sClientsets {
	return rc.clients.k8sclient
}

// GetRorClient returns the ROR client from the RorConfig.
func (rc *RorConfig) GetRorClient() *rorclient.RorClient {
	return rc.clients.rorclient
}

func (cd ConfigData) String() string {
	return string(cd)
}

func (cd ConfigData) Bool() bool {
	return cd == "true"
}
