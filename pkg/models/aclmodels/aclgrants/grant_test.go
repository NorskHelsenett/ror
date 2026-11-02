package aclgrants

import (
	"encoding/json"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclcaps"
	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclscope"
)

func TestGrantsJSONContract(t *testing.T) {
	grants := Grants{
		{
			Scope:   aclscope.ScopeProject,
			Subject: aclscope.Subject("proj-blue"),
			Access: []aclcaps.AccessTypeV3{
				"ror:read",
				"kubernetes:logon",
			},
		},
	}

	const want = `[{"scope":"Project","subject":"proj-blue","access":["ror:read","kubernetes:logon"]}]`
	encoded, err := json.Marshal(grants)
	if err != nil {
		t.Fatalf("marshal grants: %v", err)
	}
	if string(encoded) != want {
		t.Fatalf("unexpected grant JSON:\n got: %s\nwant: %s", encoded, want)
	}

	var decoded Grants
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		t.Fatalf("unmarshal grants: %v", err)
	}
	if len(decoded) != 1 {
		t.Fatalf("got %d grants, want 1", len(decoded))
	}
	if decoded[0].Scope != grants[0].Scope {
		t.Errorf("scope = %q, want %q", decoded[0].Scope, grants[0].Scope)
	}
	if decoded[0].Subject != grants[0].Subject {
		t.Errorf("subject = %q, want %q", decoded[0].Subject, grants[0].Subject)
	}
	if len(decoded[0].Access) != len(grants[0].Access) {
		t.Fatalf("got %d access values, want %d", len(decoded[0].Access), len(grants[0].Access))
	}
	for i := range grants[0].Access {
		if decoded[0].Access[i] != grants[0].Access[i] {
			t.Errorf("access[%d] = %q, want %q", i, decoded[0].Access[i], grants[0].Access[i])
		}
	}
}

func TestClaimKey(t *testing.T) {
	if ClaimKey != "ror_grants" {
		t.Fatalf("claim key = %q, want %q", ClaimKey, "ror_grants")
	}
}
