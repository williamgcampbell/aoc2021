package _2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	require.Equal(t, SolvePart2(), "1599311480")
}

func TestCalculatePositionWithAim(t *testing.T) {
	tests := map[string]struct {
		nums []string
		want int
	}{
		"Simple depth increase": {
			nums: []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			want: 900,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := calculatePosition(test.nums, true)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
