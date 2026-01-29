package projects

import "github.com/NorskHelsenett/ror/pkg/apicontracts"

type ProjectsInterface interface {
	GetById(id string) (*apicontracts.Project, error)
	Get(limit int, offset int) (*[]apicontracts.Project, error)
	GetAll() (*[]apicontracts.Project, error)
}
