package _1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	require.Equal(t, SolvePart2(), "1724")
}

func TestSlidingWindowDepthIncrease(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"Empty": {
			nums: []int{},
			want: 0,
		},
		"Simple increasing": {
			nums: []int{1, 2, 3},
			want: 0,
		},
		"Sliding window increasing": {
			nums: []int{1, 1, 1, 2, 2, 2},
			want: 3,
		},
		"Fluctuating depths": {
			nums: []int{1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2},
			want: 6,
		},
		"Sliding window increase part 2": {
			nums: []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 100},
			want: 7,
		},
		"Advent of code example": {
			nums: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			want: 5,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := slidingWindowDepthIncrease(test.nums)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}

func TestSlidingWindowDepthIncreaseFast(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"Empty": {
			nums: []int{},
			want: 0,
		},
		"Simple increasing": {
			nums: []int{1, 2, 3},
			want: 0,
		},
		"Sliding window increasing": {
			nums: []int{1, 1, 1, 2, 2, 2},
			want: 3,
		},
		"Fluctuating depths": {
			nums: []int{1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2},
			want: 6,
		},
		"Sliding window increase part 2": {
			nums: []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 100},
			want: 7,
		},
		"Advent of code example": {
			nums: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			want: 5,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := slidingWindowDepthIncreaseFast(test.nums)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
