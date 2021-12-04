package _4

import (
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

func SolvePart2() string {
	r := strings.NewReader(input)
	lines := scanner.ScanLines(r)
	return strconv.FormatInt(todo(lines), 10)
}

func todoPart2(lines []string) int64 {
	return 0
}
