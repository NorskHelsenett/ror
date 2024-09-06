package resttransport

import (
	"net/http"

	httpclient "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"

	restv1clusters "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/clusters"
	restv1datacenter "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/datacenter"
	restv1info "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/info"
	restv1metrics "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/metrics"
	restv1projects "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/projects"
	restv1resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/resources"
	restv1stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/stream"
	restv1workspaces "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/workspaces"
	restv2resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v2/resources"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v2/restclientv2self"
	v1clusters "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/clusters"
	v1datacenter "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/datacenter"
	v1info "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/info"
	v1metrics "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/metrics"
	v1projects "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/projects"
	v1resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/resources"
	v1stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/stream"
	v1workspaces "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/workspaces"
	v2resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/resources"
	v2self "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/rorclientv2self"
)

type RorHttpTransport struct {
	Client             *httpclient.HttpTransportClient
	streamClientV1     v1stream.StreamInterface
	infoClientV1       v1info.InfoInterface
	datacenterClientV1 v1datacenter.DatacenterInterface
	clustersClientV1   v1clusters.ClustersInterface
	workspacesClientV1 v1workspaces.WorkspacesInterface
	projectsClientV1   v1projects.ProjectsInterface
	resourcesClientV1  v1resources.ResourceInterface
	metricsClientV1    v1metrics.MetricsInterface
	resourcesClientV2  v2resources.ResourcesInterface
	selfClientV2       v2self.SelfInterface
}

func NewRorHttpTransport(config *httpclient.HttpTransportClientConfig) *RorHttpTransport {
	client := &httpclient.HttpTransportClient{
		Client: &http.Client{},
		Config: config,
	}
	t := &RorHttpTransport{
		Client:             client,
		streamClientV1:     restv1stream.NewV1Client(client),
		infoClientV1:       restv1info.NewV1Client(client),
		datacenterClientV1: restv1datacenter.NewV1Client(client),
		clustersClientV1:   restv1clusters.NewV1Client(client),
		selfClientV2:       restclientv2self.NewV2Client(client),
		workspacesClientV1: restv1workspaces.NewV1Client(client),
		projectsClientV1:   restv1projects.NewV1Client(client),
		resourcesClientV1:  restv1resources.NewV1Client(client),
		metricsClientV1:    restv1metrics.NewV1Client(client),
		resourcesClientV2:  restv2resources.NewV2Client(client),
	}
	return t
}

func (t *RorHttpTransport) Stream() v1stream.StreamInterface {
	return t.streamClientV1
}

func (t *RorHttpTransport) Info() v1info.InfoInterface {
	return t.infoClientV1
}

func (t *RorHttpTransport) Datacenters() v1datacenter.DatacenterInterface {
	return t.datacenterClientV1
}

func (t *RorHttpTransport) Clusters() v1clusters.ClustersInterface {
	return t.clustersClientV1
}
func (t *RorHttpTransport) Workspaces() v1workspaces.WorkspacesInterface {
	return t.workspacesClientV1
}
func (t *RorHttpTransport) Projects() v1projects.ProjectsInterface {
	return t.projectsClientV1
}
func (t *RorHttpTransport) Metrics() v1metrics.MetricsInterface {
	return t.metricsClientV1
}

func (t *RorHttpTransport) Resources() v1resources.ResourceInterface {
	return t.resourcesClientV1
}

func (t *RorHttpTransport) ResourcesV2() v2resources.ResourcesInterface {
	return t.resourcesClientV2
}

func (t *RorHttpTransport) Self() v2self.SelfInterface {
	return t.selfClientV2
}

func (t *RorHttpTransport) Ping() error {
	_, err := t.infoClientV1.GetVersion()
	return err
}
