package acl

import (
	"context"
	"errors"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclgrants"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

type testGroupReader struct {
	entries aclmodels.AclV3List
	err     error
}

func (r testGroupReader) GetByGroups(context.Context, []string) (aclmodels.AclV3List, error) {
	return r.entries, r.err
}

func TestStoreGrantSourceMapsEntries(t *testing.T) {
	source := NewStoreGrantSource(testGroupReader{
		entries: aclmodels.AclV3List{{
			Id:      "acl-entry-id",
			Group:   "team-a",
			Scope:   aclscope.ScopeProject,
			Subject: "project-1",
			Access:  []aclmodels.AccessTypeV3{"ror:read", "ror:write"},
		}},
	})

	grants, err := source.GetGrants(context.Background(), []string{"team-a"})
	if err != nil {
		t.Fatalf("GetGrants returned error: %v", err)
	}
	want := aclgrants.Grants{{
		Scope:   aclscope.ScopeProject,
		Subject: "project-1",
		Access:  []aclmodels.AccessTypeV3{"ror:read", "ror:write"},
	}}
	if len(grants) != len(want) || grants[0].Scope != want[0].Scope ||
		grants[0].Subject != want[0].Subject || len(grants[0].Access) != len(want[0].Access) {
		t.Fatalf("unexpected grants: %#v", grants)
	}
	for i := range want[0].Access {
		if grants[0].Access[i] != want[0].Access[i] {
			t.Errorf("access[%d] = %q, want %q", i, grants[0].Access[i], want[0].Access[i])
		}
	}
}

func TestStoreGrantSourcePropagatesError(t *testing.T) {
	wantErr := errors.New("store unavailable")
	source := NewStoreGrantSource(testGroupReader{err: wantErr})

	_, err := source.GetGrants(context.Background(), nil)
	if !errors.Is(err, wantErr) {
		t.Fatalf("GetGrants error = %v, want %v", err, wantErr)
	}
}
