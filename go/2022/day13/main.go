package main

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"reflect"
	"sort"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	pairs := parse(input)
	ans := 0
	for i, pair := range pairs {
		if cmp(pair[0], pair[1]) {
			ans += i + 1
		}
	}
	fmt.Printf("part 1: %d\n", ans)
}

func part2(input string) {
	lst := [][]any{
		{[]any{float64(2)}},
		{[]any{float64(6)}},
	}
	for _, p := range parse(input) {
		lst = append(lst, p[0])
		lst = append(lst, p[1])
	}
	sort.Slice(lst, func(i, j int) bool { return cmp(lst[i], lst[j]) })

	ans := 1
	for i, l := range lst {
		if fmt.Sprint(l) == "[[2]]" || fmt.Sprint(l) == "[[6]]" {
			ans *= i + 1
		}
	}
	fmt.Printf("part 2: %d\n", ans)
}

func cmp(l, r []any) bool {
	if len(l) == 0 && len(r) > 0 {
		return true
	}
	if len(l) > 0 && len(r) == 0 {
		return false
	}
	if len(l) == 0 && len(r) == 0 {
		return true
	}

	leftNum, isLeftNum := l[0].(float64)
	rightNum, isRightNum := r[0].(float64)
	leftList, isLeftList := l[0].([]any)
	rightList, isRightList := r[0].([]any)

	if isLeftNum && isRightNum && leftNum < rightNum {
		return true
	}
	if isLeftNum && isRightNum && leftNum > rightNum {
		return false
	}
	if isLeftNum && isRightNum /* leftNum == rightNum */ {
		return cmp(l[1:], r[1:])
	}
	if isLeftNum && isRightList {
		newLeftList := []any{leftNum}
		newLeftList = append(newLeftList, leftList...)
		return cmp([]any{newLeftList}, r)
	}
	if isLeftList && isRightNum {
		newRightList := []any{rightNum}
		newRightList = append(newRightList, rightList...)
		return cmp(l, []any{newRightList})
	}
	if isLeftList && isRightList && reflect.DeepEqual(leftList, rightList) {
		return cmp(l[1:], r[1:])
	}
	if isLeftList && isRightList {
		return cmp(leftList, rightList)
	}
	panic("impossibru!")
}

// parse returns a list of pairs of float64 lists or list of pairs of lists
func parse(input string) [][2][]any {
	var res [][2][]any
	for _, pair := range strings.Split(input, "\n\n") {
		lst := strings.Split(pair, "\n")

		var lst1 []any
		_ = json.Unmarshal([]byte(lst[0]), &lst1)
		var lst2 []any
		_ = json.Unmarshal([]byte(lst[1]), &lst2)

		res = append(res, [2][]any{lst1, lst2})
	}
	return res
}
