package set

import (
	"reflect"
	"testing"
)

func TestSet(t *testing.T) {
	s := New[int](1, 2, 3, 4, 5)
	added := s.Add(4)
	if added {
		t.Fatalf("expected 4 to not be added")
	}
	added = s.Add(6)
	if !added {
		t.Fatalf("expected 6 to be added")
	}

	if !s.Has(6) {
		t.Fatalf("expected 6 to be in the set")
	}

	if s.Len() != 6 {
		t.Fatalf("expected length to be 6")
	}

	if !s.HasAll(1, 2, 3, 4, 5, 6) {
		t.Fatalf("expected all items to be in the set")
	}

	if !s.Delete(6) {
		t.Fatalf("expected 6 to be deleted")
	}

	if s.Has(6) {
		t.Fatalf("expected 6 to not be in the set")
	}

	if s.Len() != 5 {
		t.Fatalf("expected length to be 5")
	}

	if s.HasAll(1, 2, 3, 4, 5, 6) {
		t.Fatalf("expected all items to not be in the set")
	}

	if !s.HasAll(1, 2, 3, 4) {
		t.Fatalf("expected all items to be in the set")
	}

	if !s.HasAny(1, 2, 3, 4, 5) {
		t.Fatalf("expected any item to be in the set")
	}

	if !s.HasAny(1, 2, 3, 4, 5, 6) {
		t.Fatalf("expected any item to be in the set")
	}

	// s = {1, 2, 3, 4, 5}
	s2 := New[int](4, 5, 6, 7, 8)

	if s.Equal(s2) {
		t.Fatalf("expected s2 and s to not be equal")
	}

	if s.IsSuperset(s2) {
		t.Fatalf("expected s2 to not be a subset of s")
	}

	if !s.IsSuperset(New[int](1, 2, 3)) {
		t.Fatalf("expected s to be a superset of {1, 2, 3}")
	}

	if !reflect.DeepEqual(s.Difference(s2), New[int](1, 2, 3)) {
		t.Fatalf("expected difference to be {1, 2, 3}")
	}

	if !reflect.DeepEqual(s.SymmetricDifference(s2), New[int](1, 2, 3, 6, 7, 8)) {
		t.Fatalf("expected symmetric difference to be {1, 2, 3, 6, 7, 8}")
	}

	if !reflect.DeepEqual(s.Union(s2), New[int](1, 2, 3, 4, 5, 6, 7, 8)) {
		t.Fatalf("expected union to be {1, 2, 3, 4, 5, 6, 7, 8}")
	}

	if !reflect.DeepEqual(s.Intersection(s2), New[int](4, 5)) {
		t.Fatalf("expected intersection to be {4, 5}")
	}

	// test clone method
	s3 := s.Clone()
	if !s3.Equal(s) {
		t.Fatalf("expected s3 to be equal to s")
	}

	// test clear method
	s3.Clear()
	if s3.Len() != 0 {
		t.Fatalf("expected s3 to be empty")
	}

	// test ToSlice method
	if reflect.DeepEqual(s.ToSlice(), []int{1, 2, 3, 4, 5}) {
		t.Fatalf("expected ToSlice to return [1, 2, 3, 4, 5]")
	}

	// test Pop method
	if _, ok := s.Pop(); !ok {
		t.Fatalf("expected Pop to return an element")
	}

	// test Pop method on empty set
	s4 := New[int]()
	if _, ok := s4.Pop(); ok {
		t.Fatalf("expected Pop to return false")
	}
}
