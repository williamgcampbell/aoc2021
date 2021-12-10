package _10

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

var (
	syntaxScores = map[string]int{
		")": 3,
		"]": 57,
		"}": 1197,
		">": 25137,
	}
	backwardsPairs = map[string]string{
		")": "(",
		"]": "[",
		"}": "{",
		">": "<",
	}
)

func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(syntaxScore(v))
}

func syntaxScore(lines []string) int {
	score := 0
	for _, line := range lines {
		score += syntaxScoreLine(line)
	}
	return score
}

func syntaxScoreLine(line string) int {
	characters := strings.Split(line, "")
	stack := make([]string, 0, len(characters))
	for _, character := range characters {
		if want, ok := backwardsPairs[character]; ok {
			n := len(stack) - 1
			got := stack[n]
			if want != got {
				return syntaxScores[character]
			}
			stack = stack[:n]
		} else {
			stack = append(stack, character)
		}
	}
	return 0
}
