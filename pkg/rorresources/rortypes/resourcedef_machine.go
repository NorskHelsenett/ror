package rortypes

import (
	"github.com/vitistack/common/pkg/v1alpha1"
)

type ResourceMachine struct {
	Spec   v1alpha1.MachineSpec   `json:"spec"`
	Status v1alpha1.MachineStatus `json:"status"`
}
