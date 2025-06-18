package apiresourcecontracts

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceProject struct {
	ApiVersion string              `json:"apiVersion"`
	Kind       string              `json:"kind"`
	Metadata   ResourceMetadata    `json:"metadata"`
	Spec       ResourceProjectSpec `json:"spec"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceProjectSpec struct {
	ProjectName string                    `json:"projectName"`
	Description string                    `json:"description"`
	Active      bool                      `json:"active"`
	CreatedTime string                    `json:"createdTime"`
	UpdatedTime string                    `json:"updatedTime"`
	Roles       []ResourceProjectSpecRole `json:"roles"`
	Workorder   string                    `json:"workorder"`
	ServiceTag  string                    `json:"serviceTag"`
	Tags        []string                  `json:"tags"`
}

// Deprecated: This type is only to be used in resource/v1 and will be deprecated
type ResourceProjectSpecRole struct {
	Upn   string `json:"upn"`
	Name  string `json:"name"`
	Role  string `json:"role"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
