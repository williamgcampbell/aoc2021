package _7

import (
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

// SolvePart2 solves part two
func (d *Day) SolvePart2() string {
	r := strings.NewReader(input)
	positions, _ := scanner.ScanCSVInt(r)
	return strconv.Itoa(minimumFuelSpent(positions[0], false))
}
