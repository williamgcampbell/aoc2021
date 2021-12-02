package _1

import (
	_ "embed"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

func SolvePart1() string {
	r := strings.NewReader(input)
	depths, err := scanner.ScanIntLines(r)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "Error scanning values: %v\n", err)
		return ""
	}
	return strconv.Itoa(depthIncrease(depths))
}

func depthIncrease(depths []int) int {
	if len(depths) == 0 {
		return 0
	}

	c := 0
	pd := depths[0]
	for _, d := range depths {
		if d > pd {
			c += 1
		}
		pd = d
	}
	return c
}
