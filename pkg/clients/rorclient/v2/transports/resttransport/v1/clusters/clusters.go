package clusters

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
		basePath: "/v1/clusters",
	}
}

func (c *V1Client) GetSelf(ctx context.Context) (apicontracts.ClusterSelf, error) {
	var selfdata apicontracts.ClusterSelf
	err := c.Client.GetJSON(ctx, c.basePath+"/self", &selfdata)
	if err != nil {
		return selfdata, err
	}

	return selfdata, nil
}

func (c *V1Client) GetById(ctx context.Context, id string) (*apicontracts.Cluster, error) {
	var cluster apicontracts.Cluster
	err := c.Client.GetJSON(ctx, c.basePath+"/"+id, &cluster)
	if err != nil {
		return nil, err
	}

	return &cluster, nil
}

func (c *V1Client) UpdateById(ctx context.Context, id string, cluster *apicontracts.Cluster) error {
	var dummy int
	return c.Client.PutJSON(ctx, c.basePath+"/"+id, cluster, &dummy)
}

func (c *V1Client) GetByFilter(ctx context.Context, filter apicontracts.Filter) (*[]apicontracts.Cluster, error) {
	var clusters apicontracts.PaginatedResult[apicontracts.Cluster]
	err := c.Client.PostJSON(ctx, c.basePath+"/filter", filter, &clusters)
	if err != nil {
		return nil, err
	}

	return &clusters.Data, nil
}

func (c *V1Client) Get(ctx context.Context, limit int, offset int) (*[]apicontracts.Cluster, error) {
	filter := apicontracts.Filter{
		Skip:  offset,
		Limit: limit,
		Sort: []apicontracts.SortMetadata{
			{
				SortField: "clustername",
				SortOrder: 1,
			},
		},
	}
	return c.GetByFilter(ctx, filter)
}

func (c *V1Client) GetAll(ctx context.Context) (*[]apicontracts.Cluster, error) {
	paginationLimit := 100
	nextBatch := 0
	var clusters []apicontracts.Cluster

	for {
		batch, err := c.Get(ctx, paginationLimit, nextBatch)
		if err != nil {
			return nil, err
		}
		if batch == nil || len(*batch) == 0 {
			return &clusters, nil
		}
		clusters = append(clusters, *batch...)
		nextBatch = nextBatch + paginationLimit
	}
}

func (c *V1Client) GetKubeconfig(ctx context.Context, clusterid, username, password string) (*apicontracts.ClusterKubeconfig, error) {
	var kubeconfig apicontracts.ClusterKubeconfig

	if len(clusterid) == 0 {
		return nil, fmt.Errorf("clusterid is required")
	}
	if len(username) == 0 {
		return nil, fmt.Errorf("username is required")
	}

	query := apicontracts.KubeconfigCredentials{
		Username: username,
		Password: password,
	}

	err := c.Client.PostJSON(ctx, c.basePath+"/"+clusterid+"/login", query, &kubeconfig)
	if err != nil {
		return nil, err
	}

	return &kubeconfig, nil
}

func (c *V1Client) Create(ctx context.Context, clusterInput apicontracts.Cluster) (string, error) {
	var clusterId string
	err := c.Client.PostJSON(ctx, c.basePath, clusterInput, &clusterId)
	if err != nil {
		return "", err
	}

	return clusterId, nil
}

func (c *V1Client) Register(ctx context.Context, data apicontracts.AgentApiKeyModel) (string, error) {
	var clusterResponse string
	err := c.Client.PostJSON(ctx, c.basePath+"/register", data, &clusterResponse)
	if err != nil {
		return "", err
	}

	return clusterResponse, nil
}

func (c *V1Client) SendHeartbeat(ctx context.Context, clusterReport apicontracts.Cluster) error {
	err := c.Client.PostJSON(ctx, "/v1/cluster/heartbeat", clusterReport, nil)
	return err
}
