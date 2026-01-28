package resttransport

import (
	"context"
	"net/http"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/transports/transportstatusinterface"
	httpclient "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/httpclient"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/sseclient/v1sseclient"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/sseclient/v2sseclient"

	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/transports/transportinterface"
	v1Acl "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/acl"
	v1clusters "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/clusters"
	v1datacenter "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/datacenter"
	v1info "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/info"
	v1metrics "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/metrics"
	v1projects "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/projects"
	v1resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/resources"
	v1stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/stream"
	v1token "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/token"
	v1workspaces "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/workspaces"
	v2apikeys "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/apikeys"
	v2resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/resources"
	v2self "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/rorclientv2self"
	v2token "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/token"
	v2stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/v2stream"
	restv1Acl "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/acl"
	restv1clusters "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/clusters"
	restv1datacenter "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/datacenter"
	restv1info "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/info"
	restv1metrics "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/metrics"
	restv1projects "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/projects"
	restv1resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/resources"
	restv1token "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/token"
	restv1stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/v1stream"
	restv1workspaces "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v1/workspaces"
	restv2apikeys "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v2/apikeys"
	restv2resources "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v2/resources"
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v2/restclientv2self"
	restv2token "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v2/token"
	restv2stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/transports/resttransport/v2/v2stream"
)

// Compile-time check to ensure resourcecache implements ResourceCacheInterface
var _ transportinterface.RorTransport = (*RorHttpTransport)(nil)

type RorHttpTransport struct {
	Client *httpclient.HttpTransportClient

	streamClientV1     v1stream.StreamInterface
	infoClientV1       v1info.InfoInterface
	datacenterClientV1 v1datacenter.DatacenterInterface
	clustersClientV1   v1clusters.ClustersInterface
	apikeysClientV2    v2apikeys.ApiKeysInterface
	workspacesClientV1 v1workspaces.WorkspacesInterface
	projectsClientV1   v1projects.ProjectsInterface
	resourcesClientV1  v1resources.ResourceInterface
	metricsClientV1    v1metrics.MetricsInterface
	resourcesClientV2  v2resources.ResourcesInterface
	aclClientV1        v1Acl.AclInterface
	selfClientV2       v2self.SelfInterface
	streamClientV2     v2stream.StreamInterface
	tokenClientV1      v1token.TokenInterface
	tokenClientV2      v2token.TokenInterface
}

func NewRorHttpTransport(config *httpclient.HttpTransportClientConfig) *RorHttpTransport {
	httpClient := &http.Client{}
	return newWithHttpClient(config, httpClient)

}

func NewWithCustomHttpClient(config *httpclient.HttpTransportClientConfig, httpClient *http.Client) *RorHttpTransport {
	return newWithHttpClient(config, httpClient)
}

func newWithHttpClient(config *httpclient.HttpTransportClientConfig, httpClient *http.Client) *RorHttpTransport {
	client := &httpclient.HttpTransportClient{
		Client: httpClient,
		Config: config,
		Status: httpclient.NewHttpTransportClientStatus(),
	}
	t := &RorHttpTransport{
		Client:             client,
		streamClientV1:     restv1stream.NewV1Client(v1sseclient.NewSSEClient(client)),
		infoClientV1:       restv1info.NewV1Client(client),
		datacenterClientV1: restv1datacenter.NewV1Client(client),
		clustersClientV1:   restv1clusters.NewV1Client(client),
		apikeysClientV2:    restv2apikeys.NewV2Client(client),
		selfClientV2:       restclientv2self.NewV2Client(client),
		workspacesClientV1: restv1workspaces.NewV1Client(client),
		projectsClientV1:   restv1projects.NewV1Client(client),
		resourcesClientV1:  restv1resources.NewV1Client(client),
		metricsClientV1:    restv1metrics.NewV1Client(client),
		resourcesClientV2:  restv2resources.NewV2Client(client),
		streamClientV2:     restv2stream.NewV2Client(v2sseclient.NewSSEClient(client)),
		aclClientV1:        restv1Acl.NewV1Client(client),
		tokenClientV1:      restv1token.NewV1Client(client),
		tokenClientV2:      restv2token.NewV2Client(client),
	}
	return t
}

func (t *RorHttpTransport) Status() transportstatusinterface.RorTransportStatus {
	return t.Client.Status
}

func (t *RorHttpTransport) Stream() v1stream.StreamInterface {
	return t.streamClientV1
}

func (t *RorHttpTransport) Info() v1info.InfoInterface {
	return t.infoClientV1
}

func (t *RorHttpTransport) Acl() v1Acl.AclInterface {
	return t.aclClientV1
}

func (t *RorHttpTransport) Datacenters() v1datacenter.DatacenterInterface {
	return t.datacenterClientV1
}

func (t *RorHttpTransport) Clusters() v1clusters.ClustersInterface {
	return t.clustersClientV1
}

func (t *RorHttpTransport) ApiKeysV2() v2apikeys.ApiKeysInterface {
	return t.apikeysClientV2
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

func (t *RorHttpTransport) AclV1() v1Acl.AclInterface {
	return t.aclClientV1
}

func (t *RorHttpTransport) Token() v1token.TokenInterface {
	return t.tokenClientV1
}

func (t *RorHttpTransport) TokenV2() v2token.TokenInterface {
	return t.tokenClientV2
}

func (t *RorHttpTransport) StreamV2() v2stream.StreamInterface {
	return t.streamClientV2
}

func (t *RorHttpTransport) Self() v2self.SelfInterface {
	return t.selfClientV2
}

func (t *RorHttpTransport) CheckConnection() error {
	_, err := t.infoClientV1.GetVersion(context.Background())
	return err
}

func (t *RorHttpTransport) Ping(ctx context.Context) bool {
	_, err := t.infoClientV1.GetVersion(ctx)
	return err == nil
}

func (t *RorHttpTransport) GetApiSecret() string {
	return t.Client.Config.AuthProvider.GetApiSecret()
}

func (t *RorHttpTransport) GetRole() string {
	return t.Client.Config.GetRole()
}

func (t *RorHttpTransport) GetTransportName() string {
	return "HTTP Transport"
}
