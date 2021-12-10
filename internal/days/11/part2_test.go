package _11

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	day := &Day{}
	require.Equal(t, day.SolvePart2(), "0")
}

func TestTodo2(t *testing.T) {
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
			actual := todo2(test.vals)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
