package projects

import (
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
		basePath: "/v1/projects",
	}
}

func (c *V1Client) GetById(id string) (*apicontracts.Project, error) {
	var project apicontracts.Project
	err := c.Client.GetJSON(c.basePath+"/"+id, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *V1Client) Get(limit int, offset int) (*[]apicontracts.Project, error) {
	var projects apicontracts.PaginatedResult[apicontracts.Project]

	filter := apicontracts.Filter{
		Skip:  offset,
		Limit: limit,
		Sort: []apicontracts.SortMetadata{
			{
				SortField: "projectname",
				SortOrder: 1,
			},
		},
	}

	err := c.Client.PostJSON(c.basePath+"/filter", filter, &projects)
	if err != nil {
		return nil, err
	}

	return &projects.Data, nil
}

func (c *V1Client) GetAll() (*[]apicontracts.Project, error) {
	paginationLimit := 100
	nextBatch := 0
	var projects []apicontracts.Project

	for {
		batch, err := c.Get(paginationLimit, nextBatch)
		if err != nil {
			return nil, err
		}
		if batch == nil || len(*batch) == 0 {
			return &projects, nil
		}
		projects = append(projects, *batch...)
		nextBatch = nextBatch + paginationLimit
	}
}
