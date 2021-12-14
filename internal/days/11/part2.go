package _11

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
	return strconv.Itoa(allFlash(v))
}

func allFlash(s []string) int {
	plane := helpers.PlanarInt(s)

	stepCount := 0
	for {
		flashes := step(plane)
		stepCount += 1
		if flashes == 100 {
			return stepCount
		}
	}
}
