package mocktransportprojects

import (
	"errors"

	"github.com/NorskHelsenett/ror/pkg/apicontracts"
)

type V1Client struct{}

func NewV1Client() *V1Client {
	return &V1Client{}
}

func (c *V1Client) GetById(id string) (*apicontracts.Project, error) {
	if id == "" {
		return nil, errors.New("project ID cannot be empty")
	}

	project := &apicontracts.Project{
		ID:   id,
		Name: "Mock Project " + id,
	}
	return project, nil
}

func (c *V1Client) Get(limit int, offset int) (*[]apicontracts.Project, error) {
	projects := []apicontracts.Project{
		{
			ID:   "mock-project-001",
			Name: "Mock Project 1",
		},
	}
	return &projects, nil
}

func (c *V1Client) GetAll() (*[]apicontracts.Project, error) {
	projects := []apicontracts.Project{
		{
			ID:   "mock-project-001",
			Name: "Mock Project 1",
		},
		{
			ID:   "mock-project-002",
			Name: "Mock Project 2",
		},
	}
	return &projects, nil
}
