package rortypes

type ResourceKubernetesMachineClass struct {
	Spec   ResourceKubernetesMachineClassSpec   `json:"spec"`
	Status ResourceKubernetesMachineClassStatus `json:"status"`
}

type ResourceKubernetesMachineClassSpec struct {
	Name   string `json:"name"`
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
	Gpu    bool   `json:"gpu"`
}

type ResourceKubernetesMachineClassStatus struct {
	Name   string `json:"name"`
	Cpu    string `json:"cpu"`
	Memory string `json:"memory"`
	Gpu    bool   `json:"gpu"`
}

// (r ResourceKubernetesMachineClass) Get returns a pointer to the resource of type ResourceKubernetesMachineClass
func (r *ResourceKubernetesMachineClass) Get() *ResourceKubernetesMachineClass {
	return r
}
