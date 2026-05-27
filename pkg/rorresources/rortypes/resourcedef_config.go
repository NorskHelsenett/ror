package rortypes

// config struct
type ResourceConfig struct {
}

// (r ResourceConfig) Get returns a pointer to the resource of type ResourceConfig
func (r *ResourceConfig) Get() *ResourceConfig {
	return r
}

// Configinterface represents the interface for resources of the type Config
type Configinterface interface {
	Get() *ResourceConfig
}
