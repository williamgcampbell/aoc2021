package _16

import (
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

// SolvePart2 solves part two
func (d *Day) SolvePart2() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(evaluate(v[0]))
}

func evaluate(hex string) int {
	bitStr := toBits(hex)
	parser := newPacketParser(bitStr)
	packets := ParseCount(parser, 1)
	return packets[0].Value()
}
