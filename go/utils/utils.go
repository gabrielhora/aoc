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
