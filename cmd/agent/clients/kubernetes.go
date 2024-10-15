package clients

import (
	kubernetesclient "github.com/NorskHelsenett/ror/pkg/clients/kubernetes"
)

var Kubernetes *kubernetesclient.K8sClientsets

func Initialize() {
	Kubernetes = kubernetesclient.NewK8sClientConfig()
}
