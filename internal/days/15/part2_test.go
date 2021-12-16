package _15

import (
	"reflect"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/williamgcampbell/aoc2021/internal/helpers"
	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	day := &Day{}
	require.Equal(t, day.SolvePart2(), "2976")
}

func TestLowestRisk2(t *testing.T) {
	tests := map[string]struct {
		vals string
		want int
	}{
		"Advent of code example": {
			vals: `1163751742
1381373672
2136511328
3694931569
7463417111
1319128137
1359912421
3125421639
1293138521
2311944581`,
			want: 315,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r := strings.NewReader(test.vals)
			v := scanner.ScanLines(r)
			actual := lowestRisk2(v)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}

func TestLargeMap(t *testing.T) {
	tests := map[string]struct {
		original string
		want     string
		times    int
	}{
		"Advent of code example": {
			original: `8`,
			want: `89123
91234
12345
23456
34567
`,
			times: 5,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r := strings.NewReader(test.original)
			v := scanner.ScanLines(r)
			actual := largeMap(helpers.PlanarInt(v), test.times)

			r2 := strings.NewReader(test.want)
			wantArr := helpers.PlanarInt(scanner.ScanLines(r2))

			if !reflect.DeepEqual(actual, wantArr) {
				t.Errorf("Got: %v, Want: %v.", actual, wantArr)
			}
		})
	}
}
