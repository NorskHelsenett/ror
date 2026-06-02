package clusterinterregator

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewClusterInterregator_NilClient_DoesNotPanic(t *testing.T) {
	assert.NotPanics(t, func() {
		_ = NewClusterInterregator(nil)
	})
}
