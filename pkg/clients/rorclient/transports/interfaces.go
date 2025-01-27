package transports

import (
	v1clusters "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/clusters"
	v1datacenter "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/datacenter"
	v1info "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/info"
	v1metrics "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/metrics"
	v1projects "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/projects"
	v1resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/resources"
	v1stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/stream"
	v1workspaces "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/workspaces"
	v2resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/resources"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/rorclientv2self"
	v2stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/v2stream"
)

type RorTransport interface {
	Stream() v1stream.StreamInterface
	Info() v1info.InfoInterface
	Datacenters() v1datacenter.DatacenterInterface
	Clusters() v1clusters.ClustersInterface
	Self() rorclientv2self.SelfInterface
	Workspaces() v1workspaces.WorkspacesInterface
	Projects() v1projects.ProjectsInterface
	Resources() v1resources.ResourceInterface
	Metrics() v1metrics.MetricsInterface
	ResourcesV2() v2resources.ResourcesInterface
	Streamv2() v2stream.StreamInterface
	Ping() error
}
