package _4

import (
	_ "embed"
	"strconv"
	"strings"

	"github.com/williamgcampbell/aoc2021/internal/scanner"
)

//go:embed input.txt
var input string

var BoardSeparator = ":"

func SolvePart1() string {
	r := strings.NewReader(input)
	game := scanner.ScanUntilEmptyLine(r, BoardSeparator)
	return strconv.Itoa(playToWin(game))
}

func playToWin(game []string) int {
	draw := game[0]

	boards := make([]*board, 0, len(game)-1)
	for i := 1; i < len(game); i++ {
		boards = append(boards, newBoard(game[i]))
	}

	drawArr := strings.Split(draw, ",")
	for _, sv := range drawArr {
		for _, brd := range boards {
			v, _ := strconv.Atoi(sv)
			if brd.markAndCheck(v) {
				return brd.score(v)
			}
		}
	}

	return 0
}

// newBoard builds a board from a colon separated set of rows.
// e.g. - 1 2; 6 10 would build a 2x2 board
func newBoard(s string) *board {
	var b [][]int
	rows := strings.Split(s, BoardSeparator)
	var sumOfUnmarked int
	for _, rowStr := range rows {
		rowStrArr := strings.Fields(rowStr)
		row := make([]int, 0, len(rowStrArr))
		for _, rv := range rowStrArr {
			val, _ := strconv.Atoi(rv)
			row = append(row, val)
			sumOfUnmarked += val
		}
		b = append(b, row)
	}

	return &board{
		board:         b,
		markedRows:    make([]int, len(b)),
		markedColumns: make([]int, len(b[0])),
		sumOfUnmarked: sumOfUnmarked,
	}
}

type board struct {
	board [][]int

	// markedRows keeps track of how many values per rows have been marked. A board is considered won
	// if the number of marked values at any position is equal to the number of rows
	markedRows []int

	// markedColumns keeps track of how many values per column have been marked. A board is considered won
	// if the number of marked values at any position is equal to the number of columns
	markedColumns []int

	// Keeps track of the sum of all unmarked values. This is used to calculate the final score
	sumOfUnmarked int
}

func (b *board) markAndCheck(number int) bool {
	for r := 0; r < len(b.board); r++ {
		row := b.board[r]
		for c := 0; c < len(row); c++ {
			v := row[c]
			if number == v {
				b.board[r][c] = -1
				b.markedRows[r]++
				b.markedColumns[c]++
				b.sumOfUnmarked -= v
				if b.markedRows[r] == len(b.board) || b.markedColumns[c] == len(b.board[0]) {
					return true
				}
				return false
			}
		}
	}
	return false
}

func (b *board) score(number int) int {
	return b.sumOfUnmarked * number
}
