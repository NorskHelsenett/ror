// THIS FILE IS GENERATED, DO NOT EDIT
// ref: cmd/generator/main.go

package rortypes

import (
	"time"

	"github.com/NorskHelsenett/ror/pkg/helpers/optional"
	"github.com/NorskHelsenett/ror/pkg/kubernetes/providers/providermodels"
	commonv1alpha1 "github.com/vitistack/common/pkg/v1alpha1"
)

// This file holds a patch mirror of a resource. a patch mirror mirrors every field
// but the field are a generic type of optional.Optional, which carries the three
// states a partial update needs: unset, null, and a value.
// See the optional package for why every field is tagged omitzero.

// ResourcePatchKubernetesCluster is the patch mirror of ResourceKubernetesCluster.
type ResourcePatchKubernetesCluster struct {
	Spec   optional.Optional[ResourcePatchKubernetesClusterSpec]   `json:"spec,omitzero"`
	Status optional.Optional[ResourcePatchKubernetesClusterStatus] `json:"status,omitzero"`
}

// ApplyTo writes the patch into target, leaving the fields the patch does
// not mention untouched.
func (p ResourcePatchKubernetesCluster) ApplyTo(target *ResourceKubernetesCluster) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	if err := optional.ApplyNested(p.Spec, &target.Spec); err != nil {
		return err
	}
	if err := optional.ApplyNested(p.Status, &target.Status); err != nil {
		return err
	}
	return nil
}

// ResourcePatchKubernetesClusterAgentStatus is the patch mirror of KubernetesClusterAgentStatus.
type ResourcePatchKubernetesClusterAgentStatus struct {
	ClusterId          optional.Optional[string]                                            `json:"clusterId,omitzero"`
	ClusterName        optional.Optional[string]                                            `json:"clusterName,omitzero"`
	KubernetesProvider optional.Optional[providermodels.ProviderType]                       `json:"kubernetesProvider,omitzero"`
	Az                 optional.Optional[string]                                            `json:"az,omitzero"`
	Region             optional.Optional[string]                                            `json:"region,omitzero"`
	Country            optional.Optional[string]                                            `json:"country,omitzero"`
	Workspace          optional.Optional[string]                                            `json:"workspaceId,omitzero"`
	Environment        optional.Optional[string]                                            `json:"environment,omitzero"`
	Datacenter         optional.Optional[string]                                            `json:"datacenter,omitzero"`
	Nodes              optional.Optional[ResourcePatchKubernetesClusterAgentStatusNodes]    `json:"nodes,omitzero"`
	Versions           optional.Optional[optional.Map[string, string]]                      `json:"versions,omitzero"`
	Urls               optional.Optional[optional.Map[string, string]]                      `json:"urls,omitzero"`
	Endpoint           optional.Optional[ResourcePatchKubernetesClusterAgentStatusEndpoint] `json:"endpoint,omitzero"`
	CreatedAt          optional.Optional[time.Time]                                         `json:"createdAt,omitzero"`
	LastSeen           optional.Optional[time.Time]                                         `json:"lastSeen,omitzero"`
}

// ApplyTo writes the patch into target, leaving the fields the patch does
// not mention untouched.
func (p ResourcePatchKubernetesClusterAgentStatus) ApplyTo(target *KubernetesClusterAgentStatus) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	optional.Apply(p.ClusterId, &target.ClusterId)
	optional.Apply(p.ClusterName, &target.ClusterName)
	optional.Apply(p.KubernetesProvider, &target.KubernetesProvider)
	optional.Apply(p.Az, &target.Az)
	optional.Apply(p.Region, &target.Region)
	optional.Apply(p.Country, &target.Country)
	optional.Apply(p.Workspace, &target.Workspace)
	optional.Apply(p.Environment, &target.Environment)
	optional.Apply(p.Datacenter, &target.Datacenter)
	if err := optional.ApplyNested(p.Nodes, &target.Nodes); err != nil {
		return err
	}
	if err := optional.ApplyNested(p.Versions, &target.Versions); err != nil {
		return err
	}
	if err := optional.ApplyNested(p.Urls, &target.Urls); err != nil {
		return err
	}
	if err := optional.ApplyNested(p.Endpoint, &target.Endpoint); err != nil {
		return err
	}
	optional.Apply(p.CreatedAt, &target.CreatedAt)
	optional.Apply(p.LastSeen, &target.LastSeen)
	return nil
}

// ResourcePatchKubernetesClusterAgentStatusEndpoint is the patch mirror of KubernetesClusterAgentStatusEndpoint.
type ResourcePatchKubernetesClusterAgentStatusEndpoint struct {
	ApiServer optional.Optional[string] `json:"apiServer,omitzero"`
	CACert    optional.Optional[string] `json:"caCert,omitzero"`
	EgressIp  optional.Optional[string] `json:"egress,omitzero"`
}

