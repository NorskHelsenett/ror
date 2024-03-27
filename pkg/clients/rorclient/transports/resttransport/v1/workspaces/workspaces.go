package workspaces

import (
	"fmt"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"

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

func (c *V1Client) GetByName(workspaceName string) (*apicontracts.Workspace, error) {
	var workspace apicontracts.Workspace
	err := c.Client.GetJSON(c.basePath+"/"+workspaceName, &workspace)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (c *V1Client) GetById(workspaceId string) (*apicontracts.Workspace, error) {
	var workspace apicontracts.Workspace
	err := c.Client.GetJSON(c.basePath+"/id/"+workspaceId, &workspace)
	if err != nil {
		return nil, err
	}

	return &workspace, nil
}

func (c *V1Client) Get() (*[]apicontracts.Workspace, error) {
	var workspaces []apicontracts.Workspace

	err := c.Client.GetJSON(c.basePath, &workspaces)
	if err != nil {
		return nil, err
	}

	return &workspaces, nil
}

func (c *V1Client) GetAll() (*[]apicontracts.Workspace, error) {
	return c.Get()
}

func (c *V1Client) GetKubeconfig(workspacename, username, password string) (*apicontracts.ClusterKubeconfig, error) {
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

	err := c.Client.PostJSON(c.basePath+"/"+workspacename+"/login", query, &kubeconfig)
	if err != nil {
		return nil, err
	}

	return &kubeconfig, nil
}
