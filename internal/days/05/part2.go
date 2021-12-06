package _5

import (
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

func (d *Day) SolvePart2() string {
	r := strings.NewReader(input)
	lines := scanner.ScanLines(r)
	return strconv.Itoa(overlappingLines(lines, false))
}
