package _7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	day := &Day{}
	require.Equal(t, day.SolvePart1(), "343605")
}

func TestMinimumFuelSpent(t *testing.T) {
	tests := map[string]struct {
		positions []int
		want      int
	}{
		"Advent of code example": {
			positions: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			want:      37,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := minimumFuelSpent(test.positions, true)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
