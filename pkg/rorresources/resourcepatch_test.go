package rorresources_test

import (
	"encoding/json"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/helpers/optional"
	"github.com/NorskHelsenett/ror/pkg/rorresources"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rordefs"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rortypes"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

const testClusterUID = "cluster-uid-1"

func newTestKubernetesCluster(t *testing.T) *rorresources.Resource {
	t.Helper()

	res := rorresources.NewRorKubernetesClusterResource()
	res.Metadata = metav1.ObjectMeta{Name: "a-cluster", UID: testClusterUID}
	res.KubernetesClusterResource.Spec.ClusterMetadata = rortypes.ResourceKubernetesClusterSpecMetadataDetails{
		Criticality:   "high",
		Sensitivity:   "internal",
		Description:   "the original description",
		SlackChannels: []string{"#ror"},
		Contacts:      []rortypes.ResourceContact{{Name: "owner", Email: "owner@example.com"}},
	}
	res.KubernetesClusterResource.Status.AgentStatus.ClusterId = "agent-cluster-id"
	res.KubernetesClusterResource.Status.AgentStatus.Region = "osl"
	res.KubernetesClusterResource.Status.AgentStatus.Endpoint.ApiServer = "https://api.example.com"
	res.KubernetesClusterResource.Status.AgentStatus.Endpoint.CACert = "a-cert"

	out, err := rorresources.NewResourceFromStruct(*res)
	require.NoError(t, err)
	return out
}

func clusterMetadataPatch(patch rortypes.ResourcePatchKubernetesClusterSpecMetadataDetails) rortypes.ResourcePatchKubernetesCluster {
	return rortypes.ResourcePatchKubernetesCluster{
		Spec: optional.NewOptional(rortypes.ResourcePatchKubernetesClusterSpec{
			ClusterMetadata: optional.NewOptional(patch),
		}),
	}
}

func TestResourcePatchAppliesTheThreeStates(t *testing.T) {
	in := newTestKubernetesCluster(t)

	patch := rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetKubernetesCluster(clusterMetadataPatch(rortypes.ResourcePatchKubernetesClusterSpecMetadataDetails{
			Description:   optional.NewOptional("a new description"),
			SlackChannels: optional.Null[[]string](),
		}))

	out, err := patch.Apply(in)
	require.NoError(t, err)

	metadata := out.KubernetesClusterResource.Spec.ClusterMetadata
	assert.Equal(t, "a new description", metadata.Description, "a value should be written")
	assert.Nil(t, metadata.SlackChannels, "null should clear the field")
	assert.Equal(t, "high", metadata.Criticality, "an unset field should be left untouched")
	assert.Equal(t, "internal", metadata.Sensitivity, "an unset field should be left untouched")
	assert.Len(t, metadata.Contacts, 1, "an unset field should be left untouched")
}

func TestResourcePatchLeavesTheInputResourceUntouched(t *testing.T) {
	in := newTestKubernetesCluster(t)

	patch := rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetKubernetesCluster(clusterMetadataPatch(rortypes.ResourcePatchKubernetesClusterSpecMetadataDetails{
			Description: optional.NewOptional("a new description"),
		}))

	out, err := patch.Apply(in)
	require.NoError(t, err)

	assert.Equal(t, "the original description", in.KubernetesClusterResource.Spec.ClusterMetadata.Description)
	assert.NotSame(t, in.KubernetesClusterResource, out.KubernetesClusterResource)
}

func TestResourcePatchNullOnNestedMirrorResetsTheWholeStruct(t *testing.T) {
	in := newTestKubernetesCluster(t)

	patch := rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetKubernetesCluster(rortypes.ResourcePatchKubernetesCluster{
			Spec: optional.NewOptional(rortypes.ResourcePatchKubernetesClusterSpec{
				ClusterMetadata: optional.Null[rortypes.ResourcePatchKubernetesClusterSpecMetadataDetails](),
			}),
		})

	out, err := patch.Apply(in)
	require.NoError(t, err)

	assert.Equal(t, rortypes.ResourceKubernetesClusterSpecMetadataDetails{}, out.KubernetesClusterResource.Spec.ClusterMetadata)
	assert.Equal(t, "agent-cluster-id", out.KubernetesClusterResource.Status.AgentStatus.ClusterId, "a sibling branch should be left untouched")
}

func TestResourcePatchAppliesDeeplyNestedFields(t *testing.T) {
	in := newTestKubernetesCluster(t)

	patch := rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetKubernetesCluster(rortypes.ResourcePatchKubernetesCluster{
			Status: optional.NewOptional(rortypes.ResourcePatchKubernetesClusterStatus{
				AgentStatus: optional.NewOptional(rortypes.ResourcePatchKubernetesClusterAgentStatus{
					Endpoint: optional.NewOptional(rortypes.ResourcePatchKubernetesClusterAgentStatusEndpoint{
						ApiServer: optional.NewOptional("https://new.example.com"),
					}),
				}),
			}),
		})

	out, err := patch.Apply(in)
	require.NoError(t, err)

	endpoint := out.KubernetesClusterResource.Status.AgentStatus.Endpoint
	assert.Equal(t, "https://new.example.com", endpoint.ApiServer)
	assert.Equal(t, "a-cert", endpoint.CACert, "an unset sibling should be left untouched")
	assert.Equal(t, "osl", out.KubernetesClusterResource.Status.AgentStatus.Region)
	assert.Equal(t, "the original description", out.KubernetesClusterResource.Spec.ClusterMetadata.Description)
}

func TestResourcePatchNullSubResourceResetsIt(t *testing.T) {
	in := newTestKubernetesCluster(t)

	patch := rorresources.NewResourcePatchKubernetesCluster(testClusterUID)
	patch.KubernetesClusterResource.SetNull()

	out, err := patch.Apply(in)
	require.NoError(t, err)

	assert.Equal(t, rortypes.ResourceKubernetesCluster{}, *out.KubernetesClusterResource)
	assert.Equal(t, testClusterUID, string(out.Metadata.UID), "the common part should be kept")
}

func TestResourcePatchKeepsTheCommonInterfaceUsable(t *testing.T) {
	in := newTestKubernetesCluster(t)

	patch := rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetKubernetesCluster(clusterMetadataPatch(rortypes.ResourcePatchKubernetesClusterSpecMetadataDetails{
			Description: optional.NewOptional("a new description"),
		}))

	out, err := patch.Apply(in)
	require.NoError(t, err)

	out.GenRorHash()
	assert.NotEmpty(t, out.RorMeta.Hash)
	assert.NotEqual(t, in.GetRorHash(), out.GetRorHash(), "a changed field should change the hash")
}

func TestResourcePatchValidate(t *testing.T) {
	validSubResource := clusterMetadataPatch(rortypes.ResourcePatchKubernetesClusterSpecMetadataDetails{
		Description: optional.NewOptional("a new description"),
	})

	tests := []struct {
		patch   *rorresources.ResourcePatch
		wantErr error
		name    string
	}{
		{
			name:    "nil patch",
			patch:   nil,
			wantErr: rorresources.ErrPatchNil,
		},
		{
			name:    "empty uid",
			patch:   rorresources.NewResourcePatchKubernetesCluster("").SetKubernetesCluster(validSubResource),
			wantErr: rorresources.ErrPatchEmptyUID,
		},
		{
			name:    "no sub resource",
			patch:   rorresources.NewResourcePatchKubernetesCluster(testClusterUID),
			wantErr: rorresources.ErrPatchEmpty,
		},
		{
			name: "sub resource of another kind",
			patch: rorresources.NewResourcePatch(
				rordefs.ResourceWorkspace.Kind,
				rordefs.ResourceWorkspace.APIVersion,
				testClusterUID,
			).SetKubernetesCluster(validSubResource),
			wantErr: rorresources.ErrPatchKindNotAllowed,
		},
		{
			name: "common only patch is valid on a kind without a mirror",
			patch: rorresources.NewResourcePatch(
				rordefs.ResourceWorkspace.Kind,
				rordefs.ResourceWorkspace.APIVersion,
				testClusterUID,
			).SetTag("env", "test"),
			wantErr: nil,
		},
		{
			name:    "valid",
			patch:   rorresources.NewResourcePatchKubernetesCluster(testClusterUID).SetKubernetesCluster(validSubResource),
			wantErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.patch.Validate()
			if test.wantErr == nil {
				assert.NoError(t, err)
				return
			}
			assert.ErrorIs(t, err, test.wantErr)
		})
	}
}

