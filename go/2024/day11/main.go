package main

import (
	_ "embed"
	"fmt"
	"hora.dev/aoc/utils"
	"sync"
	"time"
)

//go:embed sample.txt
var sample string

//go:embed input.txt
var input string

func main() {
	part1(input)
	part2(input)
}

func part1(input string) {
	stones := utils.IntList(input, " ")
	start := time.Now()
	var ans int64
	for _, stone := range stones {
		ans += blink(stone, 25)
	}
	fmt.Printf("part 1: %+v, %v\n", ans, time.Since(start))
}

func part2(input string) {
	stones := utils.IntList(input, " ")
	start := time.Now()
	results := make(chan int64, len(stones))
	var wg sync.WaitGroup
	wg.Add(len(stones))
	for _, stone := range stones {
		go func() {
			defer wg.Done()
			results <- blink(stone, 75)
		}()
	}
	go func() {
		wg.Wait()
		close(results)
	}()
	var ans int64
	for result := range results {
		ans += result
	}
	fmt.Printf("part 2: %v, %+v\n", ans, time.Since(start))
}

var memoLock sync.RWMutex
var memo = map[[2]int64]int64{}

func blink(stone int64, n int64) (res int64) {
	key := [2]int64{n, stone}

	memoLock.RLock()
	v, ok := memo[key]
	memoLock.RUnlock()

	if ok {
		return v
	} else {
		defer func() {
			memoLock.Lock()
			memo[key] = res
			memoLock.Unlock()
		}()
	}

	if n == 0 {
		res = 1
		return
	}

	if stone == 0 {
		res = blink(1, n-1)
		return
	}

	digits := countDigits(stone)
	if digits%2 == 0 {
		left, right := splitNums(stone, digits)
		res = blink(left, n-1) + blink(right, n-1)
		return
	}

	res = blink(stone*2024, n-1)
	return
}

func countDigits(num int64) int64 {
	var count int64
	for num != 0 {
		num = num / 10
		count += 1
	}
	return count
}

func setDigit(num int64, pos int64, rep int64) int64 {
	return num%pos + (rep * pos) + ((num/pos)/10)*pos*10
}

func splitNums(num int64, digits int64) (int64, int64) {
	var left, right int64
	for i := int64(0); i < digits; i++ {
		digit := num % 10
		if i < digits/2 {
			right = setDigit(right, utils.Pow(10, i), digit)
		} else {
			left = setDigit(left, utils.Pow(10, i%(digits/2)), digit)
		}
		num = num / 10
	}
	return left, right
}
