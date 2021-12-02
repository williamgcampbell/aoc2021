package _1

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

func SolvePart2() string {
	r := strings.NewReader(input)
	depths, err := scanner.ScanIntLines(r)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error scanning values: %v\n", err)
		return ""
	}
	return strconv.Itoa(slidingWindowDepthIncreaseFast(depths))
}

// slidingWindowDepthIncrease solves part two, using elements of part one
func slidingWindowDepthIncrease(depths []int) int {
	slidingDepths := make([]int, 0, len(depths)/3)
	for i := 2; i < len(depths); i++ {
		slidingDepths = append(slidingDepths, depths[i]+depths[i-1]+depths[i-2])
	}
	return depthIncrease(slidingDepths)
}

// slidingWindowDepthIncreaseFast solves part two with one less loop
func slidingWindowDepthIncreaseFast(depths []int) int {
	count := 0
	for i := 3; i < len(depths); i++ {
		one := depths[i-1] + depths[i-2] + depths[i-3]
		two := depths[i] + depths[i-1] + depths[i-2]
		if two > one {
			count++
		}
	}
	return count
}
