package _2

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	require.Equal(t, SolvePart1(), "2019945")
}

func TestCalculatePosition(t *testing.T) {
	tests := map[string]struct {
		nums []string
		want int
	}{
		"Simple depth increase": {
			nums: []string{"forward 5", "down 5", "forward 8", "up 3", "down 8", "forward 2"},
			want: 150,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := calculatePosition(test.nums, false)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
