package main

import (
	_ "embed"
	"fmt"

	"hora.dev/aoc/utils"
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
	data, _ := parse(input)
	ans := checksum(compact(data))
	fmt.Printf("part 1: %v\n", ans)
}

func part2(input string) {
	data, maxFileID := parse(input)
	ans := checksum(compactBlock(data, maxFileID))
	fmt.Printf("part 2: %v\n", ans)
}

func checksum(data []int) int64 {
	var result int64
	for i, val := range data {
		if val != -1 {
			result += int64(i) * int64(val)
		}
	}
	return result
}

func compact(data []int) []int {
	freeIdx := nextFreeIdx(data, 0)
	for i := len(data) - 1; i >= 0; i-- {
		if freeIdx >= i {
			break
		}
		data[i], data[freeIdx] = data[freeIdx], data[i]
		freeIdx = nextFreeIdx(data, freeIdx)
	}
	return data
}

func nextFreeIdx(data []int, start int) int {
	for i := start; i < len(data); i++ {
		if data[i] == -1 {
			return i
		}
	}
	panic("no free space?")
}

func compactBlock(data []int, maxFileID int) []int {
	// needle moving backwards in `data`
	curPos := len(data) - 1

	for fileID := maxFileID; fileID >= 0; fileID-- {
		// find fileID positions
		fileBlockStart, fileBlockEnd, fileBlockSize := -1, -1, 0
		for ; curPos >= 0; curPos-- {
			if fileBlockEnd == -1 && data[curPos] == fileID {
				fileBlockEnd = curPos
				continue
			}
			if fileBlockEnd > 0 && data[curPos] != fileID {
				fileBlockStart = curPos + 1
				fileBlockSize = fileBlockEnd - fileBlockStart + 1
				break
			}
		}

		freeStart := nextFreeBlock(data, fileBlockSize, fileBlockStart)
		if freeStart == -1 {
			continue // no free block big enough
		}

		// swap whole block
		for i := 0; i < fileBlockSize; i++ {
			data[freeStart+i], data[fileBlockStart+i] = data[fileBlockStart+i], data[freeStart+i]
		}
	}
	return data
}

func nextFreeBlock(data []int, size, cutoff int) int {
	for i := 0; i < cutoff; i++ {
		if data[i] == -1 {
			for j := i + 1; j < len(data); j++ {
				if data[j] != -1 {
					if j-i >= size {
						return i
					}
					break
				}
			}
		}
	}
	return -1
}

func parse(input string) ([]int, int) {
	ints := utils.IntList(input, "")
	var expanded []int
	idx := 0
	for i, val := range ints {
		if i%2 == 0 {
			for range val {
				expanded = append(expanded, idx)
			}
			idx += 1
		} else {
			for range val {
				expanded = append(expanded, -1)
			}
		}
	}
	return expanded, idx - 1
}
