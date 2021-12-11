package _11

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	day := &Day{}
	require.Equal(t, day.SolvePart1(), "1594")
}

func TestTotalFlashes(t *testing.T) {
	tests := map[string]struct {
		vals  string
		want  int
		steps int
	}{
		"Advent of code example": {
			vals: `5483143223
2745854711
5264556173
6141336146
6357385478
4167524645
2176841721
6882881134
4846848554
5283751526`,
			want:  1656,
			steps: 100,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r := strings.NewReader(test.vals)
			actual := totalFlashes(scanner.ScanLines(r), test.steps)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
