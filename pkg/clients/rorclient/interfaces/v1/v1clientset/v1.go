package v1clientset

import (
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/clientinterface"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/transports/transportinterface"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/acl"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/clusters"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/datacenter"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/info"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/metrics"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/projects"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/resources"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/stream"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/token"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/workspaces"
)

type ClientSet struct {
	transport        transportinterface.RorTransport
	streamClient     stream.StreamInterface
	infoClient       info.InfoInterface
	datacenterClient datacenter.DatacenterInterface
	clustersClient   clusters.ClustersInterface
	workspacesClient workspaces.WorkspacesInterface
	projectsClient   projects.ProjectsInterface
	resourceClient   resources.ResourceInterface
	metricsClient    metrics.MetricsInterface
	aclClient        acl.AclInterface
	tokenClient      token.TokenInterface
}

func NewV1ClientSet(transport transportinterface.RorTransport) clientinterface.RorCommonClientApiInterfaceV1 {
	return &ClientSet{
		aclClient:        transport.Acl(),
		streamClient:     transport.Stream(),
		infoClient:       transport.Info(),
		datacenterClient: transport.Datacenters(),
		clustersClient:   transport.Clusters(),
		workspacesClient: transport.Workspaces(),
		projectsClient:   transport.Projects(),
		resourceClient:   transport.Resources(),
		metricsClient:    transport.Metrics(),
		tokenClient:      transport.Token(),
	}
}

func (c *ClientSet) Acl() acl.AclInterface {
	return c.aclClient
}

func (c *ClientSet) Clusters() clusters.ClustersInterface {
	return c.clustersClient
}

func (c *ClientSet) Datacenters() datacenter.DatacenterInterface {
	return c.datacenterClient
}

func (c *ClientSet) Info() info.InfoInterface {
	return c.infoClient
}

func (c *ClientSet) Metrics() metrics.MetricsInterface {
	return c.metricsClient
}

func (c *ClientSet) Projects() projects.ProjectsInterface {
	return c.projectsClient
}

func (c *ClientSet) Resources() resources.ResourceInterface {
	return c.resourceClient
}

func (c *ClientSet) Stream() stream.StreamInterface {
	return c.streamClient
}

func (c *ClientSet) Token() token.TokenInterface {
	return c.tokenClient
}
func (c *ClientSet) Workspaces() workspaces.WorkspacesInterface {
	return c.workspacesClient
}
