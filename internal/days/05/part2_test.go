package _5

import (
	"strings"
	"testing"

	"github.com/williamgcampbell/aoc2021/internal/scanner"

	"github.com/stretchr/testify/require"
)

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	day := &Day{}
	require.Equal(t, day.SolvePart2(), "21373")
}

func TestOverlappingLines2(t *testing.T) {
	tests := map[string]struct {
		lines string
		want  int
	}{
		"Advent of code example": {
			lines: `0,9 -> 5,9
8,0 -> 0,8
9,4 -> 3,4
2,2 -> 2,1
7,0 -> 7,4
6,4 -> 2,0
0,9 -> 2,9
3,4 -> 1,4
0,0 -> 8,8
5,5 -> 8,2`,
			want: 12,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r := strings.NewReader(test.lines)
			actual := overlappingLines(scanner.ScanLines(r), false)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
