package rortypes

import (
	"github.com/vitistack/common/pkg/v1alpha1"
)

type ResourceMachine struct {
	Spec   ResourceMachineSpec   `json:"spec"`
	Status ResourceMachineStatus `json:"status"`
}

type ResourceMachineSpec struct {
	ProviderSpec *v1alpha1.MachineSpec `json:"providerSpec"`
}

type ResourceMachineStatus struct {
	ProviderStatus *v1alpha1.MachineStatus `json:"providerStatus"`
}

// (r ResourceMachine) Get returns a pointer to the resource of type ResourceMachine
func (r *ResourceMachine) Get() *ResourceMachine {
	return r
}

// Machineinterface represents the interface for resources of the type machine
type Machineinterface interface {
	Get() *ResourceMachine
}
