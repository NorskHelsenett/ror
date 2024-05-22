package apiresponses

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type WorkspacesResponse struct {
	Workspaces []apicontracts.Workspace `json:"workspaces"`
}
