package main

import (
	_ "embed"
	"fmt"
	"math"
	"regexp"
	"strings"
)

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	inst, dir := parse(input)
	steps := walkMap("AAA", inst, dir, func(pos string) bool {
		return pos != "ZZZ"
	})
	fmt.Printf("part 1: %d\n", steps)
}

func part2(input string) {
	inst, dir := parse(input)

	var starts []string
	for k := range dir {
		if strings.HasSuffix(k, "A") {
			starts = append(starts, k)
		}
	}

	var steps []int64
	for _, s := range starts {
		ss := walkMap(s, inst, dir, func(s string) bool {
			return !strings.HasSuffix(s, "Z")
		})
		steps = append(steps, int64(ss))
	}

	// all paths will end in Z in the MMC of them

	fmt.Printf("part 2: %d\n", lcm(steps...))
}

// adapted from the internets :)
func lcm(numbers ...int64) int64 {
	gcd := func(a int64, b int64) int64 {
		for b != 0 {
			a, b = b, a%b
		}
		return a
	}
	lcmPair := func(a int64, b int64) int64 {
		return int64(math.Abs(float64(a*b))) / gcd(a, b)
	}

	result := int64(1)
	for _, num := range numbers {
		result = lcmPair(result, num)
	}
	return result
}

func walkMap(start string, inst []int, dir directions, stepFn func(string) bool) int {
	var curInst int
	var steps int
	pos := start

	for {
		steps += 1
		curInst, inst = inst[0], inst[1:]
		pos = dir[pos][curInst]
		if !stepFn(pos) {
			break
		}
		inst = append(inst, curInst) // append instruction back to the queue
	}

	return steps
}

// key = "AAA", value = [2]string{"BBB", "CCC"}
type directions map[string][2]string

var lineRe = regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

func parse(input string) ([]int, directions) {
	lines := strings.Split(input, "\n")

	var instructions []int
	for _, d := range strings.Split(lines[0], "") {
		if d == "L" {
			instructions = append(instructions, 0)
		} else {
			instructions = append(instructions, 1)
		}
	}

	dir := directions{}
	for _, line := range lines[2:] {
		m := lineRe.FindAllStringSubmatch(line, -1)
		dir[m[0][1]] = [2]string{m[0][2], m[0][3]}
	}

	return instructions, dir
}
