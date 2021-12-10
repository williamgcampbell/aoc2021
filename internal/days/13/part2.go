package _13

import (
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

func (d *Day) SolvePart2() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(todo2(v))
}

func todo2(s []string) int {
	return 0
}
