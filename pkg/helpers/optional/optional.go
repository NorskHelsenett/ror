package optional

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// Optional carries the three states a field in a partial update can have:
//
//	unset  - the field was not present in the payload, leave the target untouched
//	null   - the field was explicitly set to null, clear the target
//	value  - the field holds a value, write it to the target
//
// Optional is a value type, so a struct field holding an Optional must be
// tagged with `json:",omitzero"` to keep unset fields out of the marshalled
// document. The `omitempty` tag has no effect on struct types and would emit
// unset fields as null, which in merge patch semantics means "delete".
type Optional[T any] struct {
	set   bool
	null  bool
	value T
}

// NewOptional returns an Optional holding value.
func NewOptional[T any](value T) Optional[T] {
	return Optional[T]{set: true, null: false, value: value}
}

// Null returns an Optional that is explicitly set to null.
func Null[T any]() Optional[T] {
	return Optional[T]{set: true, null: true}
}

// Unset returns an Optional that is not set at all.
func Unset[T any]() Optional[T] {
	return Optional[T]{}
}

// FromPtr converts a pointer to an Optional. A nil pointer becomes null, any
// other pointer becomes a value. FromPtr can never return an unset Optional
// since a pointer only carries two states.
func FromPtr[T any](value *T) Optional[T] {
	if value == nil {
		return Null[T]()
	}
	return NewOptional(*value)
}

// Set sets the Optional to value.
func (o *Optional[T]) Set(value T) {
	o.set = true
	o.null = false
	o.value = value
}

// SetNull sets the Optional to null.
func (o *Optional[T]) SetNull() {
	var zero T
	o.set = true
	o.null = true
	o.value = zero
}

// Unset removes any state from the Optional.
func (o *Optional[T]) Unset() {
	var zero T
	o.set = false
	o.null = false
	o.value = zero
}

func (o *Optional[T]) UnmarshalJSON(b []byte) error {
	o.set = true

	// explicit null
	if bytes.Equal(bytes.TrimSpace(b), []byte("null")) {
		o.null = true
		var zero T
		o.value = zero
		return nil
	}

	o.null = false
	return json.Unmarshal(b, &o.value)
}

// MarshalJSON renders the Optional as its value, or as null when it is null or
// unset. Unset Optionals only marshal correctly inside a struct field tagged
// with `json:",omitzero"`, which drops the field before this method is called.
func (o Optional[T]) MarshalJSON() ([]byte, error) {
	if !o.set {
		return []byte("null"), nil
	}
	if o.null {
		return []byte("null"), nil
	}
	return json.Marshal(o.value)
}

// IsZero reports whether the Optional is unset. It implements the interface
// used by encoding/json for the `omitzero` tag option, so that unset fields are
// omitted from the marshalled document instead of being emitted as null.
func (o Optional[T]) IsZero() bool { return !o.set }

func (o Optional[T]) IsSet() bool  { return o.set }
func (o Optional[T]) IsNull() bool { return o.set && o.null }
func (o Optional[T]) Value() (T, bool) {
	return o.value, o.set && !o.null
}

// ValueOr returns the held value, or fallback when the Optional is unset or null.
func (o Optional[T]) ValueOr(fallback T) T {
	if v, ok := o.Value(); ok {
		return v
	}
	return fallback
}

// Ptr returns a pointer to the held value, or nil when the Optional is unset or
// null. The returned pointer refers to a copy of the value.
func (o Optional[T]) Ptr() *T {
	v, ok := o.Value()
	if !ok {
		return nil
	}
	return &v
}

// Any returns the value as interface{} to satisfy OptionalAny interface
func (o Optional[T]) Any() (any, bool) {
	if !o.set || o.null {
		return nil, false
	}
	return o.value, true
}

func (o Optional[T]) String() string {
	switch {
	case !o.set:
		return "unset"
	case o.null:
		return "null"
	default:
		return fmt.Sprintf("value(%v)", o.value)
	}
}

func SetUpdate[T any](m map[string]any, column string, o Optional[T]) {
	if !o.IsSet() {
		return
	}
	if o.IsNull() {
		m[column] = nil
		return
	}
	if v, ok := o.Value(); ok {
		m[column] = v
	}
}

// Apply writes o into target. An unset Optional leaves target untouched, a null
// Optional resets target to its zero value and a value is assigned to target.
//
// Apply is the leaf of a patch: use it for fields that are replaced as a whole,
// such as scalars, slices, maps and structs that are not patchable themselves.
func Apply[T any](o Optional[T], target *T) {
	if target == nil || !o.set {
		return
	}
	if o.null {
		var zero T
		*target = zero
		return
	}
	*target = o.value
}

// Patcher is implemented by a patch type that knows how to apply itself to the
// type it mirrors.
type Patcher[T any] interface {
	ApplyTo(target *T) error
}

// ApplyNested writes a nested patch into target. An unset Optional leaves
// target untouched, a null Optional resets target to its zero value and a value
// is applied field by field by the nested patch.
//
// Use ApplyNested for fields that have a patch mirror of their own, so that a
// patch touching one field does not overwrite its siblings.
func ApplyNested[P Patcher[T], T any](o Optional[P], target *T) error {
	if target == nil || !o.set {
		return nil
	}
	if o.null {
		var zero T
		*target = zero
		return nil
	}
	return o.value.ApplyTo(target)
}
