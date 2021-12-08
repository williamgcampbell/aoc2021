package _8

import (
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

func (d *Day) SolvePart2() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(outputValues(v))
}

func outputValues(lines []string) int {
	sum := 0
	for _, line := range lines {
		parts := strings.Split(line, " | ")
		digits := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")

		codes := make(map[int]string)
		codes[1], digits = removeCode(digits, 2, "")
		codes[4], digits = removeCode(digits, 4, "")
		codes[7], digits = removeCode(digits, 3, "")
		codes[8], digits = removeCode(digits, 7, "")
		codes[3], digits = removeCode(digits, 5, codes[1])
		codes[5], digits = removeCode(digits, 5, difference(codes[4], codes[1]))
		codes[2], digits = removeCode(digits, 5, "")
		codes[9], digits = removeCode(digits, 6, codes[4])
		codes[0], digits = removeCode(digits, 6, codes[1])
		codes[6], digits = removeCode(digits, 6, "")

		sb := ""
		for _, s := range output {
			var key int
			for k, code := range codes {
				if len(code) == len(s) && containsSet(code, s) {
					key = k
					break
				}
			}
			sb += strconv.Itoa(key)
		}
		v, _ := strconv.Atoi(sb)
		sum += v
	}
	return sum
}

// removeCode is a specialized function which removes the first string that matches the length and contains string, and
// returns the string removed, as well as the updated array
func removeCode(inputs []string, length int, contains string) (string, []string) {
	for i, input := range inputs {
		cv := true
		if len(contains) > 0 && !containsSet(input, contains) {
			cv = false
		}
		if len(input) == length && cv {
			return input, remove(inputs, i)
		}
	}
	return "", inputs
}

func remove(s []string, i int) []string {
	s[i] = s[len(s)-1]
	return s[:len(s)-1]
}

// containsSet returns true if every character in string b exists in string a
func containsSet(a, b string) bool {
	for _, item := range b {
		found := false
		for _, item2 := range a {
			if string(item2) == string(item) {
				found = true
				break
			}
		}
		if !found {
			return false
		}
	}
	return true
}

// difference returns the character difference between a and b, not caring about order
func difference(a, b string) (diff string) {
	m := make(map[string]bool)

	for _, item := range b {
		m[string(item)] = true
	}

	for _, item := range a {
		if _, ok := m[string(item)]; !ok {
			diff += string(item)
		}
	}
	return
}
