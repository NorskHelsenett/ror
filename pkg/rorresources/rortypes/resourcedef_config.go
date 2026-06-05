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
	Key    string `json:"key"`
	Value  string `json:"value"`
	Source string `json:"source"`
	Filter string `json:"filter"`
}

// (r ResourceConfig) Get returns a pointer to the resource of type ResourceConfig
func (r *ResourceConfig) Get() *ResourceConfig {
	return r
}

func (r *ResourceConfig) ApplyOutputFilter(ctx context.Context, cr *CommonResource) error {
	identity := rorcontext.MustGetIdentityFromRorContext(ctx)
	for i, data := range r.Spec.Data {

		//make string to lower case to avoid case sensitivity issues
		// gogo agent
		if data.Filter != string(identity.Type) {
			fmt.Println("this is not in cluster")
			continue
		}
		fmt.Println("this is cluster")
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
	ApplyOutputFilter(ctx context.Context, cr *CommonResource) error
}
