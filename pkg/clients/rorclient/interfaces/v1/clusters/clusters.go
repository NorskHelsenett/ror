package clusters

import (
	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type ClustersInterface interface {
	GetSelf() (apicontracts.ClusterSelf, error)
	GetById(id string) (*apicontracts.Cluster, error)
	UpdateById(id string, cluster *apicontracts.Cluster) error
	GetByFilter(filter apicontracts.Filter) (*[]apicontracts.Cluster, error)
	Get(limit int, offset int) (*[]apicontracts.Cluster, error)
	GetAll() (*[]apicontracts.Cluster, error)
	GetKubeconfig(clusterid, username, password string) (*apicontracts.ClusterKubeconfig, error)
	Create(cluster apicontracts.Cluster) (string, error)
	Register(data apicontracts.AgentApiKeyModel) (string, error)
}
