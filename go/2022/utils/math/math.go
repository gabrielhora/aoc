package math

import (
	"math"

	"golang.org/x/exp/constraints"
)

func Max[T constraints.Integer](a, b T) T {
	return T(math.Max(float64(a), float64(b)))
}

func Min[T constraints.Integer](a, b T) T {
	return T(math.Min(float64(a), float64(b)))
}

func Abs[T constraints.Integer](x T) T {
	return T(math.Abs(float64(x)))
}

func Mod[T constraints.Integer](x, y T) T {
	if x > 0 {
		return x % y
	}
	if x < 0 {
		return (x % y) + y
	}
	return 0
}
