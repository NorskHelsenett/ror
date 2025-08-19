package mocktransportworkspaces

import (
	"errors"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type V1Client struct{}

func NewV1Client() *V1Client {
	return &V1Client{}
}

func (c *V1Client) GetByName(workspaceName string) (*apicontracts.Workspace, error) {
	if workspaceName == "" {
		return nil, errors.New("workspace name cannot be empty")
	}

	workspace := &apicontracts.Workspace{
		Name: workspaceName,
		ID:   "mock-workspace-by-name-001",
	}
	return workspace, nil
}

func (c *V1Client) GetById(workspaceId string) (*apicontracts.Workspace, error) {
	if workspaceId == "" {
		return nil, errors.New("workspace ID cannot be empty")
	}

	workspace := &apicontracts.Workspace{
		Name: "Mock Workspace " + workspaceId,
		ID:   workspaceId,
	}
	return workspace, nil
}

func (c *V1Client) Get() (*[]apicontracts.Workspace, error) {
	workspaces := []apicontracts.Workspace{
		{
			Name: "Mock Workspace 1",
			ID:   "mock-workspace-001",
		},
	}
	return &workspaces, nil
}

func (c *V1Client) GetAll() (*[]apicontracts.Workspace, error) {
	workspaces := []apicontracts.Workspace{
		{
			Name: "Mock Workspace 1",
			ID:   "mock-workspace-001",
		},
		{
			Name: "Mock Workspace 2",
			ID:   "mock-workspace-002",
		},
	}
	return &workspaces, nil
}

func (c *V1Client) GetKubeconfig(workspacename, username, password string) (*apicontracts.ClusterKubeconfig, error) {
	if workspacename == "" {
		return nil, errors.New("workspace name cannot be empty")
	}

	kubeconfig := &apicontracts.ClusterKubeconfig{
		Data:     "mock-workspace-kubeconfig-data",
		Status:   "success",
		DataType: "yaml",
	}
	return kubeconfig, nil
}
