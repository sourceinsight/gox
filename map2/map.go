package map2

import (
	"golang.org/x/exp/constraints"
	"sort"
)

// Keys returns all keys of the map.
func Keys[K comparable, V any](m map[K]V) []K {
	result := make([]K, 0, len(m))

	for k := range m {
		result = append(result, k)
	}

	return result
}

// KeysSorted returns all keys of the map in sorted order.
func KeysSorted[K constraints.Ordered, V any](m map[K]V) []K {
	keys := Keys(m)
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})
	return keys
}

// Values returns all values of the map.
func Values[K comparable, V any](m map[K]V) []V {
	result := make([]V, 0, len(m))

	for _, v := range m {
		result = append(result, v)
	}

	return result
}

// IsEmpty returns true if the map is empty.
func IsEmpty[K comparable, V any](m map[K]V) bool {
	return len(m) == 0
}

// HasKey returns true if the map contains the key.
func HasKey[K comparable, V any](m map[K]V, key K) bool {
	_, ok := m[key]
	return ok
}

// ValueOr returns the value for the key if it exists, otherwise returns the default value.
func ValueOr[K comparable, V any](m map[K]V, key K, defaultValue V) V {
	if v, ok := m[key]; ok {
		return v
	}
	return defaultValue
}

func ToSlice[K comparable, V any, R any](m map[K]V, iteratee func(key K, value V) R) []R {
	result := make([]R, 0, len(m))

	for k := range m {
		result = append(result, iteratee(k, m[k]))
	}

	return result
}
