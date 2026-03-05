package projects

import (
	"context"

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
		basePath: "/v1/projects",
	}
}

func (c *V1Client) GetById(ctx context.Context, id string) (*apicontracts.Project, error) {
	var project apicontracts.Project
	err := c.Client.GetJSON(ctx, c.basePath+"/"+id, &project)
	if err != nil {
		return nil, err
	}

	return &project, nil
}

func (c *V1Client) Get(ctx context.Context, limit int, offset int) (*[]apicontracts.Project, error) {
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

	err := c.Client.PostJSON(ctx, c.basePath+"/filter", filter, &projects)
	if err != nil {
		return nil, err
	}

	return &projects.Data, nil
}

func (c *V1Client) GetAll(ctx context.Context) (*[]apicontracts.Project, error) {
	paginationLimit := 100
	nextBatch := 0
	var projects []apicontracts.Project

	for {
		batch, err := c.Get(ctx, paginationLimit, nextBatch)
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
