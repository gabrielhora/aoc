package slice

// Split splits a list into list of lists with size `size`
func Split[T any](input []T, size int) [][]T {
	var r [][]T
	for i := 0; i < len(input); i += size {
		if i+size > len(input) {
			r = append(r, input[i:])
		} else {
			r = append(r, input[i:i+size])
		}
	}
	return r
}

func SlidingWindow[T any](input []T, size int) [][]T {
	var r [][]T
	for i, j := 0, size; j <= len(input); i, j = i+1, j+1 {
		r = append(r, input[i:j])
	}
	return r
}

func Reverse[T any](input []T) []T {
	res := make([]T, len(input))
	copy(res, input)
	for i := len(res)/2 - 1; i >= 0; i-- {
		opp := len(res) - 1 - i
		res[i], res[opp] = res[opp], res[i]
	}
	return res
}

func FindIndex[T comparable](lst []T, val T) int {
	for i, v := range lst {
		if v == val {
			return i
		}
	}
	return -1
}

func Runes(str string) []rune {
	var r []rune
	for _, l := range str {
		r = append(r, l)
	}
	return r
}
