package mocktransportclusters

import (
	"errors"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type V1Client struct{}

func NewV1Client() *V1Client {
	return &V1Client{}
}

func (c *V1Client) GetSelf() (apicontracts.ClusterSelf, error) {
	return apicontracts.ClusterSelf{
		ClusterId: "mock-cluster-001",
	}, nil
}

func (c *V1Client) GetById(id string) (*apicontracts.Cluster, error) {
	if id == "" {
		return nil, errors.New("cluster ID cannot be empty")
	}

	cluster := &apicontracts.Cluster{
		ClusterId:   id,
		ClusterName: "mock-cluster-" + id,
		Environment: "development",
	}
	return cluster, nil
}

func (c *V1Client) UpdateById(id string, cluster *apicontracts.Cluster) error {
	if id == "" {
		return errors.New("cluster ID cannot be empty")
	}
	if cluster == nil {
		return errors.New("cluster cannot be nil")
	}
	return nil
}

func (c *V1Client) GetByFilter(filter apicontracts.Filter) (*[]apicontracts.Cluster, error) {
	clusters := []apicontracts.Cluster{
		{
			ClusterId:   "mock-cluster-001",
			ClusterName: "mock-cluster-1",
			Environment: "development",
		},
		{
			ClusterId:   "mock-cluster-002",
			ClusterName: "mock-cluster-2",
			Environment: "staging",
		},
	}
	return &clusters, nil
}

func (c *V1Client) Get(limit int, offset int) (*[]apicontracts.Cluster, error) {
	clusters := []apicontracts.Cluster{
		{
			ClusterId:   "mock-cluster-001",
			ClusterName: "mock-cluster-1",
			Environment: "development",
		},
	}
	return &clusters, nil
}

func (c *V1Client) GetAll() (*[]apicontracts.Cluster, error) {
	clusters := []apicontracts.Cluster{
		{
			ClusterId:   "mock-cluster-001",
			ClusterName: "mock-cluster-1",
			Environment: "development",
		},
		{
			ClusterId:   "mock-cluster-002",
			ClusterName: "mock-cluster-2",
			Environment: "staging",
		},
	}
	return &clusters, nil
}

func (c *V1Client) GetKubeconfig(clusterid, username, password string) (*apicontracts.ClusterKubeconfig, error) {
	if clusterid == "" {
		return nil, errors.New("cluster ID cannot be empty")
	}

	kubeconfig := &apicontracts.ClusterKubeconfig{
		Data:     "mock-kubeconfig-data",
		Status:   "success",
		DataType: "yaml",
	}
	return kubeconfig, nil
}

func (c *V1Client) Create(cluster apicontracts.Cluster) (string, error) {
	if cluster.ClusterName == "" {
		return "", errors.New("cluster name cannot be empty")
	}
	return "mock-cluster-new-001", nil
}

func (c *V1Client) Register(data apicontracts.AgentApiKeyModel) (string, error) {
	if data.Identifier == "" {
		return "", errors.New("identifier cannot be empty")
	}
	return "mock-api-key-12345", nil
}
