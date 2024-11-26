package map2

import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want []K
	}
	tests := []testCase[string, int]{
		{
			name: "Test Keys",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}},
			want: []string{"a", "b", "c"},
		},
		{
			name: "Test Keys",
			args: args[string, int]{m: map[string]int{}},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Keys(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Keys() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestKeysSorted(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want []K
	}
	tests := []testCase[string, int]{
		{
			name: "Test KeysSorted",
			args: args[string, int]{m: map[string]int{"b": 2, "a": 1, "c": 3}},
			want: []string{"a", "b", "c"},
		},
		{
			name: "Test KeysSorted",
			args: args[string, int]{m: map[string]int{}},
			want: []string{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KeysSorted(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("KeysSorted() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValues(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want []V
	}
	tests := []testCase[string, int]{
		{
			name: "Test Values",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}},
			want: []int{1, 2, 3},
		},
		{
			name: "Test Values",
			args: args[string, int]{m: map[string]int{}},
			want: []int{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Values(tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Values() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsEmpty(t *testing.T) {
	type args[K comparable, V any] struct {
		m map[K]V
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want bool
	}
	tests := []testCase[string, int]{
		{
			name: "Test IsEmpty",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}},
			want: false,
		},
		{
			name: "Test IsEmpty",
			args: args[string, int]{m: map[string]int{}},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsEmpty(tt.args.m); got != tt.want {
				t.Errorf("IsEmpty() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasKey(t *testing.T) {
	type args[K comparable, V any] struct {
		m   map[K]V
		key K
	}
	type testCase[K comparable, V any] struct {
		name string
		args args[K, V]
		want bool
	}
	tests := []testCase[string, int]{
		{
			name: "Test HasKey",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}, key: "a"},
			want: true,
		},
		{
			name: "Test HasKey",
			args: args[string, int]{m: map[string]int{"a": 1, "b": 2, "c": 3}, key: "d"},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasKey(tt.args.m, tt.args.key); got != tt.want {
				t.Errorf("HasKey() = %v, want %v", got, tt.want)
			}
		})
	}
}
