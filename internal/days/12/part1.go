package _12

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
	edges := scanner.ScanLines(r)
	return strconv.Itoa(countPaths(edges, false))
}

func countPaths(edges []string, haveTime bool) int {
	em := edgeMap(edges)
	return exploreCave("start", em, make([]string, 0), haveTime)
}

func edgeMap(edges []string) map[string][]string {
	em := make(map[string][]string, len(edges)*2)
	for _, edge := range edges {
		parts := strings.Split(edge, "-")
		if em[parts[0]] == nil {
			em[parts[0]] = make([]string, 0)
		}
		if em[parts[1]] == nil {
			em[parts[1]] = make([]string, 0)
		}
		em[parts[0]] = append(em[parts[0]], parts[1])
		em[parts[1]] = append(em[parts[1]], parts[0])
	}
	return em
}

// exploreCave navigates all permutations of routes from 'start' to 'end' recursively.
// Uppercase caves can be navigated to in route any number of times.
// Lowercase caves can only be navigated to once
// Routes must be copied and passed around each iteration to track visited caves
func exploreCave(cave string, edges map[string][]string, route []string, haveTime bool) int {

	newRoute := make([]string, len(route)+1)
	copy(newRoute, route)
	newRoute[len(newRoute)-1] = cave

	c := 0
	caves := edges[cave]
	for _, nextCave := range caves {
		if nextCave == "end" {
			c += 1
		} else if bigCave(nextCave) || !contains(nextCave, newRoute) {
			c += exploreCave(nextCave, edges, newRoute, haveTime)
		} else if haveTime && smallCave(nextCave) {
			if !visitedSmallCaveTwice(newRoute) {
				c += exploreCave(nextCave, edges, newRoute, haveTime)
			}
		}
	}

	return c

}

func visitedSmallCaveTwice(route []string) bool {
	for _, r := range route {
		if smallCave(r) {
			c := count(r, route)
			if c > 1 {
				return true
			}
		}
	}
	return false
}

func smallCave(s string) bool {
	return strings.ToLower(s) == s && s != "start" && s != "end"
}

func bigCave(s string) bool {
	return strings.ToUpper(s) == s
}

func count(str string, arr []string) int {
	c := 0
	for _, val := range arr {
		if val == str {
			c += 1
		}
	}
	return c
}

func contains(str string, arr []string) bool {
	for _, val := range arr {
		if val == str {
			return true
		}
	}
	return false
}
