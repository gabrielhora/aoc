package utils

import (
	"math"
	"strconv"
	"strings"
)

func ToInt64(str string) int64 {
	val, err := strconv.ParseInt(str, 10, 64)
	if err != nil {
		panic("invalid int64 " + err.Error())
	}
	return val
}

func IntList[T int | int64](str, sep string) []T {
	var res []T
	for _, s := range strings.Split(strings.TrimSpace(str), sep) {
		res = append(res, T(ToInt64(s)))
	}
	return res
}

func GridStr(input string, lineSep, colSep string) [][]string {
	var grid [][]string
	for _, row := range strings.Split(input, lineSep) {
		line := append([]string{}, strings.Split(row, colSep)...)
		grid = append(grid, line)
	}
	return grid
}

func SlicesEqual[T comparable](slice1, slice2 []T) bool {
	if len(slice1) != len(slice2) {
		return false
	}
	for i := range slice1 {
		if slice1[i] != slice2[i] {
			return false
		}
	}
	return true
}

func Abs[T int | int64](a, b T) T {
	res := a - b
	if res < 0 {
		return res * -1
	}
	return res
}

func Pow[T int | int64](a, b T) T {
	return T(math.Pow(float64(a), float64(b)))
}

func CountDigits(num int64) int64 {
	var count int64
	for num != 0 {
		num = num / 10
		count += 1
	}
	return count
}

func ChunkNumber(num int64, digits int64, chunkSize int64) []int64 {
	var chunks []int64
	var acc int64
	var count int64
	num = reverseNum(num)
	for range digits {
		digit := num % 10
		acc = acc*10 + digit
		num /= 10
		count++
		if count == chunkSize {
			chunks = append(chunks, acc)
			acc, count = 0, 0
		}
	}
	if count > 0 {
		chunks = append(chunks, acc)
	}
	return chunks
}

func reverseNum(num int64) int64 {
	var reversed int64
	for num != 0 {
		digit := num % 10
		reversed = reversed*10 + digit
		num /= 10
	}
	return reversed
}

func setDigit(num int64, pos int64, rep int64) int64 {
	return num%pos + (rep * pos) + ((num/pos)/10)*pos*10
}
