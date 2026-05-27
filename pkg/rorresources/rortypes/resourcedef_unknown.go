package rortypes

// K8s unknown resource struct

type ResourceUnknown struct {
}

// (r ResourceUnknown) Get returns a pointer to the resource of type ResourceUnknown
func (r *ResourceUnknown) Get() *ResourceUnknown {
	return r
}
