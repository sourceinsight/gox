package set

// Set is a set of items.
type Set[T comparable] map[T]struct{}

// New creates a new set.
func New[T comparable](items ...T) Set[T] {
	ss := Set[T]{}
	_ = ss.Adds(items...)
	return ss
}

// Add adds an item to the set.
func (s Set[T]) Add(item T) bool {
	l := len(s)
	s[item] = struct{}{}
	return len(s) != l
}

// Adds adds items to the set.
func (s Set[T]) Adds(items ...T) int {
	l := len(s)
	for _, item := range items {
		s[item] = struct{}{}
	}
	return len(s) - l
}

// Delete deletes an item from the set.
func (s Set[T]) Delete(item T) bool {
	l := len(s)
	delete(s, item)
	return len(s) != l
}

// Deletes deletes items from the set.
func (s Set[T]) Deletes(items ...T) int {
	l := len(s)
	for _, item := range items {
		delete(s, item)
	}
	return l - len(s)
}

// Has returns true if the item is in the set.
func (s Set[T]) Has(item T) bool {
	_, ok := s[item]
	return ok
}

// HasAll returns true if all the items are in the set.
func (s Set[T]) HasAll(items ...T) bool {
	for _, item := range items {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// HasAny returns true if any of the items are in the set.
func (s Set[T]) HasAny(items ...T) bool {
	for _, item := range items {
		if s.Has(item) {
			return true
		}
	}
	return false
}

// Union returns a new set with all the elements in both s and s2.
func (s Set[T]) Union(s2 Set[T]) Set[T] {
	result := Set[T]{}
	result.Adds(s.ToSlice()...)
	result.Adds(s2.ToSlice()...)
	return result
}

// Len returns the number of elements in the set.
func (s Set[T]) Len() int {
	return len(s)
}

// Intersection returns a new set with elements that are in both s and s2.
func (s Set[T]) Intersection(s2 Set[T]) Set[T] {
	var walk, other Set[T]
	result := Set[T]{}
	if s.Len() < s2.Len() {
		walk = s
		other = s2
	} else {
		walk = s2
		other = s
	}
	for key := range walk {
		if other.Has(key) {
			result.Add(key)
		}
	}
	return result
}

// IsSuperset returns true if s is a superset of s2.
func (s Set[T]) IsSuperset(s2 Set[T]) bool {
	for item := range s2 {
		if !s.Has(item) {
			return false
		}
	}
	return true
}

// Difference returns a new set with elements in s but not in s2.
func (s Set[T]) Difference(s2 Set[T]) Set[T] {
	result := Set[T]{}
	for key := range s {
		if !s2.Has(key) {
			result.Add(key)
		}
	}
	return result
}

// Equal returns true if s and s2 are equal.
func (s Set[T]) Equal(s2 Set[T]) bool {
	return s.Len() == s2.Len() && s.IsSuperset(s2)
}

// ToSlice returns the elements of the set as a slice.
func (s Set[T]) ToSlice() []T {
	res := make([]T, 0, len(s))
	for key := range s {
		res = append(res, key)
	}
	return res
}

// Pop removes and returns an arbitrary item from the set.
func (s Set[T]) Pop() (T, bool) {
	for key := range s {
		s.Delete(key)
		return key, true
	}
	var zeroValue T
	return zeroValue, false
}

// Clone returns a new set with the same elements.
func (s Set[T]) Clone() Set[T] {
	result := make(Set[T], len(s))
	for key := range s {
		result.Add(key)
	}
	return result
}

// Clear removes all elements from the set.
func (s Set[T]) Clear() {
	for key := range s {
		delete(s, key)
	}
}

// SymmetricDifference returns a new set with elements that are in either of the sets and not in their intersection.
func (s Set[T]) SymmetricDifference(s2 Set[T]) Set[T] {
	return s.Difference(s2).Union(s2.Difference(s))
}
