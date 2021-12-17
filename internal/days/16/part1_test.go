package _16

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	day := &Day{}
	require.Equal(t, day.SolvePart1(), "1007")
}

func TestSumOfVersionNumbers(t *testing.T) {
	tests := map[string]struct {
		vals string
		want int
	}{
		"Advent of code example": {
			vals: "8A004A801A8002F478",
			want: 16,
		},
		"Advent of code example 2": {
			vals: "620080001611562C8802118E34",
			want: 12,
		},
		"Advent of code example 3": {
			vals: "C0015000016115A2E0802F182340",
			want: 23,
		},
		"Advent of code example 4": {
			vals: "A0016C880162017C3686B18A3D4780",
			want: 31,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := sumOfVersionNumbers(test.vals)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
