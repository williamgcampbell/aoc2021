package _9

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	day := &Day{}
	require.Equal(t, day.SolvePart1(), "475")
}

func TestSumOfRiskLevels(t *testing.T) {
	tests := map[string]struct {
		vals string
		want int
	}{
		"Advent of code example": {
			vals: `2199943210
3987894921
9856789892
8767896789
9899965678`,
			want: 15,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r := strings.NewReader(test.vals)
			actual := sumOfRiskLevels(scanner.ScanLines(r))
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
