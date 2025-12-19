package main_test

import (
	"testing"

	"github/constantoine/aoc-2025"
)

func TestFivePartOne(t *testing.T) {
	testPuzzle(t, 5, 1, func() main.PartialPuzzle {
		return main.DayFive{}
	})
}

func TestFivePartTwo(t *testing.T) {
	testPuzzle(t, 5, 2, func() main.PartialPuzzle {
		return main.DayFive{}
	})
}
