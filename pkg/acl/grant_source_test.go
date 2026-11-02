package acl

import (
	"context"
	"errors"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclgrants"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

type grantSourceTestGroupReader struct {
	entries aclmodels.AclV3List
	groups  []string
	err     error
}

func (r *grantSourceTestGroupReader) GetByGroups(_ context.Context, groups []string) (aclmodels.AclV3List, error) {
	r.groups = append([]string(nil), groups...)
	return r.entries, r.err
}

func TestStoreGrantSource_MapsEntriesToRawGrants(t *testing.T) {
	reader := &grantSourceTestGroupReader{
		entries: aclmodels.AclV3List{
			{
				Id:      "entry-1",
				Group:   "team-a",
				Scope:   aclscope.ScopeProject,
				Subject: "project-1",
				Access:  []aclmodels.AccessTypeV3{"ror:read", "ror:write"},
			},
			{
				Id:      "entry-2",
				Group:   "team-b",
				Scope:   aclscope.ScopeRor,
				Subject: aclscope.SubjectGlobal,
				Access:  []aclmodels.AccessTypeV3{"ror:owner"},
			},
		},
	}
	source := storeGrantSource{reader: reader}

	grants, err := source.GetGrants(context.Background(), []string{"team-a", "team-b"})
	if err != nil {
		t.Fatalf("GetGrants returned error: %v", err)
	}

	want := aclgrants.Grants{
		{
			Scope:   aclscope.ScopeProject,
			Subject: "project-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read", "ror:write"},
		},
		{
			Scope:   aclscope.ScopeRor,
			Subject: aclscope.SubjectGlobal,
			Access:  []aclmodels.AccessTypeV3{"ror:owner"},
		},
	}
	if len(grants) != len(want) {
		t.Fatalf("got %d grants, want %d", len(grants), len(want))
	}
	for i := range want {
		if grants[i].Scope != want[i].Scope {
			t.Errorf("grant[%d].Scope = %q, want %q", i, grants[i].Scope, want[i].Scope)
		}
		if grants[i].Subject != want[i].Subject {
			t.Errorf("grant[%d].Subject = %q, want %q", i, grants[i].Subject, want[i].Subject)
		}
		if len(grants[i].Access) != len(want[i].Access) {
			t.Fatalf("grant[%d] access count = %d, want %d", i, len(grants[i].Access), len(want[i].Access))
		}
		for j := range want[i].Access {
			if grants[i].Access[j] != want[i].Access[j] {
				t.Errorf("grant[%d].Access[%d] = %q, want %q", i, j, grants[i].Access[j], want[i].Access[j])
			}
		}
	}
	if len(reader.groups) != 2 || reader.groups[0] != "team-a" || reader.groups[1] != "team-b" {
		t.Errorf("reader got groups %v, want [team-a team-b]", reader.groups)
	}
}

func TestStoreGrantSource_EmptyEntries(t *testing.T) {
	source := storeGrantSource{reader: &grantSourceTestGroupReader{}}

	grants, err := source.GetGrants(context.Background(), nil)
	if err != nil {
		t.Fatalf("GetGrants returned error: %v", err)
	}
	if len(grants) != 0 {
		t.Fatalf("got %d grants, want none", len(grants))
	}
}

func TestStoreGrantSource_PropagatesError(t *testing.T) {
	sourceErr := errors.New("store unavailable")
	source := storeGrantSource{reader: &grantSourceTestGroupReader{err: sourceErr}}

	_, err := source.GetGrants(context.Background(), []string{"team-a"})
	if !errors.Is(err, sourceErr) {
		t.Fatalf("GetGrants error = %v, want %v", err, sourceErr)
	}
}
