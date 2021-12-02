package _2

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

func SolvePart1() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(calculatePosition(v))
}

func calculatePosition(instructions []string) int {
	depth, pos := 0, 0
	for _, i := range instructions {
		s := strings.Split(i, " ")
		dir := s[0]
		v, _ := strconv.Atoi(s[1])
		switch dir {
		case "forward":
			pos += v
		case "down":
			depth += v
		case "up":
			depth -= v
		}
	}
	return depth * pos
}
