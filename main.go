package main

import (
	"fmt"
	"log"

	"github.com/williamgcampbell/aoc2021/internal/registry"

	"github.com/williamgcampbell/aoc2021/internal/days"
)

var chorus = "On the %s day of Christmas the part %s solution was: %s\n"

func main() {
	r := registry.NewDayRegistry()
	days.RegisterAll(r)

	for i := 1; i <= 25; i++ {
		if day, ok := r[i]; ok {
			printDay(day)
		} else {
			log.Fatalf("Could not find day %d", i)
		}
	}
}

func printDay(day registry.Day) {
	printDayPart(1, day)
	printDayPart(2, day)
}

func printDayPart(part int, day registry.Day) {
	switch part {
	case 1:
		fmt.Printf(chorus, ordinalString(day.GetName()), "1", normalizeSolution(day.SolvePart1()))
	case 2:
		fmt.Printf(chorus, ordinalString(day.GetName()), "2", normalizeSolution(day.SolvePart2()))
	}
}

func normalizeSolution(r string) string {
	if r == "0" {
		return "I don't know"
	}
	return r
}

func ordinalString(i int) string {
	j := i % 10
	k := i % 100
	if j == 1 && k != 11 {
		return fmt.Sprintf("%dst", i)
	}
	if j == 2 && k != 12 {
		return fmt.Sprintf("%dnd", i)
	}
	if j == 3 && k != 13 {
		return fmt.Sprintf("%drd", i)
	}
	return fmt.Sprintf("%dth", i)
}
