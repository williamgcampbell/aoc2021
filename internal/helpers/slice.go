package helpers

import (
	"sort"
	"strconv"
	"strings"
)

// InsertSorted
func InsertSorted(ss []int, s int) []int {
	i := sort.SearchInts(ss, s)
	ss = append(ss, 0)
	copy(ss[i+1:], ss[i:])
	ss[i] = s
	return ss
}

// PlanarInt extracts a 2D array of ints from an array of strings
func PlanarInt(lines []string) [][]int {
	r := make([][]int, 0, len(lines))
	for _, line := range lines {
		ints := strings.Split(line, "")
		temp := make([]int, 0, len(ints))
		for _, is := range ints {
			i, _ := strconv.Atoi(is)
			temp = append(temp, i)
		}
		r = append(r, temp)
	}
	return r
}

// Safe returns true if i, j respectively are within the bounds of the 2D array
func Safe(hm [][]int, i, j int) bool {
	if i < 0 || j < 0 || i >= len(hm) || j >= len(hm[i]) {
		return false
	}
	return true
}
