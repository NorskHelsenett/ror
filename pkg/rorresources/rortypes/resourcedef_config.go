package rortypes

import (
	"context"
	"fmt"

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
	identity := rorcontext.MustGetIdentityFromRorContext(ctx)
	for i, data := range r.Spec.Data {
		if data.Filter != string(identity.Type) {
			continue
		}
		res, err := configservice.Template(data.Value, ctx)
		if err != nil {
			return fmt.Errorf("failed to apply config service: %w", err)
		}
		r.Spec.Data[i].Value = res
	}
	return nil

}

// Configinterface represents the interface for resources of the type Config
type Configinterface interface {
	Get() *ResourceConfig
	ApplyOutputFilter(cr *CommonResource, ctx context.Context) error
}
