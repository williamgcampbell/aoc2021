package _1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart2(t *testing.T) {
	require.Equal(t, SolvePart2(), "1724")
}

func part2Tests() []struct {
	nums []int
	want int
} {
	return []struct {
		nums []int
		want int
	}{
		{
			nums: []int{},
			want: 0,
		},
		{
			nums: []int{1, 2, 3},
			want: 0,
		},
		{
			nums: []int{1, 1, 1, 2, 2, 2},
			want: 3,
		},
		{
			nums: []int{1, 1, 1, 2, 2, 2, 1, 1, 1, 2, 2, 2},
			want: 6,
		},
		{
			nums: []int{1, 1, 1, 2, 2, 2, 3, 3, 3, 100},
			want: 7,
		},
		{
			nums: []int{199, 200, 208, 210, 200, 207, 240, 269, 260, 263},
			want: 5,
		},
	}
}

func TestSlidingWindowDepthIncrease(t *testing.T) {
	tests := part2Tests()

	for _, test := range tests {
		actual := slidingWindowDepthIncrease(test.nums)
		if actual != test.want {
			t.Errorf("Got: %d, Want: %d.", actual, test.want)
		}
	}
}

func TestSlidingWindowDepthIncreaseFast(t *testing.T) {
	tests := part2Tests()

	for _, test := range tests {
		actual := slidingWindowDepthIncreaseFast(test.nums)
		if actual != test.want {
			t.Errorf("Got: %d, Want: %d.", actual, test.want)
		}
	}
}
