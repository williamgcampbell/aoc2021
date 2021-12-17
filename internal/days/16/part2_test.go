package _16

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSolvePart2(t *testing.T) {
	t.Parallel()
	day := &Day{}
	require.Equal(t, day.SolvePart2(), "834151779165")
}

func TestEvaluate(t *testing.T) {
	tests := map[string]struct {
		vals string
		want int
	}{
		"sum of 1 and 2": {
			vals: "C200B40A82",
			want: 3,
		},
		"product of 6 and 9": {
			vals: "04005AC33890",
			want: 54,
		},
		"minimum of 7, 8, and 9": {
			vals: "880086C3E88112",
			want: 7,
		},
		"maximum of 7,8, and 9": {
			vals: "CE00C43D881120",
			want: 9,
		},
		"returns 1, because 1 is less than 15": {
			vals: "D8005AC2A8F0",
			want: 1,
		},
		"returns 0, because 5 is not greater than 15": {
			vals: "F600BC2D8F",
			want: 0,
		},
		"returns 0, because 5 is not equal to 15": {
			vals: "9C005AC2F8F0",
			want: 0,
		},
		"returns 1, because 1 + 3 == 2 * 2": {
			vals: "9C0141080250320F1802104A08",
			want: 1,
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := evaluate(test.vals)
			if actual != test.want {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
