package _8

import (
	"strings"
	"testing"

	"github.com/williamgcampbell/aoc2021/internal/scanner"

	"github.com/stretchr/testify/require"
)

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	day := &Day{}
	require.Equal(t, day.SolvePart2(), "1084606")
}

func TestOutputValues(t *testing.T) {
	tests := map[string]struct {
		val  string
		want int
	}{
		"Advent of code example": {
			val:  `acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf`,
			want: 5353,
		},
		"Advent of code example 2": {
			val:  `be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe`,
			want: 8394,
		},
		"Advent of code example 3": {
			val:  `edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc`,
			want: 9781,
		},
	}
	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r := strings.NewReader(test.val)
			actual := outputValues(scanner.ScanLines(r))
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
