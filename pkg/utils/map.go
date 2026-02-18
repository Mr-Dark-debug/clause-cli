package utils

// Keys returns all keys from a map.
func Keys[M ~map[K]V, K comparable, V any](m M) []K {
	if m == nil {
		return nil
	}
	keys := make([]K, 0, len(m))
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}

// Values returns all values from a map.
func Values[M ~map[K]V, K comparable, V any](m M) []V {
	if m == nil {
		return nil
	}
	values := make([]V, 0, len(m))
	for _, v := range m {
		values = append(values, v)
	}
	return values
}

// Entries returns all key-value pairs from a map.
func Entries[M ~map[K]V, K comparable, V any](m M) []Entry[K, V] {
	if m == nil {
		return nil
	}
	entries := make([]Entry[K, V], 0, len(m))
	for k, v := range m {
		entries = append(entries, Entry[K, V]{Key: k, Value: v})
	}
	return entries
}

// Entry represents a key-value pair.
type Entry[K comparable, V any] struct {
	Key   K
	Value V
}

// FromEntries creates a map from key-value pairs.
func FromEntries[K comparable, V any](entries []Entry[K, V]) map[K]V {
	if entries == nil {
		return nil
	}
	m := make(map[K]V, len(entries))
	for _, e := range entries {
		m[e.Key] = e.Value
	}
	return m
}

// FromPairs creates a map from key-value pairs as variadic arguments.
func FromPairs[K comparable, V any](pairs ...Pair[K, V]) map[K]V {
	if pairs == nil {
		return nil
	}
	m := make(map[K]V, len(pairs))
	for _, p := range pairs {
		m[p.Key] = p.Value
	}
	return m
}

// Pair is a convenience type for Entry.
type Pair[K comparable, V any] struct {
	Key   K
	Value V
}

// GetOrDefault retrieves a value from a map, or returns a default if the key doesn't exist.
func GetOrDefault[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if m == nil {
		return defaultValue
	}
	if v, ok := m[key]; ok {
		return v
	}
	return defaultValue
}

// GetOr retrieves a value from a map, or computes and stores a default value.
func GetOr[K comparable, V any](m map[K]V, key K, defaultValueFn func() V) V {
	if m == nil {
		return defaultValueFn()
	}
	if v, ok := m[key]; ok {
		return v
	}
	return defaultValueFn()
}

// GetOrInsert retrieves a value or inserts and returns the default.
func GetOrInsert[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if m == nil {
		return defaultValue
	}
	if v, ok := m[key]; ok {
		return v
	}
	m[key] = defaultValue
	return defaultValue
}

// GetOrInsertComputed retrieves a value or inserts and returns a computed default.
func GetOrInsertComputed[K comparable, V any](m map[K]V, key K, compute func() V) V {
	if m == nil {
		return compute()
	}
	if v, ok := m[key]; ok {
		return v
	}
	v := compute()
	m[key] = v
	return v
}

// HasKey checks if a map contains a key.
func HasKey[K comparable, V any](m map[K]V, key K) bool {
	if m == nil {
		return false
	}
	_, ok := m[key]
	return ok
}

// HasValue checks if a map contains a value.
func HasValue[M ~map[K]V, K comparable, V comparable](m M, value V) bool {
	if m == nil {
		return false
	}
	for _, v := range m {
		if v == value {
			return true
		}
	}
	return false
}

// Merge combines multiple maps into a new map.
// Later maps overwrite earlier ones for duplicate keys.
func Merge[K comparable, V any](maps ...map[K]V) map[K]V {
	if len(maps) == 0 {
		return nil
	}

	totalLen := 0
	for _, m := range maps {
		if m != nil {
			totalLen += len(m)
		}
	}

	result := make(map[K]V, totalLen)
	for _, m := range maps {
		if m != nil {
			for k, v := range m {
				result[k] = v
			}
		}
	}
	return result
}

// MergeInto merges all source maps into the destination map.
func MergeInto[K comparable, V any](dest map[K]V, sources ...map[K]V) {
	if dest == nil {
		return
	}
	for _, src := range sources {
		if src != nil {
			for k, v := range src {
				dest[k] = v
			}
		}
	}
}

// CloneMap creates a shallow copy of a map.
func CloneMap[M ~map[K]V, K comparable, V any](m M) M {
	if m == nil {
		return nil
	}
	result := make(M, len(m))
	for k, v := range m {
		result[k] = v
	}
	return result
}

