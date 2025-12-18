package main_test

import (
	"testing"

	"github/constantoine/aoc-2025"
)

func TestFourPartOne(t *testing.T) {
	testPuzzle(t, 4, 1, func() main.PartialPuzzle {
		return main.DayFour{}
	})
}

func TestFourPartTwo(t *testing.T) {
	testPuzzle(t, 4, 2, func() main.PartialPuzzle {
		return main.DayFour{}
	})
}
