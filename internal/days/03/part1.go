package _3

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

func SolvePart1() string {
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

	gammaRate := ""
	for _, mc := range mcs {
		if mc >= 0 {
			if measurement == Gamma {
				gammaRate += "1"
			} else {
				gammaRate += "0"
			}
		} else {
			if measurement == Gamma {
				gammaRate += "0"
			} else {
				gammaRate += "1"
			}
		}
	}

	return gammaRate
}
