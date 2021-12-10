package _9

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
	return strconv.Itoa(sumOfRiskLevels(v))
}

func sumOfRiskLevels(lines []string) int {
	hm := heightMap(lines)

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

func heightMap(lines []string) [][]int {
	hm := make([][]int, 0, len(lines))
	for _, r := range lines {
		heightStrs := strings.Split(r, "")
		heights := make([]int, 0, len(heightStrs))
		for _, hs := range heightStrs {
			h, _ := strconv.Atoi(hs)
			heights = append(heights, h)
		}
		hm = append(hm, heights)
	}
	return hm
}

// safe returns true if i, j respectively are within the bounds of the two dimenstional array
func safe(hm [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(hm) || j >= len(hm[i]) {
		return false
	}
	return true
}

func lowestPoint(hm [][]int, i, j, val int) bool {
	if safe(hm, i-1, j) && val >= hm[i-1][j] {
		return false
	}
	if safe(hm, i+1, j) && val >= hm[i+1][j] {
		return false
	}
	if safe(hm, i, j-1) && val >= hm[i][j-1] {
		return false
	}
	if safe(hm, i, j+1) && val >= hm[i][j+1] {
		return false
	}
	return true
}

func riskLevel(height int) int {
	return height + 1
}
