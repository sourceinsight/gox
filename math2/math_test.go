package math2

import (
	"golang.org/x/exp/constraints"
	"testing"
)

func TestMax(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		numbers []T
	}
	type testCase[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Test Max",
			args: args[int]{numbers: []int{0, 1, 2, 3, 4, 5}},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Max(tt.args.numbers...); got != tt.want {
				t.Errorf("Max() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAbs(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		x T
	}
	type testCase[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Test Abs",
			args: args[int]{x: -1},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Abs(tt.args.x); got != tt.want {
				t.Errorf("Abs() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAverage(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		numbers []T
	}
	type testCase[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Test Average",
			args: args[int]{numbers: []int{0, 1, 2, 3, 4, 5}},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Average(tt.args.numbers...); got != tt.want {
				t.Errorf("Average() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDiv(t *testing.T) {
	type args[T interface {
		constraints.Float | constraints.Integer
	}] struct {
		x T
		y T
	}
	type testCase[T interface {
		constraints.Float | constraints.Integer
	}] struct {
		name string
		args args[T]
		want float64
	}
	tests := []testCase[int]{
		{
			name: "Test Div",
			args: args[int]{x: 1, y: 2},
			want: 0.5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Div(tt.args.x, tt.args.y); got != tt.want {
				t.Errorf("Div() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMin(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		numbers []T
	}
	type testCase[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Test Min",
			args: args[int]{numbers: []int{0, 1, 2, 3, 4, 5}},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Min(tt.args.numbers...); got != tt.want {
				t.Errorf("Min() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSum(t *testing.T) {
	type args[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		numbers []T
	}
	type testCase[T interface {
		constraints.Integer | constraints.Float
	}] struct {
		name string
		args args[T]
		want T
	}
	tests := []testCase[int]{
		{
			name: "Test Sum",
			args: args[int]{numbers: []int{0, 1, 2, 3, 4, 5}},
			want: 15,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sum(tt.args.numbers...); got != tt.want {
				t.Errorf("Sum() = %v, want %v", got, tt.want)
			}
		})
	}
}
