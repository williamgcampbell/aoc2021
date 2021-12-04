package _4

import (
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

func SolvePart2() string {
	r := strings.NewReader(input)
	game := scanner.ScanUntilEmptyLine(r, BoardSeparator)
	return strconv.Itoa(playToLose(game))
}

func playToLose(game []string) int {
	draw := game[0]

	boards := make([]*board, 0, len(game)-1)
	for i := 1; i < len(game); i++ {
		boards = append(boards, newBoard(game[i]))
	}

	drawArr := strings.Split(draw, ",")
	for _, sv := range drawArr {
		tmp := boards[:0]
		v, _ := strconv.Atoi(sv)
		for _, brd := range boards {
			if !brd.markAndCheck(v) {
				tmp = append(tmp, brd)
			} else if len(boards) == 1 {
				return boards[0].score(v)
			}
		}
		boards = tmp
	}

	return 0
}
