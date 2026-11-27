package optional

// Map carries an Optional per key instead, so each key gets the same three states
// a field gets:
//
//	key absent      - nothing happens
//	key set to null - remove the key
//	key with value  - write to the key
//
// Map is meant to be wrapped in an Optional, which adds the states of the map
// as a whole:
//
//	Labels optional.Optional[optional.Map[string, string]] `json:"labels,omitzero"`
//
//	unset - leave the map untouched
//	null  - clear the whole map
//	value - merge the keys of the patch into the map
type Map[K comparable, V any] map[K]Optional[V]

// Interface guard clause for Patcher interface
var _ Patcher[map[string]string] = new(Map[string, string])

// NewMap builds a Map that sets every key of values, and removes nothing.
func NewMap[K comparable, V any](values map[K]V) Map[K, V] {
	patch := make(Map[K, V], len(values))
	for key, value := range values {
		patch[key] = NewOptional(value)
	}
	return patch
}

// Set marks key to be written with value.
func (m Map[K, V]) Set(key K, value V) Map[K, V] {
	m[key] = NewOptional(value)
	return m
}

// Get returns the value and a boolean indicating if the key is set.
// If the bool is false, the value is the zero value of V - "" for string, 0 for int, etc..
// If the bool is true, the  value is the value of V.
func (m Map[K, V]) Get(key K) (V, bool) {
	return m[key].value, m[key].IsSet()
}

// Delete marks key to be removed from the target.
func (m Map[K, V]) Delete(key K) Map[K, V] {
	m[key] = Null[V]()
	return m
}

// ApplyTo applies the map to the target map, writing values and removing keys. Unset keys are ignored.
func (m Map[K, V]) ApplyTo(target *map[K]V) error {
	if target == nil || len(m) == 0 {
		return nil
	}
	for key, value := range m {
		if !value.IsSet() {
			// An unset entry carries no intent, treat it like an absent key.
			continue
		}
		if value.IsNull() {
			delete(*target, key)
			continue
		}
		if *target == nil {
			*target = make(map[K]V, len(m))
		}
		(*target)[key], _ = value.Value()
	}
	return nil
}
