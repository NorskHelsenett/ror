package clientinterface

import (
	v1acl "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v1/acl"
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
	"github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/rorclientv2self"
	v2token "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/token"
	v2stream "github.com/NorskHelsenett/ror/pkg/clients/rorclient/interfaces/v2/v2stream"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/rorresourceowner"
)

type RorCommonClientApiInterface interface {
	Acl() v1acl.AclInterface
	ApiKeysV2() v2apikeys.ApiKeysInterface
	Clusters() v1clusters.ClustersInterface
	Datacenters() v1datacenter.DatacenterInterface
	Info() v1info.InfoInterface
	Metrics() v1metrics.MetricsInterface
	Projects() v1projects.ProjectsInterface
	Resources() v1resources.ResourceInterface
	ResourcesV2() v2resources.ResourcesInterface
	Self() rorclientv2self.SelfInterface
	Stream() v1stream.StreamInterface
	StreamV2() v2stream.StreamInterface
	Token() v1token.TokenInterface
	TokenV2() v2token.TokenInterface
	Workspaces() v1workspaces.WorkspacesInterface
}

type RorCommonClientOwnerInterface interface {
	GetOwnerref() rorresourceowner.RorResourceOwnerReference
	SetOwnerref(ownerref rorresourceowner.RorResourceOwnerReference)
}
