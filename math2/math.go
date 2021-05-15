package math2

import (
	"golang.org/x/exp/constraints"
)

// Max 最大值
func Max[T constraints.Integer | constraints.Float](numbers ...T) T {
	max := numbers[0]

	for _, v := range numbers {
		if max < v {
			max = v
		}
	}

	return max
}

// Min 最小值
func Min[T constraints.Integer | constraints.Float](numbers ...T) T {
	min := numbers[0]

	for _, v := range numbers {
		if min > v {
			min = v
		}
	}

	return min
}

// Sum 求和
func Sum[T constraints.Integer | constraints.Float](numbers ...T) T {
	var sum T

	for _, v := range numbers {
		sum += v
	}

	return sum
}

// Average 平均值
func Average[T constraints.Integer | constraints.Float](numbers ...T) T {
	var sum T
	n := T(len(numbers))

	for _, v := range numbers {
		sum += v
	}
	return sum / n
}

// Abs 绝对值
func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}

	return x
}

// Div 除法
func Div[T constraints.Float | constraints.Integer](x T, y T) float64 {
	return float64(x) / float64(y)
}
