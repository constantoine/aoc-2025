package main_test

import (
	"testing"

	"github/constantoine/aoc-2025"
)

func TestTwoPartOne(t *testing.T) {
	testPuzzle(t, 2, 1, func() main.PartialPuzzle {
		return main.DayTwo{}
	})
}

func TestTwoPartTwo(t *testing.T) {
	testPuzzle(t, 2, 2, func() main.PartialPuzzle {
		return main.DayTwo{}
	})
}
