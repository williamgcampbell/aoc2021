package _10

import (
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/helpers"
	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

var (
	completionScores = map[string]int{
		")": 1,
		"]": 2,
		"}": 3,
		">": 4,
	}
	pairs = map[string]string{
		"(": ")",
		"[": "]",
		"{": "}",
		"<": ">",
	}
)

func (d *Day) SolvePart2() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(middleCompletionScore(v))
}

func middleCompletionScore(lines []string) int {
	var scores []int
	for _, line := range lines {
		s, b := completionScore(line)
		if b {
			scores = helpers.InsertSorted(scores, s)
		}
	}
	m := len(scores) / 2
	return scores[m]
}

func completionScore(line string) (int, bool) {
	if syntaxScoreLine(line) > 0 {
		return 0, false
	}

	characters := strings.Split(line, "")
	stack := make([]string, 0, len(characters))
	for _, character := range characters {
		if _, ok := pairs[character]; ok {
			stack = append(stack, character)
		} else {
			n := len(stack) - 1
			stack = stack[:n]
		}
	}

	return tallyScore(stack), true
}

func tallyScore(remaining []string) int {
	totalScore := 0
	for i := len(remaining) - 1; i >= 0; i-- {
		s := remaining[i]
		if score, ok := completionScores[pairs[s]]; ok {
			totalScore = (totalScore * 5) + score
		}
	}
	return totalScore
}
