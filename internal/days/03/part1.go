package _3

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

// SolvePart1 solves part one
func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	lines := scanner.ScanLines(r)
	return strconv.FormatInt(powerConsumption(lines), 10)
}

type PowerConsumptionMeasurement int64

const (
	Gamma PowerConsumptionMeasurement = iota
	Epsilon
)

func powerConsumption(report []string) int64 {

	gammaRate := powerConsumptionRate(report, Gamma)
	epsilonRate := powerConsumptionRate(report, Epsilon)

	dgr, _ := strconv.ParseInt(gammaRate, 2, 64)
	der, _ := strconv.ParseInt(epsilonRate, 2, 64)

	return dgr * der
}

// powerConsumptionRate returns the most or least common bits for a corresponding position in a bit report.
func powerConsumptionRate(report []string, measurement PowerConsumptionMeasurement) string {
	// When a 1 bit is found for a position the array at that position is incremented, and vice versa for 0.
	// The resulting array will contain a positive value if 1 is the most common bit for a position
	// and negative if 0 is the most common bit.
	mcs := make([]int, len(report[0]))
	for _, bits := range report {
		for i := 0; i < len(bits); i++ {
			if string(bits[i]) == "1" {
				mcs[i]++
			} else {
				mcs[i]--
			}
		}
	}

	rate := ""
	for _, mc := range mcs {
		// Note that zero values mean there are equal number of 1s and 0s. We round up in this case.
		if mc >= 0 {
			if measurement == Gamma {
				rate += "1"
			} else {
				rate += "0"
			}
		} else {
			if measurement == Gamma {
				rate += "0"
			} else {
				rate += "1"
			}
		}
	}

	return rate
}
