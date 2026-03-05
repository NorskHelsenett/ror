package projects

import (
	"context"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type ProjectsInterface interface {
	GetById(ctx context.Context, id string) (*apicontracts.Project, error)
	Get(ctx context.Context, limit int, offset int) (*[]apicontracts.Project, error)
	GetAll(ctx context.Context) (*[]apicontracts.Project, error)
}
