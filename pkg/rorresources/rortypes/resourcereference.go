package rortypes

type RorResourceReference struct {
	Name     string    `json:"name"`
	Uid      string    `json:"uid"`
	Resource *Resource `json:"resource,omitempty"`
}

func (r RorResourceReference) String() string {
	return r.Name
}

func (r RorResourceReference) GetName() string {
	return r.Name
}

func (r RorResourceReference) GetUid() string {
	return r.Uid
}

type Resource struct {
	Uid  string `json:"uid"`
	Name string `json:"name"`
	Kind string `json:"kind"`
}
