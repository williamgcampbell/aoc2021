package _7

import (
	_ "embed"
	"math"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/helpers"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	positions, _ := scanner.ScanCSVInt(r)
	return strconv.Itoa(minimumFuelSpent(positions[0], true))
}

func minimumFuelSpent(positions []int, constantBurn bool) int {
	min := fuelSpent(0, positions, constantBurn)
	stop := len(positions)
	for i := 1; i < stop; i++ {
		spent := fuelSpent(i, positions, constantBurn)
		if spent < min {
			min = spent
		}

		// Not every position is occupied by a crab submarine.
		// As we iterate, discover the upper bound of the available positions.
		if len(positions)-1 < i && positions[i] > stop {
			stop = positions[i]
		}
	}
	return min
}

// A memoized version of the function seq
var memSeq = helpers.MemorizedIntInt(seq)

func fuelSpent(dest int, positions []int, constantBurn bool) int {
	fuel := 0
	for _, pos := range positions {
		if constantBurn {
			fuel += int(math.Abs(float64(pos - dest)))
		} else {
			moved := int(math.Abs(float64(pos - dest)))
			fuel += memSeq(moved)
		}
	}
	return fuel
}

func seq(val int) int {
	if val <= 0 {
		return 0
	}

	if val == 1 {
		return 1
	}

	return val + seq(val-1)
}