func TestResourcePatchApplyRejectsMismatchedResources(t *testing.T) {
	in := newTestKubernetesCluster(t)
	validSubResource := clusterMetadataPatch(rortypes.ResourcePatchKubernetesClusterSpecMetadataDetails{
		Description: optional.NewOptional("a new description"),
	})

	_, err := rorresources.NewResourcePatchKubernetesCluster("another-uid").
		SetKubernetesCluster(validSubResource).
		Apply(in)
	assert.ErrorIs(t, err, rorresources.ErrPatchUIDMismatch)

	_, err = rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetKubernetesCluster(validSubResource).
		Apply(rorresources.NewRorWorkspaceResource())
	assert.ErrorIs(t, err, rorresources.ErrPatchKindMismatch)

	_, err = rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetKubernetesCluster(validSubResource).
		Apply(nil)
	assert.ErrorIs(t, err, rorresources.ErrPatchNilResource)
}

func TestResourcePatchJsonRoundTripKeepsTheThreeStates(t *testing.T) {
	patch := rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetKubernetesCluster(clusterMetadataPatch(rortypes.ResourcePatchKubernetesClusterSpecMetadataDetails{
			Description:   optional.NewOptional("a new description"),
			SlackChannels: optional.Null[[]string](),
		}))

	encoded, err := json.Marshal(patch)
	require.NoError(t, err)

	assert.NotContains(t, string(encoded), "criticality", "an unset field must not be encoded")
	assert.Contains(t, string(encoded), `"slackChannels":null`, "a null field must be encoded as null")
	assert.NotContains(t, string(encoded), "status", "an unset sub tree must not be encoded")

	var decoded rorresources.ResourcePatch
	require.NoError(t, json.Unmarshal(encoded, &decoded))

	cluster, ok := decoded.KubernetesCluster()
	require.True(t, ok)
	spec, ok := cluster.Spec.Value()
	require.True(t, ok)
	metadata, ok := spec.ClusterMetadata.Value()
	require.True(t, ok)

	assert.Equal(t, "a new description", metadata.Description.ValueOr(""))
	assert.True(t, metadata.SlackChannels.IsNull())
	assert.False(t, metadata.Criticality.IsSet())
	assert.False(t, cluster.Status.IsSet())

	in := newTestKubernetesCluster(t)
	out, err := decoded.Apply(in)
	require.NoError(t, err)
	assert.Equal(t, "a new description", out.KubernetesClusterResource.Spec.ClusterMetadata.Description)
	assert.Nil(t, out.KubernetesClusterResource.Spec.ClusterMetadata.SlackChannels)
	assert.Equal(t, "high", out.KubernetesClusterResource.Spec.ClusterMetadata.Criticality)
}

