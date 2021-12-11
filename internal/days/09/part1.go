package _9

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/helpers"
	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(sumOfRiskLevels(v))
}

func sumOfRiskLevels(lines []string) int {
	hm := helpers.PlanarInt(lines)

	sum := 0
	for i, row := range hm {
		for j, height := range row {
			if lowestPoint(hm, i, j, height) {
				sum += riskLevel(height)
			}
		}
	}
	return sum
}

func lowestPoint(hm [][]int, i, j, val int) bool {
	if helpers.Safe(hm, i-1, j) && val >= hm[i-1][j] {
		return false
	}
	if helpers.Safe(hm, i+1, j) && val >= hm[i+1][j] {
		return false
	}
	if helpers.Safe(hm, i, j-1) && val >= hm[i][j-1] {
		return false
	}
	if helpers.Safe(hm, i, j+1) && val >= hm[i][j+1] {
		return false
	}
	return true
}

func riskLevel(height int) int {
	return height + 1
}
