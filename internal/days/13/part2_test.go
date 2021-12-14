package _13

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	day := &Day{}
	fmt.Println(day.SolvePart2())
	require.Equal(t, day.SolvePart2(), `
 ##  #  #  ##   ##  ###   ##   ##  #  #
#  # #  # #  # #  # #  # #  # #  # #  #
#  # #### #    #    #  # #    #  # #  #
#### #  # # ## #    ###  # ## #### #  #
#  # #  # #  # #  # #    #  # #  # #  #
#  # #  #  ###  ##  #     ### #  #  ## 
`)
}

func TestDecode(t *testing.T) {
	tests := map[string]struct {
		vals string
		want string
	}{
		"Advent of code example": {
			vals: `6,10
0,14
9,10
0,3
10,4
4,11
6,0
6,12
4,1
0,13
10,12
3,4
3,0
8,4
1,10
2,14
8,10
9,0

fold along y=7
fold along x=5`,
			want: "\n#####\n#   #\n#   #\n#   #\n#####\n",
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			r := strings.NewReader(test.vals)
			actual := decode(scanner.ScanUntilEmptyLine(r, delimiter))
			if actual != test.want {
				t.Errorf("Got: %s, Want: %s.", actual, test.want)
			}
		})
	}
}
