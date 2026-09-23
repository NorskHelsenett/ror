package authtools_test

import (
	"testing"

	"github.com/NorskHelsenett/ror/pkg/auth/authtools"

	"github.com/stretchr/testify/assert"
)

func TestQualifyGroups(t *testing.T) {
	tests := []struct {
		name   string
		groups []string
		domain string
		want   []string
	}{
		{
			name:   "appends the domain",
			groups: []string{"admins", "readers"},
			domain: "example.com",
			want:   []string{"admins@example.com", "readers@example.com"},
		},
		{
			// A name that already carries a domain must not be qualified twice.
			name:   "leaves already qualified names",
			groups: []string{"admins@example.com", "readers"},
			domain: "example.com",
			want:   []string{"admins@example.com", "readers@example.com"},
		},
		{
			// Without a domain the previous implementation produced "group@".
			name:   "leaves names when domain is unknown",
			groups: []string{"admins"},
			domain: "",
			want:   []string{"admins"},
		},
		{
			name:   "empty input yields empty output",
			groups: nil,
			domain: "example.com",
			want:   []string{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.want, authtools.QualifyGroups(tt.groups, tt.domain))
		})
	}
}

func TestQualifyGroups_DoesNotMutateInput(t *testing.T) {
	groups := []string{"admins"}
	_ = authtools.QualifyGroups(groups, "example.com")
	assert.Equal(t, []string{"admins"}, groups)
}
