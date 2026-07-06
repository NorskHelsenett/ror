package rortypes

// In tags: "parent" is the ID of the parent organizational unit.
// If empty, this is a top-level organizational unit and should be of type "organization".

type ResourceOrganizationalUnit struct {
	ID     string                           `json:"id"`
	Spec   ResourceOrganizationalUnitSpec   `json:"spec"`
	Status ResourceOrganizationalUnitStatus `json:"status"`
}

type ResourceOrganizationalUnitSpec struct {
	Name string                 `json:"name"`
	Type OrganizationalUnitType `json:"type"`
}

type ResourceOrganizationalUnitStatus struct {
	Name string                 `json:"name"`
	Type OrganizationalUnitType `json:"type"`
}

type OrganizationalUnitType string

const (
	OrganizationalUnitTypeOrganization OrganizationalUnitType = "organization"
	OrganizationalUnitTypeProject      OrganizationalUnitType = "project"
	OrganizationalUnitTypeGroup        OrganizationalUnitType = "group"
)
