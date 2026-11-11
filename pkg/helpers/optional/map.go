package optional

// Map patches a map one key at a time.
//
// A plain Optional[map[K]V] replaces the whole map, which makes it impossible
// to drop a single key without resending every other key. Map carries an
// Optional per key instead, so each key gets the same three states a field
// gets:
//
//	key absent      - leave the key in the target as it is
//	key set to null - remove the key from the target
//	key with value  - write the key to the target
//
// Map is meant to be wrapped in an Optional, which adds the states of the map
// as a whole:
//
//	Labels optional.Optional[optional.Map[string, string]] `json:"labels,omitzero"`
//
//	unset - leave the map untouched
//	null  - clear the whole map
//	value - merge the keys of the patch into the map
//
// This mirrors how a json merge patch treats an object, and it is why maps are
// patched per key while slices are replaced as a whole: a slice has no stable
// key to address an element by.
type Map[K comparable, V any] map[K]Optional[V]

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

// Delete marks key to be removed from the target.
func (m Map[K, V]) Delete(key K) Map[K, V] {
	m[key] = Null[V]()
	return m
}

// ApplyTo merges the patch into target. Keys the patch does not mention are
// left alone, keys set to null are removed and keys holding a value are
// written. The target map is allocated when it is nil and the patch has a key
// to write.
//
// ApplyTo implements Patcher, so a Map can be applied with ApplyNested.
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
