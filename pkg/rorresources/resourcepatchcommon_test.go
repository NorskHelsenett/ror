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
)

func newTaggedCluster(t *testing.T) *rorresources.Resource {
	t.Helper()

	res := newTestKubernetesCluster(t)
	res.Metadata.Labels = map[string]string{"team": "platform", "stale": "yes"}
	res.Metadata.Annotations = map[string]string{"note": "keep me"}
	res.RorMeta.Tags = []rortypes.ResourceTag{
		{Key: "env", Value: "dev"},
		{Key: "deprecated", Value: "true"},
	}
	return res
}

func TestResourcePatchPatchesLabelsPerKey(t *testing.T) {
	in := newTaggedCluster(t)

	out, err := rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetLabel("team", "infra").
		SetLabel("new", "label").
		RemoveLabel("stale").
		Apply(in)
	require.NoError(t, err)

	assert.Equal(t, map[string]string{"team": "infra", "new": "label"}, out.Metadata.Labels)
	assert.Equal(t, map[string]string{"note": "keep me"}, out.Metadata.Annotations, "annotations are untouched")
	assert.Equal(t, map[string]string{"team": "platform", "stale": "yes"}, in.Metadata.Labels, "the input is untouched")
}

func TestResourcePatchPatchesAnnotationsPerKey(t *testing.T) {
	in := newTaggedCluster(t)

	out, err := rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetAnnotation("added", "yes").
		RemoveAnnotation("note").
		Apply(in)
	require.NoError(t, err)

	assert.Equal(t, map[string]string{"added": "yes"}, out.Metadata.Annotations)
	assert.Equal(t, map[string]string{"team": "platform", "stale": "yes"}, out.Metadata.Labels, "labels are untouched")
}

func TestResourcePatchPatchesTagsByKey(t *testing.T) {
	in := newTaggedCluster(t)

	out, err := rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetTag("env", "prod").
		SetTag("owner", "team-a").
		RemoveTag("deprecated").
		Apply(in)
	require.NoError(t, err)

	assert.Equal(t, []rortypes.ResourceTag{
		{Key: "env", Value: "prod"},
		{Key: "owner", Value: "team-a"},
	}, out.RorMeta.Tags, "an updated tag keeps its position and a new tag is appended")
}

func TestResourcePatchLeavesUnmentionedCommonFieldsAlone(t *testing.T) {
	in := newTaggedCluster(t)
	in.RorMeta.Version = "v1"
	in.RorMeta.Internal = true

	out, err := rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetTag("env", "prod").
		Apply(in)
	require.NoError(t, err)

	assert.Equal(t, "v1", out.RorMeta.Version)
	assert.True(t, out.RorMeta.Internal)
	assert.Equal(t, "a-cluster", out.Metadata.Name)
	assert.Equal(t, testClusterUID, string(out.Metadata.UID))
	assert.Equal(t, "the original description", out.KubernetesClusterResource.Spec.ClusterMetadata.Description)
}

func TestResourcePatchPatchesCommonOnAKindWithoutAMirror(t *testing.T) {
	in := rorresources.NewRorWorkspaceResource()
	in.Metadata.UID = "workspace-uid"
	in.Metadata.Labels = map[string]string{"keep": "yes"}

	out, err := rorresources.NewResourcePatch(
		rordefs.ResourceWorkspace.Kind,
		rordefs.ResourceWorkspace.APIVersion,
		"workspace-uid",
	).SetLabel("added", "yes").Apply(in)
	require.NoError(t, err)

	assert.Equal(t, map[string]string{"keep": "yes", "added": "yes"}, out.Metadata.Labels)
}

func TestResourcePatchCombinesCommonAndSubResource(t *testing.T) {
	in := newTaggedCluster(t)

	out, err := rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetTag("env", "prod").
		SetLabel("team", "infra").
		SetKubernetesCluster(clusterMetadataPatch(rortypes.ResourcePatchKubernetesClusterSpecMetadataDetails{
			Description: optional.NewOptional("a new description"),
		})).
		Apply(in)
	require.NoError(t, err)

	assert.Equal(t, "prod", mustTag(t, out, "env").Value)
	assert.Equal(t, "infra", out.Metadata.Labels["team"])
	assert.Equal(t, "a new description", out.KubernetesClusterResource.Spec.ClusterMetadata.Description)
}

func TestResourcePatchCommonJsonRoundTrip(t *testing.T) {
	patch := rorresources.NewResourcePatchKubernetesCluster(testClusterUID).
		SetLabel("team", "infra").
		RemoveLabel("stale").
		RemoveTag("deprecated")

	encoded, err := json.Marshal(patch)
	require.NoError(t, err)
	assert.NotContains(t, string(encoded), "annotations", "an unset field must not be encoded")
	assert.NotContains(t, string(encoded), "kubernetescluster", "an unset sub resource must not be encoded")

	var decoded rorresources.ResourcePatch
	require.NoError(t, json.Unmarshal(encoded, &decoded))

	out, err := decoded.Apply(newTaggedCluster(t))
	require.NoError(t, err)
	assert.Equal(t, map[string]string{"team": "infra"}, out.Metadata.Labels)
	assert.Equal(t, []rortypes.ResourceTag{{Key: "env", Value: "dev"}}, out.RorMeta.Tags)
}

func mustTag(t *testing.T, res *rorresources.Resource, key string) rortypes.ResourceTag {
	t.Helper()
	tag, ok := res.RorMeta.GetTagByKey(key)
	require.True(t, ok, "expected a tag with key %q", key)
	return tag
}
