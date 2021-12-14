package _2

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

// SolvePart1 solves part one
func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(calculatePosition(v, false))
}

func calculatePosition(instructions []string, withAim bool) int {
	depth, pos, aim := 0, 0, 0
	for _, i := range instructions {
		s := strings.Split(i, " ")
		dir := s[0]
		v, _ := strconv.Atoi(s[1])
		switch dir {
		case "forward":
			pos += v
			if withAim {
				depth += aim * v
			}
		case "down":
			if withAim {
				aim += v
			} else {
				depth += v
			}
		case "up":
			if withAim {
				aim -= v
			} else {
				depth -= v
			}
		}
	}
	return depth * pos
}
