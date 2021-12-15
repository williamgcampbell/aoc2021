package _14

import (
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

// SolvePart2 solves part two
func (d *Day) SolvePart2() string {
	r := strings.NewReader(input)
	v := scanner.ScanUntilEmptyLine(r, delimiter)
	return strconv.Itoa(calculatePolymer(v, 40))
}
