package aclmodels

import (
	"testing"
	"time"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"

	"github.com/stretchr/testify/assert"
)

func qItem(id, group string, scope aclscope.Scope, subject aclscope.Subject) AclV3ListItem {
	return AclV3ListItem{Id: id, Version: 3, Group: group, Scope: scope, Subject: subject}
}

func TestAclV3ListItem_FieldValue(t *testing.T) {
	created := time.Date(2026, 1, 2, 3, 4, 5, 0, time.UTC)
	e := AclV3ListItem{Id: "id1", Group: "g", Scope: aclscope.ScopeCluster, Subject: "s", IssuedBy: "u@e.com", Created: created}

	assert.Equal(t, "g", e.FieldValue("group"))
	assert.Equal(t, "KubernetesCluster", e.FieldValue("scope"))
	assert.Equal(t, "s", e.FieldValue("subject"))
	assert.Equal(t, "u@e.com", e.FieldValue("issuedBy"))
	assert.Equal(t, "id1", e.FieldValue("id"))
	assert.Equal(t, "id1", e.FieldValue("_id"))
	assert.Equal(t, created.Format(time.RFC3339Nano), e.FieldValue("created"))
	assert.Equal(t, "", e.FieldValue("nope"))
}

// Sorted must be deterministic: equal primary keys break ties on id, regardless
// of input order, and the receiver is never mutated.
func TestAclV3List_Sorted(t *testing.T) {
	in := AclV3List{
		qItem("b", "team-a", aclscope.ScopeCluster, "c"),
		qItem("a", "team-a", aclscope.ScopeCluster, "c"),
		qItem("c", "team-b", aclscope.ScopeCluster, "c"),
	}

	asc := in.Sorted("group", true)
	assert.Equal(t, []string{"a", "b", "c"}, ids(asc), "same group ties broken by id ascending")

	desc := in.Sorted("group", false)
	assert.Equal(t, []string{"c", "a", "b"}, ids(desc), "primary desc, id tiebreaker still ascending")

	// Receiver unchanged.
	assert.Equal(t, []string{"b", "a", "c"}, ids(in), "Sorted must not mutate the receiver")
}

func TestAclV3List_Filter(t *testing.T) {
	in := AclV3List{
		qItem("1", "team-a", aclscope.ScopeCluster, "c"),
		qItem("2", "team-b", aclscope.ScopeProject, "p"),
		qItem("3", "team-a", aclscope.ScopeCluster, "c"),
	}
	got := in.Filter(func(e AclV3ListItem) bool { return e.Group == "team-a" })
	assert.Equal(t, []string{"1", "3"}, ids(got))
	assert.Len(t, in, 3, "Filter must not mutate the receiver")
}

func TestAclV3List_Page(t *testing.T) {
	in := AclV3List{qItem("1", "a", aclscope.ScopeCluster, "c"), qItem("2", "a", aclscope.ScopeCluster, "c"), qItem("3", "a", aclscope.ScopeCluster, "c")}

	assert.Equal(t, []string{"1", "2"}, ids(in.Page(0, 2)), "first page")
	assert.Equal(t, []string{"3"}, ids(in.Page(2, 2)), "partial last page")
	assert.Empty(t, ids(in.Page(5, 2)), "offset past end → empty")
	assert.Equal(t, []string{"1", "2", "3"}, ids(in.Page(0, 0)), "limit <= 0 → all from offset")
	assert.Equal(t, []string{"2", "3"}, ids(in.Page(1, -1)), "negative limit → all from offset")
	assert.Equal(t, []string{"1", "2"}, ids(in.Page(-5, 2)), "negative offset clamps to 0")
}

func ids(l AclV3List) []string {
	out := make([]string, len(l))
	for i, e := range l {
		out[i] = e.Id
	}
	return out
}
