package rortypes

import (
	"github.com/vitistack/common/pkg/v1alpha1"
)

type ResourceMachine struct {
	Spec   ResourceMachineSpec   `json:"spec"`
	Status ResourceMachineStatus `json:"status"`
}

type ResourceMachineSpec struct {
	ProviderSpec v1alpha1.MachineSpec `json:"providerSpec"`
}

type ResourceMachineStatus struct {
	ProviderStatus v1alpha1.MachineStatus `json:"providerStatus"`
}
