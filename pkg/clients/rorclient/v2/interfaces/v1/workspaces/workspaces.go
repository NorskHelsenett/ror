package workspaces

import "github.com/NorskHelsenett/ror/pkg/apicontracts"

type WorkspacesInterface interface {
	GetByName(workspaceName string) (*apicontracts.Workspace, error)
	GetById(workspaceId string) (*apicontracts.Workspace, error)
	Get() (*[]apicontracts.Workspace, error)
	GetAll() (*[]apicontracts.Workspace, error)
	GetKubeconfig(workspacename, username, password string) (*apicontracts.ClusterKubeconfig, error)
}