// Copy copies all entries from src to dst.
func Copy[K comparable, V any](dst, src map[K]V) {
	if dst == nil || src == nil {
		return
	}
	for k, v := range src {
		dst[k] = v
	}
}

// Delete removes keys from a map.
func Delete[K comparable, V any](m map[K]V, keys ...K) {
	if m == nil {
		return
	}
	for _, k := range keys {
		delete(m, k)
	}
}

// DeleteIf removes entries that match a predicate.
func DeleteIf[K comparable, V any](m map[K]V, predicate func(K, V) bool) {
	if m == nil {
		return
	}
	for k, v := range m {
		if predicate(k, v) {
			delete(m, k)
		}
	}
}

// Clear removes all entries from a map.
func Clear[K comparable, V any](m map[K]V) {
	if m == nil {
		return
	}
	for k := range m {
		delete(m, k)
	}
}

// IsEmptyMap checks if a map is empty.
func IsEmptyMap[K comparable, V any](m map[K]V) bool {
	return m == nil || len(m) == 0
}

// IsNotEmptyMap checks if a map is not empty.
func IsNotEmptyMap[K comparable, V any](m map[K]V) bool {
	return m != nil && len(m) > 0
}

// LenMap returns the length of a map.
func LenMap[K comparable, V any](m map[K]V) int {
	if m == nil {
		return 0
	}
	return len(m)
}

// ForEachMap iterates over all entries in a map.
func ForEachMap[K comparable, V any](m map[K]V, fn func(K, V)) {
	if m == nil {
		return
	}
	for k, v := range m {
		fn(k, v)
	}
}

// ForEachIf iterates over entries while the predicate returns true.
func ForEachIf[K comparable, V any](m map[K]V, fn func(K, V) bool) {
	if m == nil {
		return
	}
	for k, v := range m {
		if !fn(k, v) {
			break
		}
	}
}

// MapKeys transforms all keys in a map.
func MapKeys[K1, K2 comparable, V any](m map[K1]V, fn func(K1) K2) map[K2]V {
	if m == nil {
		return nil
	}
	result := make(map[K2]V, len(m))
	for k, v := range m {
		result[fn(k)] = v
	}
	return result
}

// MapValues transforms all values in a map.
func MapValues[K comparable, V1, V2 any](m map[K]V1, fn func(V1) V2) map[K]V2 {
	if m == nil {
		return nil
	}
	result := make(map[K]V2, len(m))
	for k, v := range m {
		result[k] = fn(v)
	}
	return result
}

// MapEntries transforms all entries in a map.
func MapEntries[K1, K2 comparable, V1, V2 any](m map[K1]V1, fn func(K1, V1) (K2, V2)) map[K2]V2 {
	if m == nil {
		return nil
	}
	result := make(map[K2]V2, len(m))
	for k, v := range m {
		k2, v2 := fn(k, v)
		result[k2] = v2
	}
	return result
}

// FilterMap returns a new map with entries that pass the predicate.
func FilterMap[K comparable, V any](m map[K]V, predicate func(K, V) bool) map[K]V {
	if m == nil {
		return nil
	}
	result := make(map[K]V)
	for k, v := range m {
		if predicate(k, v) {
			result[k] = v
		}
	}
	return result
}

// FilterKeys returns a new map with only the specified keys.
func FilterKeys[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	if m == nil || len(keys) == 0 {
		return nil
	}
	result := make(map[K]V, len(keys))
	for _, k := range keys {
		if v, ok := m[k]; ok {
			result[k] = v
		}
	}
	return result
}

// Omit returns a new map without the specified keys.
func Omit[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	if m == nil {
		return nil
	}
	omitSet := make(map[K]bool, len(keys))
	for _, k := range keys {
		omitSet[k] = true
	}

	result := make(map[K]V, len(m)-len(keys))
	for k, v := range m {
		if !omitSet[k] {
			result[k] = v
		}
	}
	return result
}

// Pick returns a new map with only the specified keys (alias for FilterKeys).
func Pick[K comparable, V any](m map[K]V, keys ...K) map[K]V {
	return FilterKeys(m, keys...)
}

// Invert creates a new map with keys and values swapped.
func Invert[M ~map[K]V, K, V comparable](m M) map[V]K {
	if m == nil {
		return nil
	}
	result := make(map[V]K, len(m))
	for k, v := range m {
		result[v] = k
	}
	return result
}

