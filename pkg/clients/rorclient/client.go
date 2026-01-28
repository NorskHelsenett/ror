package rorclient

import (
	"context"
	"fmt"

	"github.com/NorskHelsenett/ror/pkg/clients"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/clientinterface"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/transports/transportinterface"
	v1acl "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/acl"
	v1clusters "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/clusters"
	v1datacenter "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/datacenter"
	v1info "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/info"
	v1metrics "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/metrics"
	v1projects "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/projects"
	v1resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/resources"
	v1stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/stream"
	v1token "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/token"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/v1clientset"
	v1workspaces "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/workspaces"
	v2apikeys "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/apikeys"
	v2resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/resources"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/rorclientv2self"
	v2token "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/token"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/v2clientset"
	v2stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/v2stream"

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
	transportinterface.RorCommonClientTransportInterface
	clientinterface.RorCommonClientApiInterface

	clientinterface.RorCommonClientOwnerInterface
	transportinterface.RorCommonClientTransportSetterInterface

	clients.CommonClient
}

type RorClient struct {
	ownerRef *rorresourceowner.RorResourceOwnerReference

	Transport          transportinterface.RorTransport
	streamClientV1     v1stream.StreamInterface
	infoClientV1       v1info.InfoInterface
	datacenterClientV1 v1datacenter.DatacenterInterface
	clustersClientV1   v1clusters.ClustersInterface
	apikeysClientV2    v2apikeys.ApiKeysInterface
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
	v1                 clientinterface.RorCommonClientApiInterfaceV1
	v2                 clientinterface.RorCommonClientApiInterfaceV2
}

func NewRorClient(transport transportinterface.RorTransport) *RorClient {
	return &RorClient{
		Transport:          transport,
		aclClientV1:        transport.Acl(),
		streamClientV1:     transport.Stream(),
		infoClientV1:       transport.Info(),
		datacenterClientV1: transport.Datacenters(),
		clustersClientV1:   transport.Clusters(),
		apikeysClientV2:    transport.ApiKeysV2(),
		workspacesClientV1: transport.Workspaces(),
		projectsClientV1:   transport.Projects(),
		selfClientV2:       transport.Self(),
		resourceClientV1:   transport.Resources(),
		metricsClientV1:    transport.Metrics(),
		resourcesClientV2:  transport.ResourcesV2(),
		streamClientV2:     transport.StreamV2(),
		tokenClientV1:      transport.Token(),
		tokenClientV2:      transport.TokenV2(),
		v1:                 v1clientset.NewV1ClientSet(transport),
		v2:                 v2clientset.NewV2ClientSet(transport),
	}
}

func (c *RorClient) V1() clientinterface.RorCommonClientApiInterfaceV1 {
	return c.v1
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

func (c *RorClient) ApiKeysV2() v2apikeys.ApiKeysInterface {
	return c.apikeysClientV2
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

func (c *RorClient) ResourcesV2() v2resources.ResourcesInterface {
	return c.resourcesClientV2
}

func (c *RorClient) StreamV2() v2stream.StreamInterface {
	return c.streamClientV2
}

// Ping checks the connection to the transport.
// Old version used error handling, use CheckConnection instead.
func (c *RorClient) Ping() bool {
	return c.PingWithContext(context.TODO())
}

// PingWithContext checks the connection to the transport with a context.
func (c *RorClient) PingWithContext(ctx context.Context) bool {
	return c.Transport.Ping(ctx)
}

// Transport related methods

// CheckConnection checks the connection to the transport.
func (c *RorClient) CheckConnection() error {
	return c.Transport.CheckConnection()
}

func (c *RorClient) GetRole() string {
	return c.Transport.GetRole()
}

// GetApiSecret gets the API secret from the transport.
func (c *RorClient) GetApiSecret() string {
	return c.Transport.GetApiSecret()
}

// SetTransport sets the transport for the RorClient.
func (c *RorClient) SetTransport(transport transportinterface.RorTransport) {
	c.Transport = transport
}

// GetOwnerref gets the owner reference for the RorClient.
func (c *RorClient) GetOwnerref() rorresourceowner.RorResourceOwnerReference {
	if c.ownerRef == nil {
		return rorresourceowner.RorResourceOwnerReference{Scope: aclmodels.Acl2ScopeUnknown, Subject: aclmodels.Acl2RorSubjecUnknown}
	}
	return *c.ownerRef
}

// SetOwnerref sets the owner reference for the RorClient.
func (c *RorClient) SetOwnerref(ownerref rorresourceowner.RorResourceOwnerReference) {
	c.ownerRef = &ownerref
}

// CheckHealth checks the health of the RorClient.
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

// CheckHealthWithoutContext checks the health of the RorClient without a context.
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
