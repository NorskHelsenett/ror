package rortypes

type ResourceManagedDatabase struct {
	ID       string                         `json:"id"`
	Provider string                         `json:"provider"`
	Spec     *ResourceManagedDatabaseSpec   `json:"spec"`
	Status   *ResourceManagedDatabaseStatus `json:"status"`
}

type ResourceManagedDatabaseSpec struct {
	Name   string `json:"name"`
	Engine string `json:"engine"`
}

type ResourceManagedDatabaseStatus struct {
	Name   string `json:"name"`
	Engine string `json:"engine"`
}

// (r ResourceManagedDatabase) Get returns a pointer to the resource of type ResourceManagedDatabase
func (r *ResourceManagedDatabase) Get() *ResourceManagedDatabase {
	return r
}

// ManagedDatabaseinterface represents the interface for resources of the type manageddatabase
type ManagedDatabaseinterface interface {
	Get() *ResourceManagedDatabase
}

// (r *ResourceManagedDatabase) ApplyInputFilter Applies the input filter to the resource
func (r *ResourceManagedDatabase) ApplyInputFilter(cr *CommonResource) error {
	return nil
}
