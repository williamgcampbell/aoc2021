package _6

import (
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

var TotalDaysPart2 = 256

// SolvePart2 solves part two
func (d *Day) SolvePart2() string {
	r := strings.NewReader(input)
	lines, _ := scanner.ScanCSVInt(r)
	return strconv.FormatInt(countAll(lines[0], TotalDaysPart2), 10)
}
