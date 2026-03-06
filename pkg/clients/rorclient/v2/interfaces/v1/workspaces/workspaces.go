package workspaces

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type WorkspacesInterface interface {
	GetByName(ctx context.Context, workspaceName string) (*apicontracts.Workspace, error)
	GetById(ctx context.Context, workspaceId string) (*apicontracts.Workspace, error)
	Get(ctx context.Context) (*[]apicontracts.Workspace, error)
	GetAll(ctx context.Context) (*[]apicontracts.Workspace, error)
	GetKubeconfig(ctx context.Context, workspacename, username, password string) (*apicontracts.ClusterKubeconfig, error)
}
