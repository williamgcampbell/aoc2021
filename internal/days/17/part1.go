package _17

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(todo(v))
}

func todo(s []string) int {
	return 0
}
