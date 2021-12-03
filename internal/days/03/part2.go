package _3

import (
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

type LifeSupportMeasurement int64

const (
	OxygenGenerator LifeSupportMeasurement = iota
	Co2Scrubber
)

func SolvePart2() string {
	r := strings.NewReader(input)
	lines := scanner.ScanLines(r)
	return strconv.FormatInt(lifeSupportRating(lines), 10)
}

func lifeSupportRating(report []string) int64 {
	ogr := rating(report, OxygenGenerator)
	co2sr := rating(report, Co2Scrubber)

	return ogr * co2sr
}

func rating(report []string, rating LifeSupportMeasurement) int64 {
	pos := 0
	for ok := true; ok; ok = len(report) > 1 {
		filtered := make([]string, 0)

		expected := ""
		if rating == OxygenGenerator {
			gr := powerConsumptionRate(report, Gamma)
			expected = string(gr[pos])
		} else if rating == Co2Scrubber {
			er := powerConsumptionRate(report, Epsilon)
			expected = string(er[pos])
		}

		for _, bits := range report {
			if string(bits[pos]) == expected {
				filtered = append(filtered, bits)
			}
		}
		report = filtered
		pos++
	}

	r, _ := strconv.ParseInt(report[0], 2, 64)
	return r
}
