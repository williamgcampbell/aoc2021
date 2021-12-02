// Copyright 2020-present (c) Care.com, Inc.
//
// All rights reserved.
//
// This software is the confidential and proprietary information of
// Care.com, Inc.

package _1

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart1(t *testing.T) {
	require.Equal(t, SolvePart1(), "1692")
}

func TestDepthIncrease(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{
			nums: []int{1, 2, 3, 4},
			want: 3,
		},
		{
			nums: []int{1, 1, 0, 1},
			want: 1,
		},
		{
			nums: []int{},
			want: 0,
		},
	}

	for _, test := range tests {
		actual := depthIncrease(test.nums)
		if actual != test.want {
			t.Errorf("Got: %d, Want: %d.", actual, test.want)
		}
	}
}
