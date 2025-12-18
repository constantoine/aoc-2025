package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
)

type PartialPuzzle interface {
	SolveOne(input string) string
}

type CompletePuzzle interface {
	PartialPuzzle
	SolveTwo(input string) string
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: ./aoc-2025 [day] [part] < input.txt")
	}

	day, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	}

	part, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	}

	days := map[int]PartialPuzzle{
		1: DayOne{},
		2: DayTwo{},
		3: DayThree{},
		4: DayFour{},
	}

	solver, ok := days[day]
	if !ok {
		panic("invalid day")
	}

	input, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic("could not read input: " + err.Error())
	}

	switch part {
	case 1:
		fmt.Println(solver.SolveOne(string(input)))
	case 2:
		solver, ok := solver.(CompletePuzzle)
		if !ok {
			panic("puzzle only implements part one")
		}

		fmt.Println(solver.SolveTwo(string(input)))
	default:
		panic("invalid part")
	}
}
