package rortypes

type ResourceProvider struct {
}

// (r ResourceProvider) Get returns a pointer to the resource of type ResourceProvider
func (r *ResourceProvider) Get() *ResourceProvider {
	return r
}

// Providerinterface represents the interface for resources of the type provider
type Providerinterface interface {
	Get() *ResourceProvider
}
