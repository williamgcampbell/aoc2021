package _4

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	require.Equal(t, SolvePart1(), "0")
}

func TestTodoPart1(t *testing.T) {
	tests := map[string]struct {
		report []string
		want   int64
	}{
		"Advent of code example": {
			report: []string{""},
			want:   0,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := todo(test.report)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
