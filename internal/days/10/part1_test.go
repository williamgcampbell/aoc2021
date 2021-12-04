package _10

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart1(t *testing.T) {
	require.Equal(t, SolvePart1(), "0")
}

func TestTodo(t *testing.T) {
	tests := map[string]struct {
		vals []string
		want int
	}{
		"Advent of code example": {
			vals: []string{},
			want: 0,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := todo(test.vals)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
