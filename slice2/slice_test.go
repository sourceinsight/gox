package slice2

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type args[T comparable] struct {
		collection []T
		element    T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Test Contains",
			args: args[int]{collection: []int{0, 1, 2, 3, 4, 5}, element: 5},
			want: true,
		},
		{
			name: "Test Contains",
			args: args[int]{collection: []int{0, 1, 2, 3, 4, 5}, element: 6},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.collection, tt.args.element); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	type args[T any] struct {
		collection []T
		size       int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[int]{
		{
			name: "Test Chunk",
			args: args[int]{collection: []int{0, 1, 2, 3, 4}, size: 2},
			want: [][]int{{0, 1}, {2, 3}, {4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.collection, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	type args[T comparable] struct {
		collection []T
		value      T
	}
	type testCase[T comparable] struct {
		name      string
		args      args[T]
		wantCount int
	}
	tests := []testCase[int]{
		{
			name:      "Test Count",
			args:      args[int]{collection: []int{0, 1, 2, 3, 4, 5}, value: 5},
			wantCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := Count(tt.args.collection, tt.args.value); gotCount != tt.wantCount {
				t.Errorf("Count() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args[V any] struct {
		collection []V
		predicate  func(item V, index int) bool
	}
	type testCase[V any] struct {
		name string
		args args[V]
		want []V
	}
	tests := []testCase[int]{
		{
			name: "Test Filter",
			args: args[int]{collection: []int{0, 1, 2, 3, 4, 5}, predicate: func(item int, index int) bool {
				return item%2 == 0
			},
			},
			want: []int{0, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.collection, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[T any, R any] struct {
		collection []T
		iteratee   func(item T, index int) R
	}
	type testCase[T any, R any] struct {
		name string
		args args[T, R]
		want []R
	}
	tests := []testCase[int, int]{
		{
			name: "Test Map",
			args: args[int, int]{collection: []int{0, 1, 2, 3, 4, 5}, iteratee: func(item int, index int) int {
				return item * 2

			},
			},
			want: []int{0, 2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.collection, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Map() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPruneEmptyString(t *testing.T) {
	type args struct {
		v []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "Test PruneEmptyString",
			args: args{v: []string{"", "1", "2", "", "3", ""}},
			want: []string{"1", "2", "3"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PruneEmptyString(tt.args.v); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PruneEmptyString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPruneEqual(t *testing.T) {
	type args[T comparable] struct {
		inputSlice []T
		equalTo    T
	}
	type testCase[T comparable] struct {
		name  string
		args  args[T]
		wantR []T
	}
	tests := []testCase[int]{
		{
			name:  "Test PruneEqual",
			args:  args[int]{inputSlice: []int{0, 1, 2, 3, 4, 5, 5}, equalTo: 5},
			wantR: []int{0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotR := PruneEqual(tt.args.inputSlice, tt.args.equalTo); !reflect.DeepEqual(gotR, tt.wantR) {
				t.Errorf("PruneEqual() = %v, want %v", gotR, tt.wantR)
			}
		})
	}
}

func TestUniq(t *testing.T) {
	type args[T comparable] struct {
		collection []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "Test Uniq",
			args: args[int]{collection: []int{0, 1, 2, 3, 4, 5, 5}},
			want: []int{0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Uniq(tt.args.collection); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uniq() = %v, want %v", got, tt.want)
			}
		})
	}
}
