package rortypes

import (
	"bytes"
	"context"
	"fmt"
	"text/template"

	"github.com/NorskHelsenett/ror/pkg/context/rorcontext"
	"github.com/NorskHelsenett/ror/pkg/services/configservice"
)

// config struct
type ResourceConfig struct {
	Spec ResourceConfigSpec `json:"spec"`
}

type ResourceConfigSpec struct {
	Data []ResourceConfigData `json:"resourceConfigData"`
}

type ResourceConfigData struct {
	Value  string `json:"value"`
	Source string `json:"source"`
	Filter string `json:"filter"`
}

// (r ResourceConfig) Get returns a pointer to the resource of type ResourceConfig
func (r *ResourceConfig) Get() *ResourceConfig {
	return r
}

func (r *ResourceConfig) ApplyOutputFilter(cr *CommonResource, ctx context.Context) error {

	res, err := configservice.ConfigService(r, ctx)
	if err != nil {
		return fmt.Errorf("failed to apply config service: %w", err)
	}



}

// Configinterface represents the interface for resources of the type Config
type Configinterface interface {
	Get() *ResourceConfig
	ApplyOutputFilter(cr *CommonResource, ctx context.Context) error
}
