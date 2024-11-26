package slice2

import "math/rand"

// Contains if the element is in the ss
func Contains[T comparable](ss []T, element T) bool {
	for _, item := range ss {
		if item == element {
			return true
		}
	}

	return false
}

// PruneEmptyString removes empty strings
func PruneEmptyString(v []string) []string {
	return PruneEqual(v, "")
}

// PruneEqual removes all elements from the input slice that are equal to the specified value.
func PruneEqual[T comparable](ss []T, equalTo T) (r []T) {
	for i := range ss {
		if ss[i] != equalTo {
			r = append(r, ss[i])
		}
	}

	return
}

// UniqStable returns a slice with all duplicate values removed in the order they first appeared
func UniqStable[T comparable](ss []T) []T {
	result := make([]T, 0, len(ss))
	seen := make(map[T]struct{}, len(ss))

	for _, item := range ss {
		if _, ok := seen[item]; ok {
			continue
		}

		seen[item] = struct{}{}
		result = append(result, item)
	}

	return result
}

// Uniq returns a slice with all duplicate values removed
func Uniq[T comparable](ss []T) []T {
	return UniqStable(ss)
}

// Chunk Chunks an array into smaller arrays of a specified size
func Chunk[T any](ss []T, size int) [][]T {
	if size <= 0 {
		panic("Second parameter must be greater than 0")
	}

	chunksNum := len(ss) / size
	if len(ss)%size != 0 {
		chunksNum += 1
	}

	result := make([][]T, 0, chunksNum)

	for i := 0; i < chunksNum; i++ {
		last := (i + 1) * size
		if last > len(ss) {
			last = len(ss)
		}
		result = append(result, ss[i*size:last])
	}

	return result
}

// Count the occurrences of a value in a ss
func Count[T comparable](ss []T, value T) (count int) {
	for _, item := range ss {
		if item == value {
			count++
		}
	}

	return count
}

// Filter the array based on a predicate
func Filter[V any](ss []V, predicate func(item V, index int) bool) []V {
	result := make([]V, 0, len(ss))

	for i, item := range ss {
		if predicate(item, i) {
			result = append(result, item)
		}
	}

	return result
}

// Map produces a new array of values by mapping each value in list through a transformation function
func Map[T any, R any](ss []T, iteratee func(item T, index int) R) []R {
	result := make([]R, len(ss))

	for i, item := range ss {
		result[i] = iteratee(item, i)
	}

	return result
}

// First returns the first element of an array
func First[T any](ss []T) T {
	if len(ss) == 0 {
		var zero T
		return zero
	}

	return ss[0]
}

// FirstOr returns the first element of an array or a default value if the array is empty
func FirstOr[T any](ss []T, defaultValue T) T {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[0]
}

// Last returns the last element of an array
func Last[T any](ss []T) T {
	if len(ss) == 0 {
		var zero T
		return zero
	}

	return ss[len(ss)-1]
}

// LastOr returns the last element of an array or a default value if the array is empty
func LastOr[T any](ss []T, defaultValue T) T {
	if len(ss) == 0 {
		return defaultValue
	}

	return ss[len(ss)-1]
}

// IsEmpty checks if the array is empty
func IsEmpty[T comparable](ss []T) bool {
	return len(ss) == 0
}

func Equal[T comparable](s1, s2 []T) bool {
	if len(s1) != len(s2) {
		return false
	}

	for idx := range s1 {
		if s1[idx] != s2[idx] {
			return false
		}
	}

	return true
}

// Clone returns a shallow copy of a portion of an array into a new array object
func Clone[T comparable](ss []T) []T {
	result := make([]T, len(ss))
	copy(result, ss)
	return result
}

// PickRandom returns a random element from the array
func PickRandom[T any](ss []T) T {
	return ss[rand.Intn(len(ss))]
}

// Diff returns the difference between two arrays
func Diff[T comparable](s1 []T, s2 []T) (d1, d2 []T) {
	f := func(ss1, ss2 []T) (result []T) {
		set := make(map[T]struct{}, len(ss1))

		for _, s := range ss1 {
			set[s] = struct{}{}
		}

		for _, s := range ss2 {
			if _, ok := set[s]; ok {
				delete(set, s)
			} else {
				result = append(result, s)
			}
		}
		return
	}

	d1 = f(s2, s1)
	d2 = f(s1, s2)

	return
}

// Merge merges multiple arrays into a single array
func Merge[T comparable](ss ...[]T) []T {
	var result []T
	for _, s := range ss {
		result = append(result, s...)
	}
	return Uniq(result)
}
