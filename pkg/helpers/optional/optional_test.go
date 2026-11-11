package optional_test

import (
	"encoding/json"
	"errors"
	"testing"

	"github.com/NorskHelsenett/ror/pkg/helpers/optional"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

type patchDoc struct {
	Name  optional.Optional[string] `json:"name,omitzero"`
	Count optional.Optional[int]    `json:"count,omitzero"`
}

func TestOptionalDecodesThreeStates(t *testing.T) {
	tests := []struct {
		name     string
		payload  string
		wantSet  bool
		wantNull bool
		wantVal  string
	}{
		{name: "absent", payload: `{}`, wantSet: false, wantNull: false},
		{name: "explicit null", payload: `{"name":null}`, wantSet: true, wantNull: true},
		{name: "value", payload: `{"name":"ror"}`, wantSet: true, wantNull: false, wantVal: "ror"},
		{name: "zero value", payload: `{"name":""}`, wantSet: true, wantNull: false, wantVal: ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var doc patchDoc
			require.NoError(t, json.Unmarshal([]byte(tt.payload), &doc))

			assert.Equal(t, tt.wantSet, doc.Name.IsSet())
			assert.Equal(t, tt.wantNull, doc.Name.IsNull())

			value, ok := doc.Name.Value()
			assert.Equal(t, tt.wantSet && !tt.wantNull, ok)
			assert.Equal(t, tt.wantVal, value)
		})
	}
}

// Unset optionals must disappear from the document. Without omitzero they would
// be emitted as null, which in merge patch semantics means "delete".
func TestOptionalOmitzeroDropsUnsetFields(t *testing.T) {
	doc := patchDoc{Name: optional.NewOptional("ror")}

	encoded, err := json.Marshal(doc)
	require.NoError(t, err)
	assert.JSONEq(t, `{"name":"ror"}`, string(encoded))

	doc.Count.SetNull()
	encoded, err = json.Marshal(doc)
	require.NoError(t, err)
	assert.JSONEq(t, `{"name":"ror","count":null}`, string(encoded))
}

func TestOptionalStateTransitions(t *testing.T) {
	var value optional.Optional[int]
	assert.True(t, value.IsZero())

	value.Set(0)
	assert.True(t, value.IsSet())
	assert.False(t, value.IsNull())
	assert.False(t, value.IsZero())
	assert.Equal(t, 7, optional.NewOptional(7).ValueOr(1))
	assert.Equal(t, 1, optional.Null[int]().ValueOr(1))
	assert.Equal(t, 1, optional.Unset[int]().ValueOr(1))

	value.SetNull()
	assert.True(t, value.IsNull())
	assert.Nil(t, value.Ptr())

	value.Unset()
	assert.True(t, value.IsZero())
	assert.False(t, value.IsSet())
}

func TestOptionalFromPtr(t *testing.T) {
	assert.True(t, optional.FromPtr[string](nil).IsNull())

	value := "ror"
	fromPtr := optional.FromPtr(&value)
	assert.Equal(t, "ror", fromPtr.ValueOr(""))
	require.NotNil(t, fromPtr.Ptr())

	// Ptr returns a copy, mutating it must not reach the optional.
	*fromPtr.Ptr() = "changed"
	assert.Equal(t, "ror", fromPtr.ValueOr(""))
}

type nested struct {
	Label string
	Score int
}

type target struct {
	Name   string
	Nested nested
}

type nestedPatch struct {
	Label optional.Optional[string]
}

func (p nestedPatch) ApplyTo(t *nested) error {
	if t == nil {
		return errNilTarget
	}
	optional.Apply(p.Label, &t.Label)
	return nil
}

var errNilTarget = errors.New("nil target")

func TestApplyHandlesTheThreeStates(t *testing.T) {
	value := target{Name: "original"}
	optional.Apply(optional.Unset[string](), &value.Name)
	assert.Equal(t, "original", value.Name, "unset leaves the target untouched")

	optional.Apply(optional.NewOptional("changed"), &value.Name)
	assert.Equal(t, "changed", value.Name, "a value is written to the target")

	optional.Apply(optional.Null[string](), &value.Name)
	assert.Equal(t, "", value.Name, "null resets the target")
}

func TestApplyNestedHandlesTheThreeStates(t *testing.T) {
	value := target{Nested: nested{Label: "original", Score: 7}}

	require.NoError(t, optional.ApplyNested(optional.Unset[nestedPatch](), &value.Nested))
	assert.Equal(t, nested{Label: "original", Score: 7}, value.Nested, "unset leaves the target untouched")

	require.NoError(t, optional.ApplyNested(optional.NewOptional(nestedPatch{Label: optional.NewOptional("changed")}), &value.Nested))
	assert.Equal(t, nested{Label: "changed", Score: 7}, value.Nested, "a nested patch only touches the fields it mentions")

	require.NoError(t, optional.ApplyNested(optional.Null[nestedPatch](), &value.Nested))
	assert.Equal(t, nested{}, value.Nested, "null resets the whole nested struct")
}

func TestApplyNestedPropagatesErrors(t *testing.T) {
	var missing *nested
	assert.NoError(t, optional.ApplyNested(optional.NewOptional(nestedPatch{}), missing), "a nil target is a no-op")

	value := nested{Label: "original"}
	assert.ErrorIs(t, optional.ApplyNested(optional.NewOptional(failingPatch{}), &value), errPatchFailed)
}

type failingPatch struct{}

func (failingPatch) ApplyTo(*nested) error { return errPatchFailed }

var errPatchFailed = errors.New("patch failed")
