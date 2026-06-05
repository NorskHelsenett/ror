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
	Filter string                        `json:"filter"`
	Data   map[string]ResourceConfigData `json:"data"` //map of key value pairs, where value can be a template string to be resolved by configservice
}

type ResourceConfigData string

// (r ResourceConfig) Get returns a pointer to the resource of type ResourceConfig
func (r *ResourceConfig) Get() *ResourceConfig {
	return r
}

//value: {{arguser: vault(/bla/bla/bla), argopass: vault(/bla/bla/bla)}}

func (r *ResourceConfig) ApplyOutputFilter(ctx context.Context, cr *CommonResource) error {
	identity := rorcontext.MustGetIdentityFromRorContext(ctx)
	for i, data := range r.Spec.Data {

		//make string to lower case to avoid case sensitivity issues
		// gogo agent
		if r.Spec.Filter != string(identity.Type) {
			fmt.Println("this is not in cluster")
			continue
		}
		fmt.Println("this is cluster")
		res, err := configservice.Template(string(data), ctx)
		if err != nil {
			return fmt.Errorf("failed to apply config service: %w", err)
		}
		r.Spec.Data[i] = ResourceConfigData(res)
	}
	return nil

}

// Configinterface represents the interface for resources of the type Config
type Configinterface interface {
	Get() *ResourceConfig
	ApplyOutputFilter(ctx context.Context, cr *CommonResource) error
}