// GroupBy groups values by a key function.
func GroupBy[T any, K comparable](slice []T, keyFn func(T) K) map[K][]T {
	if slice == nil {
		return nil
	}
	result := make(map[K][]T)
	for _, v := range slice {
		key := keyFn(v)
		result[key] = append(result[key], v)
	}
	return result
}

// GroupByMap groups values and transforms them by a key function.
func GroupByMap[T any, K comparable, V any](slice []T, fn func(T) (K, V)) map[K][]V {
	if slice == nil {
		return nil
	}
	result := make(map[K][]V)
	for _, item := range slice {
		k, v := fn(item)
		result[k] = append(result[k], v)
	}
	return result
}

// CountBy counts occurrences by a key function.
func CountBy[T any, K comparable](slice []T, keyFn func(T) K) map[K]int {
	if slice == nil {
		return nil
	}
	result := make(map[K]int)
	for _, v := range slice {
		key := keyFn(v)
		result[key]++
	}
	return result
}

// ToSlice converts a map to a slice using a transform function.
func ToSlice[K comparable, V any, T any](m map[K]V, fn func(K, V) T) []T {
	if m == nil {
		return nil
	}
	result := make([]T, 0, len(m))
	for k, v := range m {
		result = append(result, fn(k, v))
	}
	return result
}

// ToSet creates a set (map with bool values) from a slice.
func ToSet[T comparable](slice []T) map[T]bool {
	if slice == nil {
		return nil
	}
	result := make(map[T]bool, len(slice))
	for _, v := range slice {
		result[v] = true
	}
	return result
}

// SetFromKeys creates a set from map keys.
func SetFromKeys[K comparable, V any](m map[K]V) map[K]bool {
	if m == nil {
		return nil
	}
	result := make(map[K]bool, len(m))
	for k := range m {
		result[k] = true
	}
	return result
}

// EqualMaps checks if two maps are equal.
func EqualMaps[M1 ~map[K]V1, M2 ~map[K]V2, K comparable, V1, V2 comparable](m1 M1, m2 M2, valueEq func(V1, V2) bool) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if m1 == nil || m2 == nil {
		return false
	}
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok {
			return false
		}
		if valueEq != nil {
			if !valueEq(v1, v2) {
				return false
			}
		} else {
			var v1Typed V1 = any(v2).(V1)
			if v1 != v1Typed {
				return false
			}
		}
	}
	return true
}

// EqualSimple checks if two maps with the same value type are equal.
func EqualSimple[M ~map[K]V, K, V comparable](m1, m2 M) bool {
	if m1 == nil && m2 == nil {
		return true
	}
	if m1 == nil || m2 == nil {
		return false
	}
	if len(m1) != len(m2) {
		return false
	}
	for k, v1 := range m1 {
		v2, ok := m2[k]
		if !ok || v1 != v2 {
			return false
		}
	}
	return true
}

// FindKey finds a key by its value.
func FindKey[K comparable, V comparable](m map[K]V, value V) (K, bool) {
	if m == nil {
		var zero K
		return zero, false
	}
	for k, v := range m {
		if v == value {
			return k, true
		}
	}
	var zero K
	return zero, false
}

// FindKeyBy finds a key by a predicate on its value.
func FindKeyBy[K comparable, V any](m map[K]V, predicate func(V) bool) (K, bool) {
	if m == nil {
		var zero K
		return zero, false
	}
	for k, v := range m {
		if predicate(v) {
			return k, true
		}
	}
	var zero K
	return zero, false
}

// Update updates a value in a map using a function.
func Update[K comparable, V any](m map[K]V, key K, fn func(V) V) {
	if m == nil {
		return
	}
	if v, ok := m[key]; ok {
		m[key] = fn(v)
	}
}

// UpdateOrInsert updates a value or inserts a default.
func UpdateOrInsert[K comparable, V any](m map[K]V, key K, fn func(V) V, defaultValue V) {
	if m == nil {
		return
	}
	if v, ok := m[key]; ok {
		m[key] = fn(v)
	} else {
		m[key] = defaultValue
	}
}

// UpdateWithDefault updates a value or inserts a computed default.
func UpdateWithDefault[K comparable, V any](m map[K]V, key K, fn func(V) V, defaultFn func() V) {
	if m == nil {
		return
	}
	if v, ok := m[key]; ok {
		m[key] = fn(v)
	} else {
		m[key] = defaultFn()
	}
}