// ApplyTo writes the patch into target, leaving the fields the patch does
// not mention untouched.
func (p ResourcePatchKubernetesClusterAgentStatusEndpoint) ApplyTo(target *KubernetesClusterAgentStatusEndpoint) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	optional.Apply(p.ApiServer, &target.ApiServer)
	optional.Apply(p.CACert, &target.CACert)
	optional.Apply(p.EgressIp, &target.EgressIp)
	return nil
}

// ResourcePatchKubernetesClusterAgentStatusNodes is the patch mirror of KubernetesClusterAgentStatusNodes.
type ResourcePatchKubernetesClusterAgentStatusNodes struct {
	ControllPlane optional.Optional[[]KubernetesClusterAgentStatusNodesNodepoolsNodes] `json:"controlPlane,omitzero"`
	Nodepools     optional.Optional[[]KubernetesClusterAgentStatusNodesNodepools]      `json:"nodepools,omitzero"`
}

// ApplyTo writes the patch into target, leaving the fields the patch does
// not mention untouched.
func (p ResourcePatchKubernetesClusterAgentStatusNodes) ApplyTo(target *KubernetesClusterAgentStatusNodes) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	optional.Apply(p.ControllPlane, &target.ControllPlane)
	optional.Apply(p.Nodepools, &target.Nodepools)
	return nil
}

// ResourcePatchKubernetesClusterSpec is the patch mirror of ResourceKubernetesClusterSpec.
type ResourcePatchKubernetesClusterSpec struct {
	ClusterMetadata optional.Optional[ResourcePatchKubernetesClusterSpecMetadataDetails] `json:"clusterMetadata,omitzero"`
	VitiSpec        optional.Optional[commonv1alpha1.KubernetesClusterSpec]              `json:"vitiSpec,omitzero"`
}

// ApplyTo writes the patch into target, leaving the fields the patch does
// not mention untouched.
func (p ResourcePatchKubernetesClusterSpec) ApplyTo(target *ResourceKubernetesClusterSpec) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	if err := optional.ApplyNested(p.ClusterMetadata, &target.ClusterMetadata); err != nil {
		return err
	}
	optional.Apply(p.VitiSpec, &target.VitiSpec)
	return nil
}

// ResourcePatchKubernetesClusterSpecMetadataDetails is the patch mirror of ResourceKubernetesClusterSpecMetadataDetails.
type ResourcePatchKubernetesClusterSpecMetadataDetails struct {
	Criticality   optional.Optional[string]            `json:"criticality,omitzero"`
	Sensitivity   optional.Optional[string]            `json:"sensitivity,omitzero"`
	Description   optional.Optional[string]            `json:"description,omitzero"`
	Contacts      optional.Optional[[]ResourceContact] `json:"contacts,omitzero"`
	SlackChannels optional.Optional[[]string]          `json:"slackChannels,omitzero"`
}

// ApplyTo writes the patch into target, leaving the fields the patch does
// not mention untouched.
func (p ResourcePatchKubernetesClusterSpecMetadataDetails) ApplyTo(target *ResourceKubernetesClusterSpecMetadataDetails) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	optional.Apply(p.Criticality, &target.Criticality)
	optional.Apply(p.Sensitivity, &target.Sensitivity)
	optional.Apply(p.Description, &target.Description)
	optional.Apply(p.Contacts, &target.Contacts)
	optional.Apply(p.SlackChannels, &target.SlackChannels)
	return nil
}

// ResourcePatchKubernetesClusterStatus is the patch mirror of ResourceKubernetesClusterStatus.
type ResourcePatchKubernetesClusterStatus struct {
	ProviderStatus optional.Optional[commonv1alpha1.KubernetesClusterStatus]    `json:"providerstatus,omitzero"`
	AgentStatus    optional.Optional[ResourcePatchKubernetesClusterAgentStatus] `json:"agentstatus,omitzero"`
}

// ApplyTo writes the patch into target, leaving the fields the patch does
// not mention untouched.
func (p ResourcePatchKubernetesClusterStatus) ApplyTo(target *ResourceKubernetesClusterStatus) error {
	if target == nil {
		return ErrPatchNilTarget
	}
	optional.Apply(p.ProviderStatus, &target.ProviderStatus)
	if err := optional.ApplyNested(p.AgentStatus, &target.AgentStatus); err != nil {
		return err
	}
	return nil
}

// KubernetesClusterPatchinterface represents the interface for patches of the type kubernetescluster,
// implemented by ResourcePatchKubernetesCluster.
type KubernetesClusterPatchinterface interface {
	ApplyTo(target *ResourceKubernetesCluster) error
}
