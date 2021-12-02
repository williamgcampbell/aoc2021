package _2

import (
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

func SolvePart2() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(calculatePositionWithAim(v))
}

func calculatePositionWithAim(instructions []string) int {
	depth, pos, aim := 0, 0, 0
	for _, i := range instructions {
		s := strings.Split(i, " ")
		dir := s[0]
		v, _ := strconv.Atoi(s[1])
		switch dir {
		case "forward":
			pos += v
			depth += aim * v
		case "down":
			aim += v
		case "up":
			aim -= v
		}
	}
	return depth * pos
}
