package _11

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/helpers"
	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

// SolvePart1 solves part one
func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(totalFlashes(v, 100))
}

func totalFlashes(s []string, afterSteps int) int {
	plane := helpers.PlanarInt(s)

	flashes := 0
	for i := 0; i < afterSteps; i++ {
		fs := step(plane)
		flashes += fs
	}
	return flashes
}

func step(plane [][]int) int {
	for i, row := range plane {
		for j := range row {
			plane[i][j] += 1
		}
	}

	flashed := make([][]bool, len(plane))
	for i, v := range plane {
		flashed[i] = make([]bool, len(v))
	}

	count := 0
	for i, row := range plane {
		for j := range row {
			count += flash(plane, i, j, flashed)
		}
	}

	for i, row := range plane {
		for j := range row {
			if plane[i][j] > 9 {
				plane[i][j] = 0
			}
		}
	}
	return count
}

// flash performs a "flash", which increments all adjacent octopuses
func flash(plane [][]int, row, col int, flashed [][]bool) int {
	if !shouldFlash(plane[row][col]) || flashed[row][col] {
		return 0
	}

	// clever way to represent all eight locations surrounding a value in a 2D space
	rn := []int{-1, -1, -1, 0, 0, 1, 1, 1}
	cn := []int{0, -1, 1, -1, 1, 1, -1, 0}

	flashed[row][col] = true

	count := 1
	for k := 0; k < len(rn); k++ {
		rni := row + rn[k]
		cni := col + cn[k]
		if helpers.Safe(plane, rni, cni) && !flashed[rni][cni] {
			plane[rni][cni] += 1
			count += flash(plane, rni, cni, flashed)
		}
	}
	return count
}

func shouldFlash(v int) bool {
	return v > 9
}