// A second kind is covered to show the dispatcher routes by kind rather than
// only working for the one kind it was first written against.
func TestResourcePatchAppliesToAnotherKind(t *testing.T) {
	res := rorresources.NewRorOrganizationalUnitResource()
	res.Metadata = metav1.ObjectMeta{Name: "an-ou", UID: "ou-uid-1"}
	res.OrganizationalUnitResource.ID = "ou-1"
	res.OrganizationalUnitResource.Spec.Name = "the original name"
	res.OrganizationalUnitResource.Spec.Type = rortypes.OrganizationalUnitTypeProject

	in, err := rorresources.NewResourceFromStruct(*res)
	require.NoError(t, err)

	patch := rorresources.NewResourcePatchOrganizationalUnit("ou-uid-1").
		SetOrganizationalUnit(rortypes.ResourcePatchOrganizationalUnit{
			Spec: optional.NewOptional(rortypes.ResourcePatchOrganizationalUnitSpec{
				Name: optional.NewOptional("a new name"),
			}),
		})

	out, err := patch.Apply(in)
	require.NoError(t, err)

	assert.Equal(t, "a new name", out.OrganizationalUnitResource.Spec.Name)
	assert.Equal(t, rortypes.OrganizationalUnitTypeProject, out.OrganizationalUnitResource.Spec.Type, "an unset field should be left untouched")
	assert.Equal(t, "ou-1", out.OrganizationalUnitResource.ID, "an unset field should be left untouched")
	assert.Equal(t, "the original name", in.OrganizationalUnitResource.Spec.Name, "the input must not be mutated")
}
