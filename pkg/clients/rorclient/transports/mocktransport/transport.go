package mocktransport

import (
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportacl"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportclusters"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportdatacenter"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportinfo"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportmetrics"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportprojects"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportresources"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportresourcesv2"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportself"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportstatus"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportstream"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportstreamv2"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/mocktransport/mocktransportworkspaces"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/transportstatus"
	v1Acl "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/acl"
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
	v2stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/v2stream"
)

// Compile-time check to ensure RorMockTransport implements RorTransport
var _ transports.RorTransport = (*RorMockTransport)(nil)

type RorMockTransport struct {
	status             transportstatus.RorTransportStatus
	streamClientV1     v1stream.StreamInterface
	infoClientV1       v1info.InfoInterface
	datacenterClientV1 v1datacenter.DatacenterInterface
	clustersClientV1   v1clusters.ClustersInterface
	workspacesClientV1 v1workspaces.WorkspacesInterface
	projectsClientV1   v1projects.ProjectsInterface
	resourcesClientV1  v1resources.ResourceInterface
	metricsClientV1    v1metrics.MetricsInterface
	resourcesClientV2  v2resources.ResourcesInterface
	aclClientV1        v1Acl.AclInterface
	selfClientV2       v2self.SelfInterface
	streamClientV2     v2stream.StreamInterface
	apiSecret          string
	role               string
}

func NewRorMockTransport() *RorMockTransport {
	t := &RorMockTransport{
		status:             mocktransportstatus.NewMockTransportStatus(),
		infoClientV1:       mocktransportinfo.NewV1Client(),
		streamClientV1:     mocktransportstream.NewV1Client(),
		datacenterClientV1: mocktransportdatacenter.NewV1Client(),
		clustersClientV1:   mocktransportclusters.NewV1Client(),
		workspacesClientV1: mocktransportworkspaces.NewV1Client(),
		projectsClientV1:   mocktransportprojects.NewV1Client(),
		resourcesClientV1:  mocktransportresources.NewV1Client(),
		metricsClientV1:    mocktransportmetrics.NewV1Client(),
		resourcesClientV2:  mocktransportresourcesv2.NewV2Client(),
		aclClientV1:        mocktransportacl.NewV1Client(),
		selfClientV2:       mocktransportself.NewV2Client(),
		streamClientV2:     mocktransportstreamv2.NewV2Client(),
		apiSecret:          "mock-api-secret-12345",
		role:               "mock-role",
	}
	return t
}

func (t *RorMockTransport) Status() transportstatus.RorTransportStatus {
	return t.status
}

func (t *RorMockTransport) Stream() v1stream.StreamInterface {
	return t.streamClientV1
}

func (t *RorMockTransport) Info() v1info.InfoInterface {
	return t.infoClientV1
}

func (t *RorMockTransport) Datacenters() v1datacenter.DatacenterInterface {
	return t.datacenterClientV1
}

func (t *RorMockTransport) Clusters() v1clusters.ClustersInterface {
	return t.clustersClientV1
}

func (t *RorMockTransport) Self() v2self.SelfInterface {
	return t.selfClientV2
}

func (t *RorMockTransport) Workspaces() v1workspaces.WorkspacesInterface {
	return t.workspacesClientV1
}

func (t *RorMockTransport) Projects() v1projects.ProjectsInterface {
	return t.projectsClientV1
}

func (t *RorMockTransport) Resources() v1resources.ResourceInterface {
	return t.resourcesClientV1
}

func (t *RorMockTransport) Metrics() v1metrics.MetricsInterface {
	return t.metricsClientV1
}

func (t *RorMockTransport) ResourcesV2() v2resources.ResourcesInterface {
	return t.resourcesClientV2
}

func (t *RorMockTransport) Streamv2() v2stream.StreamInterface {
	return t.streamClientV2
}

func (t *RorMockTransport) AclV1() v1Acl.AclInterface {
	return t.aclClientV1
}

func (t *RorMockTransport) Ping() error {
	// Mock implementation - always return nil (success)
	return nil
}

func (t *RorMockTransport) GetApiSecret() string {
	return t.apiSecret
}

func (t *RorMockTransport) GetRole() string {
	return t.role
}

// SetApiSecret allows setting the API secret for testing purposes
func (t *RorMockTransport) SetApiSecret(secret string) {
	t.apiSecret = secret
}

// SetRole allows setting the role for testing purposes
func (t *RorMockTransport) SetRole(role string) {
	t.role = role
}
