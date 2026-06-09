package rortypes

// K8s unknown resource struct

type ResourceUnknown struct {
}

// (r ResourceUnknown) Get returns a pointer to the resource of type ResourceUnknown
func (r *ResourceUnknown) Get() *ResourceUnknown {
	return r
}

// Unknowninterface represents the interface for resources of the type unknown
type Unknowninterface interface {
	Get() *ResourceUnknown
}
