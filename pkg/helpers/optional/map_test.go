package optional_test

import (
	"encoding/json"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/helpers/optional"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMapAppliesPerKey(t *testing.T) {
	target := map[string]string{"keep": "me", "drop": "me", "change": "old"}

	patch := optional.Map[string, string]{}
	patch.Set("change", "new")
	patch.Set("add", "value")
	patch.Delete("drop")

	require.NoError(t, patch.ApplyTo(&target))

	assert.Equal(t, map[string]string{"keep": "me", "change": "new", "add": "value"}, target)
}

func TestMapAllocatesANilTarget(t *testing.T) {
	var target map[string]string
	require.NoError(t, optional.NewMap(map[string]string{"a": "b"}).ApplyTo(&target))
	assert.Equal(t, map[string]string{"a": "b"}, target)

	var untouched map[string]string
	require.NoError(t, optional.Map[string, string]{}.Delete("a").ApplyTo(&untouched))
	assert.Nil(t, untouched, "deleting from a nil map must not allocate one")
}

func TestMapWrappedInAnOptionalCarriesTheThreeStates(t *testing.T) {
	type doc struct {
		Labels optional.Optional[optional.Map[string, string]] `json:"labels,omitzero"`
	}

	original := map[string]string{"a": "1", "b": "2"}

	target := maps(original)
	var untouched doc
	require.NoError(t, json.Unmarshal([]byte(`{}`), &untouched))
	require.NoError(t, optional.ApplyNested(untouched.Labels, &target))
	assert.Equal(t, original, target, "an unset map is left untouched")

	target = maps(original)
	var cleared doc
	require.NoError(t, json.Unmarshal([]byte(`{"labels":null}`), &cleared))
	require.NoError(t, optional.ApplyNested(cleared.Labels, &target))
	assert.Empty(t, target, "a null map is cleared")

	target = maps(original)
	var merged doc
	require.NoError(t, json.Unmarshal([]byte(`{"labels":{"a":null,"c":"3"}}`), &merged))
	require.NoError(t, optional.ApplyNested(merged.Labels, &target))
	assert.Equal(t, map[string]string{"b": "2", "c": "3"}, target, "a map value is merged per key")
}

func maps(in map[string]string) map[string]string {
	out := make(map[string]string, len(in))
	for key, value := range in {
		out[key] = value
	}
	return out
}
