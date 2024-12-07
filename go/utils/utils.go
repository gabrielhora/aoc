package utils

import (
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

func IntList(str, sep string) []int64 {
	var res []int64
	for _, s := range strings.Split(str, sep) {
		res = append(res, ToInt64(s))
	}
	return res
}

func GridStr(input string, lineSep, colSep string) [][]string {
	var grid [][]string
	for _, row := range strings.Split(input, lineSep) {
		var line []string
		for _, col := range strings.Split(row, colSep) {
			line = append(line, col)
		}
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
