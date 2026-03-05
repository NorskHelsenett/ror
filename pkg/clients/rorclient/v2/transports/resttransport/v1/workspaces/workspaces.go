package workspaces

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/transports/resttransport/httpclient"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type V1Client struct {
	Client   *httpclient.HttpTransportClient
	basePath string
}

func NewV1Client(client *httpclient.HttpTransportClient) *V1Client {
	return &V1Client{
		Client:   client,
		basePath: "/v1/workspaces",
	}
}

func (c *V1Client) GetByName(ctx context.Context, workspaceName string) (*apicontracts.Workspace, error) {
	var workspace apicontracts.Workspace
	err := c.Client.GetJSON(ctx, c.basePath+"/"+workspaceName, &workspace)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (c *V1Client) GetById(ctx context.Context, workspaceId string) (*apicontracts.Workspace, error) {
	var workspace apicontracts.Workspace
	err := c.Client.GetJSON(ctx, c.basePath+"/id/"+workspaceId, &workspace)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (c *V1Client) Get(ctx context.Context) (*[]apicontracts.Workspace, error) {
	var workspaces []apicontracts.Workspace

	err := c.Client.GetJSON(ctx, c.basePath, &workspaces)
	if err != nil {
		return nil, err
	}

	return &workspaces, nil
}

func (c *V1Client) GetAll(ctx context.Context) (*[]apicontracts.Workspace, error) {
	return c.Get(ctx)
}

func (c *V1Client) GetKubeconfig(ctx context.Context, workspacename, username, password string) (*apicontracts.ClusterKubeconfig, error) {
	var kubeconfig apicontracts.ClusterKubeconfig

	if len(workspacename) == 0 {
		return nil, fmt.Errorf("clusterid is required")
	}
	if len(username) == 0 {
		return nil, fmt.Errorf("username is required")
	}

	query := apicontracts.KubeconfigCredentials{
		Username: username,
		Password: password,
	}

	err := c.Client.PostJSON(ctx, c.basePath+"/"+workspacename+"/login", query, &kubeconfig)
	if err != nil {
		return nil, err
	}

	return &kubeconfig, nil
}
