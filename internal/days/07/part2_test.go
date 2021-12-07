package _7

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	day := &Day{}
	require.Equal(t, day.SolvePart2(), "96744904")
}

func TestMinimumFuelSpent2(t *testing.T) {
	tests := map[string]struct {
		positions []int
		want      int
	}{
		"Advent of code example": {
			positions: []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14},
			want:      168,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := minimumFuelSpent(test.positions, false)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}

func TestSeq(t *testing.T) {
	tests := map[string]struct {
		val  int
		want int
	}{
		"0": {
			val:  0,
			want: 0,
		},
		"1": {
			val:  1,
			want: 1,
		},
		"3": {
			val:  3,
			want: 6,
		},
		"11": {
			val:  11,
			want: 66,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := seq(test.val)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
