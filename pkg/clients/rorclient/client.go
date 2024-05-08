package rorclient

import (
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports"
	v1clusters "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/clusters"
	v1datacenter "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/datacenter"
	v1info "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/info"
	v1metrics "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/metrics"
	v1projects "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/projects"
	v1resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/resources"
	v1workspaces "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/workspaces"
	v2resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/resources"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/rorclientv2self"
)

type RorConfig struct {
	Host string
}

type RorClient struct {
	Transport          transports.RorTransport
	infoClientV1       v1info.InfoInterface
	datacenterClientV1 v1datacenter.DatacenterInterface
	clustersClientV1   v1clusters.ClustersInterface
	workspacesClientV1 v1workspaces.WorkspacesInterface
	selfClientV2       rorclientv2self.SelfInterface
	projectsClientV1   v1projects.ProjectsInterface
	resourceClientV1   v1resources.ResourceInterface
	metricsClientV1    v1metrics.MetricsInterface
	resourcesClientV2  v2resources.ResourcesInterface
}

func NewRorClient(transport transports.RorTransport) *RorClient {
	return &RorClient{
		Transport:          transport,
		infoClientV1:       transport.Info(),
		datacenterClientV1: transport.Datacenters(),
		clustersClientV1:   transport.Clusters(),
		workspacesClientV1: transport.Workspaces(),
		projectsClientV1:   transport.Projects(),
		selfClientV2:       transport.Self(),
		resourceClientV1:   transport.Resources(),
		metricsClientV1:    transport.Metrics(),
		resourcesClientV2:  transport.ResourcesV2(),
	}
}

func (c *RorClient) SetTransport(transport transports.RorTransport) {
	c.Transport = transport
}
func (c *RorClient) Info() v1info.InfoInterface {
	return c.infoClientV1
}
func (c *RorClient) Self() rorclientv2self.SelfInterface {
	return c.selfClientV2
}

func (c *RorClient) Datacenters() v1datacenter.DatacenterInterface {
	return c.datacenterClientV1
}
func (c *RorClient) Clusters() v1clusters.ClustersInterface {
	return c.clustersClientV1
}
func (c *RorClient) Workspaces() v1workspaces.WorkspacesInterface {
	return c.workspacesClientV1
}

func (c *RorClient) Projects() v1projects.ProjectsInterface {
	return c.projectsClientV1
}
func (c *RorClient) Metrics() v1metrics.MetricsInterface {
	return c.metricsClientV1
}

func (c *RorClient) Resources() v1resources.ResourceInterface {
	return c.resourceClientV1
}

func (c *RorClient) Ping() error {
	return c.Transport.Ping()
}

func (c *RorClient) ResourceV2() v2resources.ResourcesInterface {
	return c.resourcesClientV2
}
