package set

import "golang.org/x/exp/constraints"

type Set[T comparable] map[T]any

func (s Set[T]) Push(val T) {
	s[val] = nil
}

// Intersection returns a new Set with values that are in both sets
func (s Set[T]) Intersection(b Set[T]) Set[T] {
	inter := Set[T]{}
	for x := range s {
		if _, ok := b[x]; ok {
			inter[x] = nil
		}
	}
	return inter
}

// Slice converts the Set into a slice
func (s Set[T]) Slice() []T {
	var ks []T
	for k, _ := range s {
		ks = append(ks, k)
	}
	return ks
}

// Equal checks if `other` has same elements as `s`
func (s Set[T]) Equal(other Set[T]) bool {
	if len(s) != len(other) {
		return false
	}
	for k, _ := range s {
		if _, ok := other[k]; !ok {
			return false
		}
	}
	return true
}

func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// FromValues creates a set from a list of values
func FromValues[T comparable](vals ...T) Set[T] {
	s := Set[T]{}
	for _, v := range vals {
		s[v] = nil
	}
	return s
}

// FromRange creates a set from `start` to `finish` (inclusive on both ends)
func FromRange[T constraints.Integer](start, finish T) Set[T] {
	s := Set[T]{}
	for i := start; i <= finish; i++ {
		s[i] = nil
	}
	return s
}

// Runes turns a string into a set of runes
func Runes(str string) Set[rune] {
	res := Set[rune]{}
	for _, l := range str {
		res[l] = nil
	}
	return res
}
