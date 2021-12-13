package _13

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

var (
	delimeter            = ";"
	instructionFormat, _ = regexp.Compile("fold along (?P<dir>[x|y])=(?P<val>[0-9]+)")
	present              struct{}
)

func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	v := scanner.ScanUntilEmptyLine(r, delimeter)
	return strconv.Itoa(visibleDots(v, 1))
}

func visibleDots(lines []string, folds int) int {
	dots := strings.Split(lines[0], delimeter)
	coords := make(map[string]struct{})
	for _, dot := range dots {
		coords[dot] = present
	}

	instructions := strings.Split(lines[1], delimeter)
	for i, instruction := range instructions {
		if folds > i {
			coords = fold(coords, instruction)
		}
	}

	return len(coords)
}

// fold will fold the coordinates onto each other
// up for horizontal lines and left for vertical lines
func fold(coords map[string]struct{}, instruction string) map[string]struct{} {
	inst := instructionFormat.FindStringSubmatch(instruction)
	dir := inst[1]
	val, _ := strconv.Atoi(inst[2])
	m := make(map[string]struct{})
	for dot, _ := range coords {
		coord := newCoord(dot)
		if dir == "x" {
			if coord.x > val {
				coord.x = val*2 - coord.x
			}
		} else if dir == "y" {
			if coord.y > val {
				coord.y = val*2 - coord.y
			}
		}
		m[coord.String()] = present
	}
	return m
}

func newCoord(s string) *coordinate {
	coord := strings.Split(s, ",")
	x, _ := strconv.Atoi(coord[0])
	y, _ := strconv.Atoi(coord[1])
	return &coordinate{x, y}
}

type coordinate struct {
	x, y int
}

func (c *coordinate) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}
