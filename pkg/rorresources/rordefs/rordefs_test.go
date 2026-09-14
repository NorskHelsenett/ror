package rordefs_test

import (
	"testing"

	"github.com/NorskHelsenett/ror/pkg/models/aclmodels/aclcaps"
	"github.com/NorskHelsenett/ror/pkg/rorresources/rordefs"
	"github.com/stretchr/testify/assert"
)

func TestProtectedByKind(t *testing.T) {
	assert.Equal(t, aclcaps.CapRorConfig, rordefs.Resourcedefs.ProtectedByKind("Config"))
	assert.Equal(t, aclcaps.Capability(""), rordefs.Resourcedefs.ProtectedByKind("Pod"))
	assert.Equal(t, aclcaps.Capability(""), rordefs.Resourcedefs.ProtectedByKind("DoesNotExist"))
}
