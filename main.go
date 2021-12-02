package main

import (
	_1 "aoc2021/internal/days/01"
	_2 "aoc2021/internal/days/02"
	_3 "aoc2021/internal/days/03"
	"fmt"
)

var chorus = "On the %s day of Christmas the part %s solution was: %s\n"

func main() {
	fmt.Printf(chorus, "1st", "1", _1.SolvePart1())
	fmt.Printf(chorus, "1st", "2", _1.SolvePart2())
	fmt.Printf(chorus, "2nd", "1", _2.SolvePart1())
	fmt.Printf(chorus, "2nd", "2", _2.SolvePart2())
	fmt.Printf(chorus, "3rd", "1", _3.SolvePart1())
	fmt.Printf(chorus, "3rd", "2", _3.SolvePart2())
}
