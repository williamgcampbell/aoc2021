package _1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	require.Equal(t, SolvePart1(), "1692")
}

func TestDepthIncrease(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"Simple depth increase": {
			nums: []int{1, 2, 3, 4},
			want: 3,
		},
		"Depth up and down": {
			nums: []int{1, 1, 0, 1},
			want: 1,
		},
		"Empty depths": {
			nums: []int{},
			want: 0,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := depthIncrease(test.nums)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
