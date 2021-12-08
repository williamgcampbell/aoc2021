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
		code := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")

		codes := make(map[int]string)
		// This was a pain and order does matter here.
		// There are a handful of uniquely identifiable numbers just by the length of the code
		codes[1], code = removeCode(code, 2, "")
		codes[4], code = removeCode(code, 4, "")
		codes[7], code = removeCode(code, 3, "")
		codes[8], code = removeCode(code, 7, "")

		// These are more complex. For instance, the number 3 unique in that it is made up of 5 codes and contains
		// all-of-the codes from the number 1. In order to identify 3 we need to populate 1 first.
		codes[3], code = removeCode(code, 5, codes[1])

		// 5 will have codes from 4, as long as you take away the upper right value. We do this by taking the difference
		// with the codes from 1. It loses the entire right side but it's still unique.
		codes[5], code = removeCode(code, 5, difference(codes[4], codes[1]))

		// At this point there are no longer any codes of length 5 that isn't a 2
		codes[2], code = removeCode(code, 5, "")

		// 9 contains everything from 4, but has six values
		codes[9], code = removeCode(code, 6, codes[4])

		// Now that 9 is out we can identify 0 by looking at six-digit codes with all codes from 1
		codes[0], code = removeCode(code, 6, codes[1])

		// This is the last one so code should only have one value left, sanity check on length
		codes[6], code = removeCode(code, 6, "")

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
