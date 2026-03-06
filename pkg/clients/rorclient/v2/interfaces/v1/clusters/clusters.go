package clusters

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type ClustersInterface interface {
	GetSelf(ctx context.Context) (apicontracts.ClusterSelf, error)
	GetById(ctx context.Context, id string) (*apicontracts.Cluster, error)
	UpdateById(ctx context.Context, id string, cluster *apicontracts.Cluster) error
	GetByFilter(ctx context.Context, filter apicontracts.Filter) (*[]apicontracts.Cluster, error)
	Get(ctx context.Context, limit int, offset int) (*[]apicontracts.Cluster, error)
	GetAll(ctx context.Context) (*[]apicontracts.Cluster, error)
	GetKubeconfig(ctx context.Context, clusterid, username, password string) (*apicontracts.ClusterKubeconfig, error)
	Create(ctx context.Context, cluster apicontracts.Cluster) (string, error)
	Register(ctx context.Context, data apicontracts.AgentApiKeyModel) (string, error)
	SendHeartbeat(ctx context.Context, clusterReport apicontracts.Cluster) error
}
