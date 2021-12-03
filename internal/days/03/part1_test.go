package _3

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	require.Equal(t, SolvePart1(), "2640986")
}

func TestCalculatePowerConsumption(t *testing.T) {
	tests := map[string]struct {
		report []string
		want   int64
	}{
		"Simple depth increase": {
			report: []string{"00100", "11110", "10110", "10111", "10101", "01111", "00111", "11100", "10000", "11001", "00010", "01010"},
			want:   198,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := powerConsumption(test.report)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
