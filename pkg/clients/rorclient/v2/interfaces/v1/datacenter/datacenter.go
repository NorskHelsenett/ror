package datacenter

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type DatacenterInterface interface {
	Get(ctx context.Context) (*[]apicontracts.Datacenter, error)
	GetById(ctx context.Context, id string) (*apicontracts.Datacenter, error)
	GetByName(ctx context.Context, name string) (*apicontracts.Datacenter, error)
	Post(ctx context.Context, data apicontracts.DatacenterModel) (*apicontracts.Datacenter, error)
	Put(ctx context.Context, id string, data apicontracts.DatacenterModel) (*apicontracts.Datacenter, error)
}
