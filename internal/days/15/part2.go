package _15

import (
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/helpers"
	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

// SolvePart2 solves part two
func (d *Day) SolvePart2() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(lowestRisk2(v))
}

func lowestRisk2(lines []string) int {
	rvs := helpers.PlanarInt(lines)
	rvs = largeMap(rvs, 5)
	start := loc{
		row: 0,
		col: 0,
	}

	r := buildRiskScore(rvs, start)

	return r[len(r)-1][len(r[len(r)-1])-1]
}

func largeMap(original [][]int, times int) [][]int {
	r := make([][]int, len(original)*times)
	for i, row := range original {
		for t := 0; t < times; t++ {
			c := make([]int, len(row)*times)
			r[i+(len(original)*t)] = c
		}
	}
	// what have I done?!
	for i, row := range original {
		for rt := 0; rt < times; rt++ {
			nextR := i + (len(original) * rt)
			for j, val := range row {
				for ct := 0; ct < times; ct++ {
					nextC := j + (len(row) * ct)
					nextVal := (val+rt+ct)%10 + (val+rt+ct)/10
					r[nextR][nextC] = nextVal
				}
			}
		}
	}
	return r
}
