package _12

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

func TestSolvePart1(t *testing.T) {
	t.Parallel()
	day := &Day{}
	require.Equal(t, day.SolvePart1(), "4378")
}

func TestCountPaths(t *testing.T) {
	tests := map[string]struct {
		vals string
		want int
	}{
		"Advent of code example": {
			vals: `start-A
start-b
A-c
A-b
b-d
A-end
b-end`,
			want: 10,
		},
		"Advent of code example 2": {
			vals: `dc-end
HN-start
start-kj
dc-start
dc-HN
LN-dc
HN-end
kj-sa
kj-HN
kj-dc`,
			want: 19,
		},
		"Advent of code example 3": {
			vals: `fs-end
he-DX
fs-he
start-DX
pj-DX
end-zg
zg-sl
zg-pj
pj-he
RW-he
fs-DX
pj-RW
zg-RW
start-pj
he-WI
zg-he
pj-fs
start-RW`,
			want: 226,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r := strings.NewReader(test.vals)
			actual := countPaths(scanner.ScanLines(r), false)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
