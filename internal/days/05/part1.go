package _5

import (
	_ "embed"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

var format, _ = regexp.Compile("(?P<x1>[0-9]+),(?P<y1>[0-9]+) -> (?P<x2>[0-9]+),(?P<y2>[0-9]+)")

// SolvePart1 solves part one
func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	lines := scanner.ScanLines(r)
	return strconv.Itoa(overlappingLines(lines, true))
}

func overlappingLines(lines []string, skipDiagonal bool) int {
	d := newDiagram()
	for _, line := range lines {
		c1, c2 := extractCoordinate(line)
		if !skipDiagonal || !diagonalLine(c1, c2) {
			d.addLine(c1, c2)
		}
	}
	return d.atLeastTwoCount
}

func extractCoordinate(line string) (coordinate, coordinate) {
	coords := format.FindStringSubmatch(line)
	x1, _ := strconv.Atoi(coords[1])
	y1, _ := strconv.Atoi(coords[2])
	x2, _ := strconv.Atoi(coords[3])
	y2, _ := strconv.Atoi(coords[4])
	return coordinate{x1, y1}, coordinate{x2, y2}
}

func diagonalLine(c1, c2 coordinate) bool {
	return c1.x != c2.x && c1.y != c2.y
}

func newDiagram() *diagram {
	return &diagram{
		matrix:          make(map[string]int),
		atLeastTwoCount: 0,
	}
}

type diagram struct {
	matrix map[string]int

	atLeastTwoCount int
}

func (d *diagram) addLine(start, end coordinate) {
	l := newLineIt(start, end)
	for {
		p, stop := l.Next()
		coord := fmt.Sprintf("%d,%d", p.x, p.y)
		if v, ok := d.matrix[coord]; ok {
			d.matrix[coord]++
			if v == 1 {
				d.atLeastTwoCount++
			}
		} else {
			d.matrix[coord] = 1
		}

		if stop {
			return
		}
	}
}

func newLineIt(start, end coordinate) *lineIt {
	return &lineIt{
		start: start,
		end:   end,
		next:  start,
	}
}

// lineIt is a straight line iterator.
type lineIt struct {
	start coordinate
	end   coordinate
	next  coordinate
}

// Next returns the next coordinate and whether the end coordinate has been reached
func (l *lineIt) Next() (p coordinate, stop bool) {
	r := l.next
	if l.next.x != l.end.x || l.next.y != l.end.y {
		l.next.x = nextCoord(l.next.x, l.end.x)
		l.next.y = nextCoord(l.next.y, l.end.y)
		return r, false
	}
	return r, true
}

func nextCoord(v int, end int) int {
	if v != end {
		if v < end {
			return v + 1
		}
		return v - 1
	}
	return v
}

type coordinate struct {
	x, y int
}
