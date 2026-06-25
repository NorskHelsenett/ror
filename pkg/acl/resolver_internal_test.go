package acl

import (
	"testing"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"

	"github.com/stretchr/testify/assert"
)

// skipExpansion is a narrowing optimisation: it may only return true when no
// descendant could ever pass the filter. Returning true incorrectly would drop
// legitimately-matching descendants (too narrow); the bigger risk is that a
// future change makes it return true when subjects are set or when the filter
// admits other scopes, so these cases are pinned explicitly.
func TestOwnerrefFilter_skipExpansion(t *testing.T) {
	tests := []struct {
		name       string
		filter     OwnerrefFilter
		entryScope aclscope.Scope
		want       bool
	}{
		{
			name:       "empty filter never skips (would need every descendant)",
			filter:     OwnerrefFilter{},
			entryScope: "KubernetesCluster",
			want:       false,
		},
		{
			name:       "single scope equal to entry, no subjects → skip",
			filter:     OwnerrefFilter{Scopes: []aclscope.Scope{"KubernetesCluster"}},
			entryScope: "KubernetesCluster",
			want:       true,
		},
		{
			name:       "single scope different from entry → must expand",
			filter:     OwnerrefFilter{Scopes: []aclscope.Scope{"KubernetesCluster"}},
			entryScope: "Project",
			want:       false,
		},
		{
			name:       "multiple scopes incl. one other than entry → must expand",
			filter:     OwnerrefFilter{Scopes: []aclscope.Scope{"KubernetesCluster", "Project"}},
			entryScope: "KubernetesCluster",
			want:       false,
		},
		{
			name:       "scope matches entry but subjects set → must expand",
			filter:     OwnerrefFilter{Scopes: []aclscope.Scope{"KubernetesCluster"}, Subjects: []aclscope.Subject{"cluster-1"}},
			entryScope: "KubernetesCluster",
			want:       false,
		},
		{
			name:       "subjects only, no scopes → must expand",
			filter:     OwnerrefFilter{Subjects: []aclscope.Subject{"cluster-1"}},
			entryScope: "KubernetesCluster",
			want:       false,
		},
		{
			name:       "duplicate scope equal to entry → skip",
			filter:     OwnerrefFilter{Scopes: []aclscope.Scope{"Workspace", "Workspace"}},
			entryScope: "Workspace",
			want:       true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, tt.filter.skipExpansion(tt.entryScope))
		})
	}
}
