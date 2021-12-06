package _6

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	day := &Day{}
	require.Equal(t, day.SolvePart1(), "345387")
}

func TestCountAll(t *testing.T) {
	tests := map[string]struct {
		vals      []int
		afterDays int
		want      int64
	}{
		"Advent of code example": {
			vals:      []int{3, 4, 3, 1, 2},
			afterDays: 80,
			want:      5934,
		},
		"Advent of code example Part 2": {
			vals:      []int{3, 4, 3, 1, 2},
			afterDays: 256,
			want:      26984457539,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := countAll(test.vals, test.afterDays)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
