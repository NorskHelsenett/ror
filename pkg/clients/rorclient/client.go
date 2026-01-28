package rorclient

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/clients"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports"
	v1acl "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/acl"
	v1clusters "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/clusters"
	v1datacenter "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/datacenter"
	v1info "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/info"
	v1metrics "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/metrics"
	v1projects "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/projects"
	v1resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/resources"
	v1stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/stream"
	v1token "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/token"
	v1workspaces "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v1/workspaces"
	v2clusters "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/clusters"
	v2resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/resources"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/rorclientv2self"
	v2token "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/token"
	v2stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/v2/v2stream"
	"github.com/NorskHelsenett/ror/pkg/helpers/rorhealth"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
)

type RorConfig struct {
	Host string
}

// Compile-time check to ensure RorClient implements RorClientInterface
var _ RorClientInterface = (*RorClient)(nil)

type RorClientInterface interface {
	GetRole() string
	GetApiSecret() string
	GetOwnerref() rorresourceowner.RorResourceOwnerReference
	SetOwnerref(ownerref rorresourceowner.RorResourceOwnerReference)
	CheckConnection() error

	Acl() v1acl.AclInterface
	Clusters() v1clusters.ClustersInterface
	ClustersV2() v2clusters.ClustersInterface
	Datacenters() v1datacenter.DatacenterInterface
	Info() v1info.InfoInterface
	Metrics() v1metrics.MetricsInterface
	Projects() v1projects.ProjectsInterface
	ResourceV2() v2resources.ResourcesInterface
	Resources() v1resources.ResourceInterface
	Self() rorclientv2self.SelfInterface
	SetTransport(transport transports.RorTransport)
	Stream() v1stream.StreamInterface
	StreamV2() v2stream.StreamInterface
	Workspaces() v1workspaces.WorkspacesInterface
	Token() v1token.TokenInterface
	TokenV2() v2token.TokenInterface

	clients.CommonClient
}

type RorClient struct {
	ownerRef *rorresourceowner.RorResourceOwnerReference

	Transport          transports.RorTransport
	streamClientV1     v1stream.StreamInterface
	infoClientV1       v1info.InfoInterface
	datacenterClientV1 v1datacenter.DatacenterInterface
	clustersClientV1   v1clusters.ClustersInterface
	clustersClientV2   v2clusters.ClustersInterface
	workspacesClientV1 v1workspaces.WorkspacesInterface
	selfClientV2       rorclientv2self.SelfInterface
	projectsClientV1   v1projects.ProjectsInterface
	resourceClientV1   v1resources.ResourceInterface
	metricsClientV1    v1metrics.MetricsInterface
	resourcesClientV2  v2resources.ResourcesInterface
	streamClientV2     v2stream.StreamInterface
	aclClientV1        v1acl.AclInterface
	tokenClientV1      v1token.TokenInterface
	tokenClientV2      v2token.TokenInterface
}

func NewRorClient(transport transports.RorTransport) *RorClient {
	return &RorClient{
		Transport:          transport,
		aclClientV1:        transport.Acl(),
		streamClientV1:     transport.Stream(),
		infoClientV1:       transport.Info(),
		datacenterClientV1: transport.Datacenters(),
		clustersClientV1:   transport.Clusters(),
		clustersClientV2:   transport.ClustersV2(),
		workspacesClientV1: transport.Workspaces(),
		projectsClientV1:   transport.Projects(),
		selfClientV2:       transport.Self(),
		resourceClientV1:   transport.Resources(),
		metricsClientV1:    transport.Metrics(),
		resourcesClientV2:  transport.ResourcesV2(),
		streamClientV2:     transport.Streamv2(),
		tokenClientV1:      transport.Token(),
		tokenClientV2:      transport.TokenV2(),
	}
}

func (c *RorClient) SetTransport(transport transports.RorTransport) {
	c.Transport = transport
}
func (c *RorClient) Stream() v1stream.StreamInterface {
	return c.streamClientV1
}

func (c *RorClient) Acl() v1acl.AclInterface {
	return c.aclClientV1
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

func (c *RorClient) ClustersV2() v2clusters.ClustersInterface {
	return c.clustersClientV2
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

func (c *RorClient) Token() v1token.TokenInterface {
	return c.tokenClientV1
}

func (c *RorClient) TokenV2() v2token.TokenInterface {
	return c.tokenClientV2
}

func (c *RorClient) CheckConnection() error {
	return c.Transport.CheckConnection()
}

// Ping checks the connection to the transport.
// Old version used error handling, use CheckConnection instead.
func (c *RorClient) Ping() bool {
	return c.PingWithContext(context.TODO())
}

func (c *RorClient) PingWithContext(ctx context.Context) bool {
	return c.Transport.Ping(ctx)
}

func (c *RorClient) ResourceV2() v2resources.ResourcesInterface {
	return c.resourcesClientV2
}

func (c *RorClient) StreamV2() v2stream.StreamInterface {
	return c.streamClientV2
}

func (c *RorClient) GetRole() string {
	return c.Transport.GetRole()
}

func (c *RorClient) GetApiSecret() string {
	return c.Transport.GetApiSecret()
}

func (c *RorClient) GetOwnerref() rorresourceowner.RorResourceOwnerReference {
	if c.ownerRef == nil {
		return rorresourceowner.RorResourceOwnerReference{Scope: aclmodels.Acl2ScopeUnknown, Subject: aclmodels.Acl2RorSubjecUnknown}
	}
	return *c.ownerRef
}

func (c *RorClient) SetOwnerref(ownerref rorresourceowner.RorResourceOwnerReference) {
	c.ownerRef = &ownerref
}

func (c *RorClient) CheckHealth(ctx context.Context) []rorhealth.Check {
	healthChecks := []rorhealth.Check{}
	if !c.Transport.Ping(ctx) {
		healthChecks = append(healthChecks, rorhealth.Check{
			ComponentID: "Transport",
			Status:      rorhealth.StatusFail,
			Output:      fmt.Sprintf("%s could not be connected", c.Transport.GetTransportName()),
		})
	}
	return healthChecks
}

func (c *RorClient) CheckHealthWithoutContext() []rorhealth.Check {
	healthChecks := []rorhealth.Check{}
	if !c.Transport.Ping(context.Background()) {
		healthChecks = append(healthChecks, rorhealth.Check{
			ComponentID: "Transport",
			Status:      rorhealth.StatusFail,
			Output:      fmt.Sprintf("%s could not be connected", c.Transport.GetTransportName()),
		})
	}
	return healthChecks
}
