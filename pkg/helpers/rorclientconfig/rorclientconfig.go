package rorclientconfig

import (
	kubernetesclient "github.com/NorskHelsenett/ror/pkg/clients/kubernetes"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
)

type RorClientInterface interface {
	RorAPIClientInterface
	RorOwnerrefInterface
	RorKubernetesInterface
}

type RorAPIClientInterface interface {
	GetRorClient() (rorclient.RorClientInterface, error)
}
type RorOwnerrefInterface interface {
	GetOwnerref() rorresourceowner.RorResourceOwnerReference
	SetOwnerref(ownerref rorresourceowner.RorResourceOwnerReference)
}

type RorKubernetesInterface interface {
	GetKubernetesClientSet() *kubernetesclient.K8sClientsets
}
