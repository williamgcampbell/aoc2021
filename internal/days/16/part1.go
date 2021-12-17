package _16

import (
	_ "embed"
	"fmt"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/helpers"
	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

// SolvePart1 solves part one
func (d *Day) SolvePart1() string {
	r := strings.NewReader(input)
	v := scanner.ScanLines(r)
	return strconv.Itoa(sumOfVersionNumbers(v[0]))
}

func sumOfVersionNumbers(hex string) int {
	bitStr := toBits(hex)
	parser := newPacketParser(bitStr)
	packets := ParseCount(parser, 1)

	sumVersion := 0
	for len(packets) > 0 {
		pkt := packets[0]
		packets = packets[1:]
		sumVersion += pkt.Version()
		for _, p := range pkt.ChildPackets() {
			packets = append(packets, p)
		}
	}
	return sumVersion
}

func toBits(code string) string {
	var sb strings.Builder
	for i := 0; i < len(strings.TrimSpace(code)); i++ {
		sub := helpers.MustParseUint(string(code[i]), 16, 4)
		sb.WriteString(fmt.Sprintf("%04b", sub))
	}

	return sb.String()
}
