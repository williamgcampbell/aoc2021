package _8

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(easyDigitOutputCount(v))
}

func easyDigitOutputCount(lines []string) int {
	count := 0
	for _, str := range lines {
		output := strings.Split(str, "|")[1]
		nums := strings.Split(output, " ")
		for _, digits := range nums {
			switch len(digits) {
			case 2: // 1
				fallthrough
			case 3: // 7
				fallthrough
			case 4: // 4
				fallthrough
			case 7: // 8
				count++
			}
		}
	}
	return count
}
