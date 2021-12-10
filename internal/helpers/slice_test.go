package helpers

import (
	"reflect"
	"testing"
)

func TestTodo2(t *testing.T) {
	tests := map[string]struct {
		vals   []int
		insert int
		want   []int
	}{
		"Empty array works": {
			vals:   []int{},
			insert: 1,
			want:   []int{1},
		},
		"Basic insert": {
			vals:   []int{1, 2, 3, 5},
			insert: 4,
			want:   []int{1, 2, 3, 4, 5},
		},
		"Value inserted before existing value": {
			vals:   []int{1, 2, 3, 4},
			insert: 2,
			want:   []int{1, 2, 2, 3, 4},
		},
	}

	for name, test := range tests {
		test := test
		t.Run(name, func(t *testing.T) {
			t.Parallel()
			actual := InsertSorted(test.vals, test.insert)
			if !reflect.DeepEqual(actual, test.want) {
				t.Errorf("Got: %d, Want: %d.", actual, test.want)
			}
		})
	}
}
