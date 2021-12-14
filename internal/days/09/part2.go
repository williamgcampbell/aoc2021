package _9

import (
	"sort"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/helpers"
	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

// SolvePart2 solves part two
func (d *Day) SolvePart2() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(threeLargestBasins(v))
}

func threeLargestBasins(lines []string) int {
	hm := helpers.PlanarInt(lines)

	visited := make([][]bool, len(hm))
	for i, v := range hm {
		visited[i] = make([]bool, len(v))
	}

	sizes := make([]int, 0)
	for i, row := range hm {
		for j, height := range row {
			if inBasin(height) && !visited[i][j] {
				size := walkBasin(hm, i, j, visited)
				sizes = append(sizes, size)
			}
		}
	}

	// sorted in ascending order
	sort.Ints(sizes)
	return sizes[len(sizes)-1] * sizes[len(sizes)-2] * sizes[len(sizes)-3]
}

// walkBasin will do a depth first search of a basin, starting from a location, and
// returns the size of the basin.
// Visited is used to identify if a location on the map has already been walked.
func walkBasin(hm [][]int, row, col int, visited [][]bool) int {
	// clever way to represent all four locations surrounding a value in a 2D space
	rn := []int{-1, 0, 0, 1}
	cn := []int{0, -1, 1, 0}

	visited[row][col] = true

	count := 1
	for k := 0; k < 4; k++ {
		rni := row + rn[k]
		cni := col + cn[k]
		if helpers.Safe(hm, rni, cni) {
			if inBasin(hm[rni][cni]) && !visited[rni][cni] {
				count += walkBasin(hm, rni, cni, visited)
			}
		}
	}
	return count
}

func inBasin(height int) bool {
	return height != 9
}
