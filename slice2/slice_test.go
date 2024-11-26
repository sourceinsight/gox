package slice2

import (
	"reflect"
	"testing"
)

func TestContains(t *testing.T) {
	type args[T comparable] struct {
		ss      []T
		element T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "Test Contains",
			args: args[int]{ss: []int{0, 1, 2, 3, 4, 5}, element: 5},
			want: true,
		},
		{
			name: "Test Contains",
			args: args[int]{ss: []int{0, 1, 2, 3, 4, 5}, element: 6},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Contains(tt.args.ss, tt.args.element); got != tt.want {
				t.Errorf("Contains() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunk(t *testing.T) {
	type args[T any] struct {
		ss   []T
		size int
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[int]{
		{
			name: "Test Chunk",
			args: args[int]{ss: []int{0, 1, 2, 3, 4}, size: 2},
			want: [][]int{{0, 1}, {2, 3}, {4}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Chunk(tt.args.ss, tt.args.size); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Chunk() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCount(t *testing.T) {
	type args[T comparable] struct {
		ss    []T
		value T
	}
	type testCase[T comparable] struct {
		name      string
		args      args[T]
		wantCount int
	}
	tests := []testCase[int]{
		{
			name:      "Test Count",
			args:      args[int]{ss: []int{0, 1, 2, 3, 4, 5}, value: 5},
			wantCount: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotCount := Count(tt.args.ss, tt.args.value); gotCount != tt.wantCount {
				t.Errorf("Count() = %v, want %v", gotCount, tt.wantCount)
			}
		})
	}
}

func TestFilter(t *testing.T) {
	type args[V any] struct {
		ss        []V
		predicate func(item V, index int) bool
	}
	type testCase[V any] struct {
		name string
		args args[V]
		want []V
	}
	tests := []testCase[int]{
		{
			name: "Test Filter",
			args: args[int]{ss: []int{0, 1, 2, 3, 4, 5}, predicate: func(item int, index int) bool {
				return item%2 == 0
			},
			},
			want: []int{0, 2, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Filter(tt.args.ss, tt.args.predicate); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Filter() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMap(t *testing.T) {
	type args[T any, R any] struct {
		ss       []T
		iteratee func(item T, index int) R
	}
	type testCase[T any, R any] struct {
		name string
		args args[T, R]
		want []R
	}
	tests := []testCase[int, int]{
		{
			name: "Test Map",
			args: args[int, int]{ss: []int{0, 1, 2, 3, 4, 5}, iteratee: func(item int, index int) int {
				return item * 2

			},
			},
			want: []int{0, 2, 4, 6, 8, 10},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Map(tt.args.ss, tt.args.iteratee); !reflect.DeepEqual(got, tt.want) {
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

func TestUniqStable(t *testing.T) {
	type args[T comparable] struct {
		ss []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "Test Uniq",
			args: args[int]{ss: []int{0, 1, 2, 3, 4, 5, 5}},
			want: []int{0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UniqStable(tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Uniq() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirst(t *testing.T) {
	type args[T any] struct {
		ss []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Test First",
			args: args[int]{ss: []int{0, 1, 2, 3, 4, 5}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := First(tt.args.ss); got != tt.want {
				t.Errorf("First() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFirstOr(t *testing.T) {
	type args[T any] struct {
		ss []T
		or T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Test FirstOr",
			args: args[int]{ss: []int{0, 1, 2, 3, 4, 5}, or: 10},
			want: 0,
		},
		{
			name: "Test FirstOr",
			args: args[int]{ss: []int{}, or: 10},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstOr(tt.args.ss, tt.args.or); got != tt.want {
				t.Errorf("FirstOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLast(t *testing.T) {
	type args[T any] struct {
		ss []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Test Last",
			args: args[int]{ss: []int{0, 1, 2, 3, 4, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Last(tt.args.ss); got != tt.want {
				t.Errorf("Last() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLastOr(t *testing.T) {
	type args[T any] struct {
		ss []T
		or T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Test LastOr",
			args: args[int]{ss: []int{0, 1, 2, 3, 4, 5}, or: 10},
			want: 5,
		},
		{
			name: "Test LastOr",
			args: args[int]{ss: []int{}, or: 10},
			want: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LastOr(tt.args.ss, tt.args.or); got != tt.want {
				t.Errorf("LastOr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type args[T any] struct {
		ss []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "empty",
			args: args[int]{ss: []int{}},
			want: true,
		},
		{
			name: "not emtpy",
			args: args[int]{ss: []int{0, 1, 2, 3, 4, 5}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.ss); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestClone(t *testing.T) {
	type args[T any] struct {
		ss []T
	}
	type testCase[T any] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "Test Clone",
			args: args[int]{ss: []int{0, 1, 2, 3, 4, 5}},
			want: []int{0, 1, 2, 3, 4, 5},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Clone(tt.args.ss); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Clone() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqual(t *testing.T) {
	type args[T comparable] struct {
		s1 []T
		s2 []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want bool
	}
	tests := []testCase[int]{
		{
			name: "equal",
			args: args[int]{s1: []int{0, 1, 2, 3, 4, 5}, s2: []int{0, 1, 2, 3, 4, 5}},
			want: true,
		},
		{
			name: "not equal",
			args: args[int]{s1: []int{0, 1, 2, 3, 4, 5}, s2: []int{0, 1, 2, 3, 4}},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Equal(tt.args.s1, tt.args.s2); got != tt.want {
				t.Errorf("Equal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiff(t *testing.T) {
	type args[T comparable] struct {
		s1 []T
		s2 []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want [][]T
	}
	tests := []testCase[int]{
		{
			name: "Test Diff",
			args: args[int]{s1: []int{1, 2, 3, 4}, s2: []int{3, 4, 5, 6}},
			want: [][]int{{1, 2}, {5, 6}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got1, got2 := Diff(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got1, tt.want[0]) || !reflect.DeepEqual(got2, tt.want[1]) {
				t.Errorf("Diff() = %v, %v, want %v", got1, got2, tt.want)
			}
		})
	}
}

func TestMerge(t *testing.T) {
	type args[T comparable] struct {
		s1 []T
		s2 []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want []T
	}
	tests := []testCase[int]{
		{
			name: "Test Merge",
			args: args[int]{s1: []int{1, 2, 3, 4}, s2: []int{3, 4, 5, 6}},
			want: []int{1, 2, 3, 4, 5, 6},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Merge(tt.args.s1, tt.args.s2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Merge() = %v, want %v", got, tt.want)
			}
		})
	}
}
