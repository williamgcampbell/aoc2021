package _15

import (
	_ "embed"
	"fmt"
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
	return strconv.Itoa(lowestRisk(v))
}

func lowestRisk(lines []string) int {
	rvs := helpers.PlanarInt(lines)

	start := loc{
		row: 0,
		col: 0,
	}

	r := buildRiskScore(rvs, start)

	return r[len(r)-1][len(r[len(rvs)-1])-1]
}

func buildRiskScore(input [][]int, start loc) [][]int {
	// clever way to represent all four locations surrounding a value in a 2D space
	rn := []int{-1, 0, 0, 1}
	cn := []int{0, -1, 1, 0}

	visited := make([][]bool, len(input))
	riskAgg := make([][]int, len(input))
	for i, v := range input {
		visited[i] = make([]bool, len(v))
		riskAgg[i] = make([]int, len(v))
	}

	var pq []loc
	// kick things off by adding start location to the stack
	pq = append(pq, start)

	for len(pq) > 0 {
		// It's important to always pull the start since the priority queue is ordered by risk, ascending
		// This ensures that you are navigating the graph from smallest risk to largest risk
		current := pq[0]
		pq = pq[1:]

		visited[current.row][current.col] = true
		currRiskAgg := riskAgg[current.row][current.col]
		for k := 0; k < 4; k++ {
			nextLoc := loc{
				row: current.row + rn[k],
				col: current.col + cn[k],
			}
			if helpers.Safe(input, nextLoc.row, nextLoc.col) && !visited[nextLoc.row][nextLoc.col] {
				nextRisk := input[nextLoc.row][nextLoc.col]
				// encountering a value that has no risk set (0) or has a higher risk than the current path means we are on the least risky path
				if riskAgg[nextLoc.row][nextLoc.col] == 0 || currRiskAgg+nextRisk < riskAgg[nextLoc.row][nextLoc.col] {
					riskAgg[nextLoc.row][nextLoc.col] = currRiskAgg + nextRisk
					nextLoc.risk = currRiskAgg + nextRisk
					pq = insert(pq, nextLoc)
				}
			}
		}
	}
	return riskAgg
}

// insert adds val in the priority queue in order of risk value.
// This is an important property of Dijkstra's algorithm to ensure you are always navigating the smallest path first
func insert(pq []loc, val loc) []loc {
	if len(pq) == 0 {
		pq = append(pq, val)
		return pq
	}

	var inserted bool
	for k, v := range pq {
		if val.risk < v.risk {
			if k > 0 {
				pq = append(pq[:k+1], pq[k:]...)
				pq[k] = val
			} else {
				pq = append([]loc{val}, pq...)
			}
			inserted = true
			break
		}
	}
	if !inserted {
		pq = append(pq, val)
	}
	return pq
}

type loc struct {
	row  int
	col  int
	risk int
}

func (l *loc) String() string {
	return fmt.Sprintf("%d,%d", l.row, l.col)
}
