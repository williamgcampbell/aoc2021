package _14

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

var (
	delimiter = ";"
)

// SolvePart1 solves part one
func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	v := scanner.ScanUntilEmptyLine(r, delimiter)
	return strconv.Itoa(calculatePolymer(v, 10))
}

// calculatePolymer keeps a map of pair -> count to avoid keeping the full polymer in memory,
// which gets completely unmanageable quickly.
func calculatePolymer(lines []string, steps int) int {
	template := make(map[string]int)
	polymer := lines[0]
	for i := 1; i < len(polymer); i++ {
		pair := polymer[i-1 : i+1]
		if _, ok := template[pair]; !ok {
			template[pair] = 0
		}
		template[pair]++
	}

	rules := getRules(lines[1])
	for i := 0; i < steps; i++ {
		template = step(template, rules)
	}

	// counting pairs means that we assume the end of one pair is the beginning
	// of another pair. This is not the case for the ends of the code
	start, end := string(polymer[0]), string(polymer[len(polymer)-1])
	cm := countMap(template, start, end)

	return calculate(cm)
}

func getRules(r string) map[string]string {
	pairInsertions := strings.Split(r, delimiter)
	pairs := make(map[string]string)
	for _, pair := range pairInsertions {
		p := strings.Split(pair, " -> ")
		pairs[p[0]] = p[1]
	}
	return pairs
}

func step(template map[string]int, rules map[string]string) map[string]int {
	temp := copyMap(template)
	for pair, count := range template {
		if val, ok := rules[pair]; ok {
			applyRule(temp, pair, val, count)
		}
	}
	return temp
}

func applyRule(template map[string]int, originalPair, val string, count int) {
	firstP := string(originalPair[0]) + val
	if _, ok := template[firstP]; !ok {
		template[firstP] = 0
	}

	secondP := val + string(originalPair[1])
	if _, ok := template[secondP]; !ok {
		template[secondP] = 0
	}

	template[firstP] += count
	template[secondP] += count
	template[originalPair] -= count
}

func copyMap(o map[string]int) map[string]int {
	n := make(map[string]int)
	for k, v := range o {
		n[k] = v
	}
	return n
}

func calculate(cm map[string]int) int {
	most, least := 0, 0
	for _, cnt := range cm {
		if most == 0 || cnt > most {
			most = cnt
		}
		if least == 0 || cnt < least {
			least = cnt
		}
	}
	return most - least
}

func countMap(pairMap map[string]int, start, end string) map[string]int {
	r := make(map[string]int)
	for pair, v := range pairMap {
		f := string(pair[0])
		if _, ok := r[f]; !ok {
			r[f] = 0
		}

		s := string(pair[1])
		if _, ok := r[s]; !ok {
			r[s] = 0
		}

		r[s] += v
		r[f] += v
	}

	// The start/end aren't represented twice like all other values. Double it.
	r[start] += 1
	r[end] += 1

	for k := range r {
		r[k] = r[k] / 2
	}

	return r
}
