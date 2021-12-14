package _13

import (
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

// SolvePart2 solves part two
func (d *Day) SolvePart2() string {
	r := strings.NewReader(input)
	v := scanner.ScanUntilEmptyLine(r, delimiter)
	return decode(v)
}

func decode(lines []string) string {
	dots := strings.Split(lines[0], delimiter)
	coords := make(map[string]struct{})
	for _, dot := range dots {
		coords[dot] = present
	}

	instructions := strings.Split(lines[1], delimiter)
	for _, instruction := range instructions {
		coords = fold(coords, instruction)
	}

	return asciiString(coords)
}

func asciiString(m map[string]struct{}) string {

	mx := 0
	my := 0
	for c, _ := range m {
		coord := newCoord(c)
		if coord.x > mx {
			mx = coord.x
		}
		if coord.y > my {
			my = coord.y
		}
	}

	r := strings.Builder{}
	r.WriteString("\n")
	for y := 0; y <= my; y++ {
		for x := 0; x <= mx; x++ {
			c := &coordinate{x, y}
			if _, ok := m[c.String()]; ok {
				r.WriteRune('#')
			} else {
				r.WriteRune(' ')
			}
		}
		r.WriteString("\n")
	}
	return r.String()
}
