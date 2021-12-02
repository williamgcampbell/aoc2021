package _3

import (
	_ "embed"
	"strings"
)

//go:embed input.txt
var input string

func SolvePart1() string {
	_ = strings.NewReader(input)
	return ""
}
